#!/bin/bash
# countdownclilinux installer
# Usage: curl -sSL https://raw.githubusercontent.com/basheer78655/countdown/main/install.sh | sudo bash

set -e

REPO="basheer78655/countdown"
BINARY_NAME="countdownclilinux"
INSTALL_DIR="/usr/local/bin"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${CYAN}╔══════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║   ${GREEN}countdownclilinux${CYAN} installer            ║${NC}"
echo -e "${CYAN}║   Terminal countdown timer for Linux     ║${NC}"
echo -e "${CYAN}╚══════════════════════════════════════════╝${NC}"
echo ""

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
    linux)  OS="linux" ;;
    darwin) OS="darwin" ;;
    *)
        echo -e "${RED}Error: Unsupported OS: $OS${NC}"
        exit 1
        ;;
esac

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
    x86_64|amd64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *)
        echo -e "${RED}Error: Unsupported architecture: $ARCH${NC}"
        exit 1
        ;;
esac

echo -e "${YELLOW}Detected:${NC} ${OS}/${ARCH}"

# Get latest release tag
echo -e "${YELLOW}Fetching latest release...${NC}"
LATEST_TAG=$(curl -sSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')

if [ -z "$LATEST_TAG" ]; then
    echo -e "${RED}Error: Could not fetch latest release${NC}"
    exit 1
fi

echo -e "${GREEN}Latest version:${NC} ${LATEST_TAG}"

# Build download URL
ARCHIVE_NAME="${BINARY_NAME}_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${ARCHIVE_NAME}"

# Download and install
TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR" EXIT

echo -e "${YELLOW}Downloading ${ARCHIVE_NAME}...${NC}"
curl -sSL "$DOWNLOAD_URL" -o "${TEMP_DIR}/${ARCHIVE_NAME}"

echo -e "${YELLOW}Extracting...${NC}"
tar -xzf "${TEMP_DIR}/${ARCHIVE_NAME}" -C "$TEMP_DIR"

echo -e "${YELLOW}Installing to ${INSTALL_DIR}...${NC}"
install -m 755 "${TEMP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"

echo ""
echo -e "${GREEN}✓ ${BINARY_NAME} ${LATEST_TAG} installed successfully!${NC}"
echo -e "${CYAN}  Run: ${BINARY_NAME} 10s${NC}"
echo ""
