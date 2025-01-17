// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title Staking
 * @notice The EVM interface to the consensus chain's x/staking module.
 *         Calls are proxied, and not executed syncronously. Their execution is left to
 *         the consensus chain, and they may fail.
 * @dev This contract is predeployed, and requires storage slots to be set in genesis.
 *      initialize(...) is called pre-deployment, in sctips/genesis/AllocPredeploys.s.sol
 *      iniitializers on the implementation are disabled via manual storage updates, rather than in a constructor.
 *      If an new implementation is required, a constructor should be added.
 */
contract Staking is OwnableUpgradeable {
    /**
     * @notice Emitted when a validator is created
     * @param validator     (MsgCreateValidator.validator_addr) The address of the validator to create
     * @param pubkey        (MsgCreateValidator.pubkey) The validators consensus public key. 33 bytes compressed secp256k1 public key
     * @param deposit       (MsgCreateValidator.selfDelegation) The validators initial stake
     */
    event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit);

    /**
     * @notice Emitted when a delegation is made to a validator
     * @param delegator     (MsgDelegate.delegator_addr) The address of the delegator
     * @param validator     (MsgDelegate.validator_addr) The address of the validator to delegate to
     * @param amount        (MsgDelegate.amount) The amount of tokens to delegate
     */
    event Delegate(address indexed delegator, address indexed validator, uint256 amount);

    /**
     * @notice Emitted when a validator is allowed to create a validator
     * @param validator     The validator address
     */
    event ValidatorAllowed(address indexed validator);

    /**
     * @notice Emitted when a validator is disallowed to create a validator
     * @param validator     The validator address
     */
    event ValidatorDisallowed(address indexed validator);

    /**
     * @notice Emitted when the allowlist is enabled
     */
    event AllowlistEnabled();

    /**
     * @notice Emitted when the allowlist is disabled
     */
    event AllowlistDisabled();

    /**
     * @notice The minimum deposit required to create a validator
     */
    uint256 public constant MinDeposit = 100 ether;

    /**
     * @notice The minimum delegation required to delegate to a validator
     */
    uint256 public constant MinDelegation = 1 ether;

    /**
     * @notice True of the validator allowlist is enabled.
     */
    bool public isAllowlistEnabled;

    /**
     * @notice True if the validator address is allowed to create a validator.
     */
    mapping(address => bool) public isAllowedValidator;

    constructor() {
        _disableInitializers();
    }

    function initialize(address owner_, bool isAllowlistEnabled_) public initializer {
        __Ownable_init(owner_);
        isAllowlistEnabled = isAllowlistEnabled_;
    }

    /**
     * @notice Create a new validator
     * @param pubkey The validators consensus public key. 33 bytes compressed secp256k1 public key
     * @dev Proxies x/staking.MsgCreateValidator
     */
    function createValidator(bytes calldata pubkey) external payable {
        require(!isAllowlistEnabled || isAllowedValidator[msg.sender], "Staking: not allowed");
        require(pubkey.length == 33, "Staking: invalid pubkey length");
        require(msg.value >= MinDeposit, "Staking: insufficient deposit");

        emit CreateValidator(msg.sender, pubkey, msg.value);
    }

    /**
     * @notice Increase your validators self delegation.
     *         NOTE: Only self delegations to existing validators are currently supported.
     *         If msg.sender is not a validator, the delegation will be lost.
     * @dev Proxies x/staking.MsgDelegate
     */
    function delegate(address validator) external payable {
        require(msg.value >= MinDelegation, "Staking: insufficient deposit");

        // only support self delegation for now
        require(msg.sender == validator, "Staking: only self delegation");

        emit Delegate(msg.sender, validator, msg.value);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                                  Admin                                   //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Enable the validator allowlist
     */
    function enableAllowlist() external onlyOwner {
        isAllowlistEnabled = true;
        emit AllowlistEnabled();
    }

    /**
     * @notice Disable the validator allowlist
     */
    function disableAllowlist() external onlyOwner {
        isAllowlistEnabled = false;
        emit AllowlistDisabled();
    }

    /**
     * @notice Add validators to allow list
     */
    function allowValidators(address[] calldata validators) external onlyOwner {
        for (uint256 i = 0; i < validators.length; i++) {
            isAllowedValidator[validators[i]] = true;
            emit ValidatorAllowed(validators[i]);
        }
    }

    /**
     * @notice Remove validators from allow list
     */
    function disallowValidators(address[] calldata validators) external onlyOwner {
        for (uint256 i = 0; i < validators.length; i++) {
            isAllowedValidator[validators[i]] = false;
            emit ValidatorDisallowed(validators[i]);
        }
    }
}
