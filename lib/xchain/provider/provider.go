// Package provider is the implementation of the Provider interface.
// it abstracts connecting to different rollup chains and collecting
// XMsgs and XReceipts to construct XBlock and deliver them to the
// calling module.
package provider

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/stream"
	"github.com/omni-network/omni/lib/xchain"
)

var _ xchain.Provider = (*Provider)(nil)

// Provider stores the source chain configuration and the global quit channel.
type Provider struct {
	network     netconf.Network
	rpcClients  map[uint64]ethclient.Client // store config for every chain ID
	backoffFunc func(context.Context) (func(), func())
}

// New instantiates the provider instance which will be ready to accept
// subscriptions for respective destination XBlocks.
func New(network netconf.Network, rpcClients map[uint64]ethclient.Client) *Provider {
	backoffFunc := func(ctx context.Context) (func(), func()) {
		return expbackoff.NewWithReset(ctx, expbackoff.WithFastConfig())
	}

	return &Provider{
		network:     network,
		rpcClients:  rpcClients,
		backoffFunc: backoffFunc,
	}
}

// StreamAsync starts a goroutine that streams xblocks asynchronously forever.
// It returns immediately. It only returns an error if the chainID in invalid.
// This is the async version of StreamBlocks.
// It retries forever (with backoff) on all fetch and callback errors.
func (p *Provider) StreamAsync(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	callback xchain.ProviderCallback,
) error {
	if _, _, err := p.getChain(chainID); err != nil {
		return err
	}

	go func() {
		err := p.stream(ctx, chainID, fromHeight, callback, true)
		if err != nil { // RetryCallback==true so this should only ever return nil on ctx cancel.
			log.Error(ctx, "Streaming xprovider blocks failed unexpectedly [BUG]", err)
		}
	}()

	return nil
}

// StreamBlocks blocks, streaming all xblocks from the chain as they become available (finalized).
// It retries forever (with backoff) on all fetch errors. It however returns the first callback error.
// It returns nil when the context is canceled.
func (p *Provider) StreamBlocks(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	callback xchain.ProviderCallback,
) error {
	return p.stream(ctx, chainID, fromHeight, callback, false)
}

func (p *Provider) stream(
	ctx context.Context,
	chainID uint64,
	fromHeight uint64,
	callback xchain.ProviderCallback,
	retryCallback bool,
) error {
	// retrieve the respective config
	chain, _, err := p.getChain(chainID)
	if err != nil {
		return err
	}

	deps := stream.Deps[xchain.Block]{
		FetchBatch: func(ctx context.Context, chainID uint64, height uint64) ([]xchain.Block, error) {
			xBlock, exists, err := p.GetBlock(ctx, chainID, height)
			if err != nil {
				return nil, err
			} else if !exists {
				return nil, nil
			}

			return []xchain.Block{xBlock}, nil
		},
		Backoff:       p.backoffFunc,
		ElemLabel:     "attestation",
		RetryCallback: retryCallback,
		Verify: func(ctx context.Context, block xchain.Block, h uint64) error {
			if block.SourceChainID != chainID {
				return errors.New("invalid block source chain id")
			} else if block.BlockHeight != h {
				return errors.New("invalid block height")
			}

			return nil
		},
		IncFetchErr: func() {
			fetchErrTotal.WithLabelValues(chain.Name).Inc()
		},
		IncCallbackErr: func() {
			callbackErrTotal.WithLabelValues(chain.Name).Inc()
		},
		SetStreamHeight: func(h uint64) {
			streamHeight.WithLabelValues(chain.Name).Set(float64(h))
		},
	}

	cb := (stream.Callback[xchain.Block])(callback)

	// Start streaming from chain's deploy height as per config.
	if fromHeight < chain.DeployHeight {
		fromHeight = chain.DeployHeight
	}

	ctx = log.WithCtx(ctx, "chain", chain.Name)
	log.Info(ctx, "Streaming xprovider blocks", "from_height", fromHeight)

	return stream.Stream(ctx, deps, chainID, fromHeight, cb)
}

// getChain provides the configuration of the given chainID.
func (p *Provider) getChain(chainID uint64) (netconf.Chain, ethclient.Client, error) {
	chain, ok := p.network.Chain(chainID)
	if !ok {
		return netconf.Chain{}, nil, errors.New("unknown chain ID for network")
	}

	client, ok := p.rpcClients[chainID]
	if !ok {
		return netconf.Chain{}, nil, errors.New("no rpc client for chain ID")
	}

	return chain, client, nil
}
