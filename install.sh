#!/bin/sh
#
# Install script for lock (Go binary)
#

set -e

INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
REPO="ronthekiehn/lock"
TMP_DIR="$(mktemp -d)"

cleanup() {
    rm -rf "$TMP_DIR"
}
trap cleanup EXIT

echo "🔒 Installing lock..."

OS="$(uname -s)"
ARCH="$(uname -m)"

if [ "$OS" != "Darwin" ]; then
    echo "❌ This installer currently supports macOS only."
    exit 1
fi

case "$ARCH" in
    arm64|aarch64)
        GOARCH="arm64"
        ;;
    x86_64|amd64)
        GOARCH="amd64"
        ;;
    *)
        echo "❌ Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

ASSET="lock_darwin_${GOARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/latest/download/${ASSET}"
ARCHIVE_PATH="$TMP_DIR/$ASSET"

echo "Downloading latest release binary..."
curl -fsSL "$URL" -o "$ARCHIVE_PATH"

echo "Extracting..."
tar -xzf "$ARCHIVE_PATH" -C "$TMP_DIR"

if [ ! -f "$TMP_DIR/lock" ]; then
    echo "❌ Could not find lock binary in release archive"
    exit 1
fi

chmod +x "$TMP_DIR/lock"

if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating $INSTALL_DIR..."
    sudo mkdir -p "$INSTALL_DIR"
fi

echo "Installing to $INSTALL_DIR (requires sudo)..."
sudo install -m 0755 "$TMP_DIR/lock" "$INSTALL_DIR/lock"

echo "✅ lock installed successfully!"
echo ""
echo "Usage:"
echo "  lock x.com"
echo "  lock -n \"ship checkout\" x.com reddit.com"
echo "  lock --status"
