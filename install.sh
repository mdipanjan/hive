#!/bin/bash

set -e

REPO="mdipanjan/hive"
BINARY="hive"
INSTALL_DIR="/usr/local/bin"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
    darwin|linux) ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Get latest version
VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | cut -d '"' -f 4)

if [ -z "$VERSION" ]; then
    echo "Failed to get latest version"
    exit 1
fi

echo "Installing $BINARY $VERSION for $OS/$ARCH..."

# Download
URL="https://github.com/$REPO/releases/download/$VERSION/${BINARY}_${OS}_${ARCH}.tar.gz"
TMP_DIR=$(mktemp -d)

curl -sL "$URL" -o "$TMP_DIR/$BINARY.tar.gz"
tar -xzf "$TMP_DIR/$BINARY.tar.gz" -C "$TMP_DIR"

# Install
if [ -w "$INSTALL_DIR" ]; then
    mv "$TMP_DIR/$BINARY" "$INSTALL_DIR/"
else
    sudo mv "$TMP_DIR/$BINARY" "$INSTALL_DIR/"
fi

# Cleanup
rm -rf "$TMP_DIR"

echo "Installed $BINARY to $INSTALL_DIR/$BINARY"
echo "Run 'hive' to start"
