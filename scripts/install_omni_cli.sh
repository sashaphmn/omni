#!/usr/bin/env bash

set -e

OS=$(uname -s)
ARCH=$(uname -m)

case $ARCH in
    arm64) ARCH="arm64" ;;
    aarch64) ARCH="arm64" ;;
    x86_64) ARCH="x86_64" ;;
    *) echo "Unsupported architecture"; exit 1 ;;
esac

URL="https://github.com/omni-network/omni/releases/latest/download/omni_${OS}_${ARCH}.tar.gz"
TARGET="$HOME/bin/omni"
SHELL_PROFILE=""
FILE=""

echo "ℹ️ Downloading omni from $URL"

# Create target directory
echo "ℹ️ Installing omni to $TARGET"
mkdir -p "$(dirname "${TARGET}")"

curl -L -s "$URL" | tar -xz -C "$(dirname "${TARGET}")" omni
chmod +x "$TARGET"

# Add to PATH if not already there
if ! echo $PATH | grep -q "$HOME/bin"; then
    if [[ $SHELL == *"zsh"* ]]; then
        FILE=".zshrc"
        SHELL_PROFILE="$HOME/$FILE"
    elif [[ $SHELL == *"bash"* ]]; then
        FILE=".bashrc"
        SHELL_PROFILE="$HOME/$FILE"
    else
        echo "Unknown shell: $SHELL. Please add $HOME/bin to your PATH manually."
        exit 1
    fi

    echo "ℹ️ Adding '\$HOME/bin' to your PATH in $SHELL_PROFILE"
    echo 'export PATH="$PATH:$HOME/bin"' >> "$SHELL_PROFILE"
    export PATH=\$PATH:$HOME/bin
fi

# Check if omni is correctly installed
if command -v omni &> /dev/null; then
    echo "✅ omni is now installed. Run first 'source ~/$FILE' and then 'omni --help' to get started."
else
    echo "❌ Error: omni executable not found in PATH. You may need to add '\$HOME/bin' to your PATH manually."
fi