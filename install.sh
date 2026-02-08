#!/bin/bash
#
# Install script for lock
#

set -e

INSTALL_DIR="/usr/local/bin"
REPO_URL="https://raw.githubusercontent.com/ronthekiehn/lock/main/lock"

echo "🔒 Installing lock..."

# Check if /usr/local/bin exists
if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating $INSTALL_DIR..."
    sudo mkdir -p "$INSTALL_DIR"
fi

# Download the script
echo "Downloading lock script..."
sudo curl -fsSL "$REPO_URL" -o "$INSTALL_DIR/lock"

# Make it executable
sudo chmod +x "$INSTALL_DIR/lock"

echo "✅ lock installed successfully!"
echo ""
echo "Usage:"
echo "  lock x.com"
echo "  lock reddit.com"
echo ""
echo "Note: You may need to restart your terminal or run 'hash -r' for the command to be available."
