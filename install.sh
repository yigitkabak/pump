#!/bin/bash

# Color codes for terminal output
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color
BOLD='\033[1m'

echo -e "${BLUE}${BOLD}========================================${NC}"
echo -e "${BLUE}${BOLD}    Pump CLI Installation Script        ${NC}"
echo -e "${BLUE}${BOLD}        (.pmp Mod File Support)         ${NC}"
echo -e "${BLUE}${BOLD}========================================${NC}"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed on your system.${NC}"
    echo -e "Please install Go first: https://golang.org/doc/install"
    exit 1
fi

echo -e "${GREEN}✓ Go is installed${NC}"

# Create a temporary directory
TEMP_DIR=$(mktemp -d)
echo -e "${YELLOW}Working in temporary directory: ${TEMP_DIR}${NC}"

cd "$TEMP_DIR" || exit 1

# Clone the Pump CLI GitHub repository
echo -e "${YELLOW}Cloning Pump CLI repository...${NC}"
git clone https://github.com/yigitkabak/pump.git pump
cd pump || exit 1
echo -e "${GREEN}✓ Repository cloned${NC}"

# Initialize Go module (if needed)
if [ ! -f "go.mod" ]; then
    echo -e "${YELLOW}Initializing Go module...${NC}"
    go mod init pump
    echo -e "${GREEN}✓ go.mod created${NC}"
else
    echo -e "${GREEN}✓ go.mod already exists${NC}"
fi

# Build the binary
echo -e "${YELLOW}Building Pump CLI...${NC}"
go build -o pump
echo -e "${GREEN}✓ Built pump binary${NC}"

# Determine install location
INSTALL_DIR=""
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
else
    INSTALL_DIR="$HOME/.local/bin"
    mkdir -p "$INSTALL_DIR"
    
    if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
        echo -e "${YELLOW}Adding $INSTALL_DIR to your PATH${NC}"
        if [ -f "$HOME/.bashrc" ]; then
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.bashrc"
            echo -e "${YELLOW}Added to .bashrc, you may need to run: source ~/.bashrc${NC}"
        elif [ -f "$HOME/.zshrc" ]; then
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.zshrc"
            echo -e "${YELLOW}Added to .zshrc, you may need to run: source ~/.zshrc${NC}"
        else
            echo -e "${RED}Couldn't find .bashrc or .zshrc to update PATH${NC}"
            echo -e "${YELLOW}Please add this to your shell configuration:${NC}"
            echo -e "export PATH=\"$HOME/.local/bin:\$PATH\""
        fi
    fi
fi

# Install the binary
echo -e "${YELLOW}Installing Pump CLI to $INSTALL_DIR...${NC}"
cp pump "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/pump"
echo -e "${GREEN}✓ Installed pump to $INSTALL_DIR${NC}"

# Clean up
echo -e "${YELLOW}Cleaning up temporary files...${NC}"
cd "$HOME"
rm -rf "$TEMP_DIR"
echo -e "${GREEN}✓ Cleaned up temporary files${NC}"

# Final message
echo -e "\n${GREEN}${BOLD}=== Installation Complete! ===${NC}"
echo -e "You can now use the ${BOLD}pump${NC} command from anywhere."
echo -e "\n${YELLOW}Test it out with:${NC}"
echo -e "  pump help"
echo -e "  pump init   ${GRAY}(creates mod.pmp file)${NC}"
echo -e "  pump mod    ${GRAY}(installs from mod.pmp)${NC}"
echo -e "  pump install <paket>"

if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo -e "\n${YELLOW}NOTE: You may need to restart your terminal or run:${NC}"
    echo -e "  export PATH=\"$INSTALL_DIR:\$PATH\""
    echo -e "to use the ${BOLD}pump${NC} command in this session."
fi

echo -e "\n${BLUE}${BOLD}Happy coding with .pmp! 🚀${NC}"
