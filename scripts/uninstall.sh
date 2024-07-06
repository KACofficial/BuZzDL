#!/bin/bash
# Uninstall script for BuZzDL
# Author: @KACofficial

# Ensure we exit on first error and all commands are run as intended
set -euo pipefail

# Check if BuZzDL binary exists
if [ -f "/usr/local/bin/buzzdl" ]; then
    # Remove BuZzDL binary
    echo "Removing BuZzDL binary..."
    sudo rm /usr/local/bin/buzzdl
    echo "BuZzDL binary removed."
else
    echo "BuZzDL is not installed."
fi

# Clean up any remaining files or directories (optional, adjust as needed)
# echo "Cleaning up additional files if any..."

echo "BuZzDL uninstallation complete."