#!/bin/bash
#
# Install script for lock
#

set -e

# Check OS
if [[ "$(uname)" != "Darwin" ]]; then
    echo "❌ Error: lock is currently only supported on macOS"
    echo "   Detected OS: $(uname)"
    exit 1
fi

INSTALL_DIR="/usr/local/bin"
VERSION="${LOCK_VERSION:-main}"
REPO_URL="https://raw.githubusercontent.com/ronthekiehn/lock/${VERSION}/lock"
TEMP_FILE="/tmp/lock.$$"

echo "🔒 Installing lock (version: ${VERSION})..."

# Download the script as unprivileged user
echo "Downloading lock script..."
curl -fsSL "$REPO_URL" -o "$TEMP_FILE"

# Make it executable
chmod +x "$TEMP_FILE"

# Install to /usr/local/bin (requires sudo)
echo "Installing to $INSTALL_DIR (requires sudo)..."
sudo install -m 0755 "$TEMP_FILE" "$INSTALL_DIR/lock"

# Clean up
rm -f "$TEMP_FILE"

echo "✅ lock installed successfully!"
echo ""
echo "Usage:"
echo "  lock x.com"
echo "  lock reddit.com"
echo ""
echo "Note: You may need to restart your terminal or run 'hash -r' for the command to be available."
