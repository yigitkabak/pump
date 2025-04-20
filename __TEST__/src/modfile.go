package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const modFilePath = "mod.pmp"

func InstallFromModFile() {
	if _, err := os.Stat(modFilePath); os.IsNotExist(err) {
		PrintError("mod.pmp dosyası bulunamadı. pump init ile oluşturabilirsin.")
		os.Exit(1)
	}

	file, err := os.Open(modFilePath)
	if err != nil {
		PrintError("Dosya okunamadı: " + err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var modules []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			modules = append(modules, line)
		}
	}

	if err := scanner.Err(); err != nil {
		PrintError("Dosya okuma hatası: " + err.Error())
		os.Exit(1)
	}

	if len(modules) == 0 {
		fmt.Printf("%s⚠️ mod.pmp içinde kuracak modül yok.%s\n", Yellow, Reset)
		os.Exit(0)
	}

	fmt.Printf("%s📦 %d modül bulundu. Kurulum başlıyor...%s\n", Blue, len(modules), Reset)

	success, fail := 0, 0
	for _, moduleName := range modules {
		if InstallPackage(moduleName) {
			success++
		} else {
			fail++
		}
	}

	fmt.Printf("\n%s📊 Kurulum Özeti:%s\n", Bold, Reset)
	fmt.Printf("%s✅ Başarılı: %d%s\n", Green, success, Reset)
	if fail > 0 {
		fmt.Printf("%s❌ Başarısız: %d%s\n", Red, fail, Reset)
	}
}

func CreateModFile() {
	if _, err := os.Stat(modFilePath); err == nil {
		fmt.Printf("%s⚠️ mod.pmp zaten mevcut.%s\n", Yellow, Reset)
		return
	}

	file, err := os.Create(modFilePath)
	if err != nil {
		PrintError("Dosya oluşturulamadı: " + err.Error())
		os.Exit(1)
	}
	defer file.Close()

	content := `# Pump Modül Listesi
# Her satıra bir npm paketi ekle
# Örnek:
# react
# express
`
	file.WriteString(content)
	fmt.Printf("%s✅ mod.pmp başarıyla oluşturuldu.%s\n", Green, Reset)
}
