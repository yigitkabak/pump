#!/bin/bash

# Terminal renk kodları
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # Renk sıfırlama
BOLD='\033[1m'

echo -e "${BLUE}${BOLD}========================================${NC}"
echo -e "${BLUE}${BOLD}    Pump CLI Kurulum Script'i           ${NC}"
echo -e "${BLUE}${BOLD}========================================${NC}"

# Go kurulumu kontrol
if ! command -v go &> /dev/null; then
    echo -e "${RED}Hata: Sistemde Go kurulu değil.${NC}"
    echo -e "Lütfen önce Go kur: https://golang.org/doc/install"
    exit 1
fi

echo -e "${GREEN}✓ Go yüklü${NC}"

# Geçici klasör oluştur
TEMP_DIR=$(mktemp -d)
echo -e "${YELLOW}Geçici dizin oluşturuldu: ${TEMP_DIR}${NC}"

cd "$TEMP_DIR" || exit 1

# Pump reposunu klonla
echo -e "${YELLOW}Pump CLI GitHub reposu klonlanıyor...${NC}"
git clone https://github.com/yigitkabak/pump.git pump
cd pump || exit 1
echo -e "${GREEN}✓ Repo klonlandı${NC}"

# go.mod dosyası var mı kontrol et
if [ ! -f "go.mod" ]; then
    echo -e "${YELLOW}go.mod oluşturuluyor...${NC}"
    go mod init pump
    go mod tidy
    echo -e "${GREEN}✓ go.mod oluşturuldu${NC}"
else
    echo -e "${GREEN}✓ go.mod zaten mevcut${NC}"
fi

# Pump CLI'yi build et
echo -e "${YELLOW}Pump CLI derleniyor...${NC}"
go build -o pump main.go
echo -e "${GREEN}✓ Derleme tamamlandı${NC}"

# Kurulum dizinini belirle
INSTALL_DIR=""
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
else
    INSTALL_DIR="$HOME/.local/bin"
    mkdir -p "$INSTALL_DIR"
    
    # Otomatik PATH'e ekle
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        echo -e "${YELLOW}PATH'e $INSTALL_DIR ekleniyor...${NC}"
        echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.profile"
        export PATH="$HOME/.local/bin:$PATH"
        echo -e "${GREEN}✓ PATH güncellendi ve geçerli oturuma eklendi${NC}"
    fi
fi

# Binary'i taşı
echo -e "${YELLOW}Pump CLI $INSTALL_DIR dizinine yükleniyor...${NC}"
cp pump "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/pump"
echo -e "${GREEN}✓ Yüklendi: $INSTALL_DIR/pump${NC}"

# Geçici dosyaları temizle
echo -e "${YELLOW}Geçici dosyalar siliniyor...${NC}"
cd "$HOME" || exit
rm -rf "$TEMP_DIR"
echo -e "${GREEN}✓ Temizlik tamamlandı${NC}"

# Bilgilendirme
echo -e "\n${GREEN}${BOLD}=== Kurulum Tamamlandı! ===${NC}"
echo -e "Artık ${BOLD}pump${NC} komutunu istediğin yerden kullanabilirsin."
echo -e "\n${YELLOW}Örnek komutlar:${NC}"
echo -e "  pump help"
echo -e "  pump init"
echo -e "  pump install express"

echo -e "\n${BLUE}${BOLD}Keyifli kodlamalar! 🚀${NC}"
