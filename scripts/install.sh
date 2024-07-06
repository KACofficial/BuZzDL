#!/bin/bash
# Install script for BuZzDL
# Author: @KACofficial

# Ensure we exit on first error and all commands are run as intended
set -euo pipefail

# Run Go mod tidy to clean up the go.mod and go.sum files
echo "Tidying up Go modules..."
go mod tidy

# Build the Go binary with optimizations
echo "Building BuZzDL binary..."
sudo go build -ldflags="-w -s" -o /usr/local/bin/buzzdl

# Ensure the binary is executable
echo "Setting execute permissions for BuZzDL..."
sudo chmod +x /usr/local/bin/buzzdl

# Confirm installation success to the user
echo "BuZzDL installed successfully!"
echo "To use BuZzDL, type 'buzzdl' in your terminal."

echo "Done!"