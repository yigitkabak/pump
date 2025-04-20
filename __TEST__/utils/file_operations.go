package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/yigitkabak/pump/__TEST__/constants"
)

func ReadModFile() []string {
	modFilePath := "mod.pmp"

	if !ModFileExists() {
		fmt.Printf("%s❌ Hata: mod.pmp dosyası bulunamadı.%s\n", constants.ColorRed, constants.ColorReset)
		fmt.Printf("Şu komutla oluşturabilirsin: %spump init%s\n", constants.ColorCyan, constants.ColorReset)
		os.Exit(1)
	}

	file, err := os.Open(modFilePath)
	if err != nil {
		fmt.Printf("%s❌ Dosya okunamadı: %s%s\n", constants.ColorRed, err.Error(), constants.ColorReset)
		os.Exit(1)
	}
	defer file.Close()

	var modules []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if isValidModule(line) {
			modules = append(modules, line)
		}
	}

	return modules
}

// Yardımcı fonksiyonlar...
func isValidModule(line string) bool {
	return line != "" && !strings.HasPrefix(line, "#")
}

func ModFileExists() bool {
	_, err := os.Stat("mod.pmp")
	return !os.IsNotExist(err)
}

func CreateModFile() error {
	content := `# Pump Modül Listesi...`
	return os.WriteFile("mod.pmp", []byte(content), 0644)
}
