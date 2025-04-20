#!/bin/bash

# Color codes for terminal output
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color
BOLD='\033[1m'

echo -e "${BLUE}${BOLD}========================================${NC}"
echo -e "${BLUE}${BOLD}        Pump CLI Kurulum Scripti        ${NC}"
echo -e "${BLUE}${BOLD}========================================${NC}"

# Go kurulumu kontrolü
if ! command -v go &> /dev/null; then
    echo -e "${RED}Hata: Go yüklü değil.${NC}"
    echo -e "Go yüklemek için: https://golang.org/doc/install"
    exit 1
fi

echo -e "${GREEN}✓ Go yüklü${NC}"

# Geçici dizin oluştur
TEMP_DIR=$(mktemp -d)
echo -e "${YELLOW}Geçici dizin oluşturuldu: ${TEMP_DIR}${NC}"

cd "$TEMP_DIR" || exit 1

# Repo klonla
echo -e "${YELLOW}Pump CLI reposu klonlanıyor...${NC}"
git clone https://github.com/yigitkabak/pump.git pump
cd pump/__TEST__ || exit 1
echo -e "${GREEN}✓ Repo klonlandı${NC}"

# Go modülü kontrol et
if [ ! -f "go.mod" ]; then
    echo -e "${YELLOW}Go modülü başlatılıyor...${NC}"
    go mod init pump
    echo -e "${GREEN}✓ go.mod dosyası oluşturuldu${NC}"
else
    echo -e "${GREEN}✓ go.mod zaten mevcut${NC}"
fi

# Bağımlılıkları çek
echo -e "${YELLOW}Bağımlılıklar yükleniyor...${NC}"
go mod tidy
echo -e "${GREEN}✓ Bağımlılıklar tamamlandı${NC}"

# CLI derlemesi
echo -e "${YELLOW}Pump CLI derleniyor...${NC}"
go build -o pump main.go
echo -e "${GREEN}✓ Derleme tamamlandı${NC}"

# Kurulum dizini belirle
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
else
    INSTALL_DIR="$HOME/.local/bin"
    mkdir -p "$INSTALL_DIR"

    if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
        echo -e "${YELLOW}$INSTALL_DIR dizini PATH'e ekleniyor${NC}"
        SHELL_RC="$HOME/.bashrc"
        [ -f "$HOME/.zshrc" ] && SHELL_RC="$HOME/.zshrc"

        echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$SHELL_RC"
        echo -e "${YELLOW}$SHELL_RC dosyasına eklendi. Aktif etmek için:${NC}"
        echo -e "  source $SHELL_RC"
    fi
fi

# Binary'i taşı
echo -e "${YELLOW}Pump CLI $INSTALL_DIR dizinine kopyalanıyor...${NC}"
cp pump "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/pump"
echo -e "${GREEN}✓ Kurulum tamamlandı${NC}"

# Temizleme
echo -e "${YELLOW}Geçici dosyalar siliniyor...${NC}"
cd "$HOME" || exit
rm -rf "$TEMP_DIR"
echo -e "${GREEN}✓ Temizlik tamamlandı${NC}"

# Bilgilendirme
echo -e "\n${GREEN}${BOLD}=== Kurulum Tamamlandı! ===${NC}"
echo -e "Artık terminalde ${BOLD}pump${NC} komutunu kullanabilirsin."
echo -e "\n${YELLOW}Örnek komutlar:${NC}"
echo -e "  pump help"
echo -e "  pump init"
echo -e "  pump install react"

if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo -e "\n${YELLOW}UYARI: Bu oturumda komutu kullanmak için şunu çalıştırman gerekebilir:${NC}"
    echo -e "  export PATH=\"$INSTALL_DIR:\$PATH\""
fi

echo -e "\n${BLUE}${BOLD}İyi kodlamalar! 🚀${NC}"
