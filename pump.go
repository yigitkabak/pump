package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Terminal renk kodları
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
	colorBold   = "\033[1m"
)

func main() {
	args := os.Args[1:]

	// Termux fix: Eğer args içinde kendi binary path'in çıkmışsa temizle
	if len(args) == 1 && strings.Contains(args[0], "/pump") {
		args = []string{}
	}

	if len(args) == 0 {
		printHelp()
		os.Exit(0)
	}

	switch args[0] {
	case "version":
		fmt.Println("Pump v0.1.0")

	case "help":
		printHelp()

	case "install", "i":
		if len(args) < 2 {
			fmt.Printf("%s❌ Hata: Kurulacak paket ismini belirtmelisin.%s\n", colorRed, colorReset)
			os.Exit(1)
		}
		packageToInstall := args[1]
		installPackage(packageToInstall)

	case "mod":
		installFromModFile()

	case "init":
		createModFile()

	default:
		fmt.Printf("%s❌ Geçersiz komut \"%s\"! Yardım için: pump help%s\n", colorRed, args[0], colorReset)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Printf("\n%s%s📖 PUMP KOMUT REHBERİ%s\n\n", colorBold, colorBlue, colorReset)
	fmt.Printf("%s ➜ %spump install <paket>%s %s# Belirtilen npm paketini kurar.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
	fmt.Printf("%s ➜ %spump i <paket>%s %s# install için kısa yazım.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
	fmt.Printf("%s ➜ %spump mod%s %s# mod.npr dosyasındaki tüm paketleri kurar.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
	fmt.Printf("%s ➜ %spump init%s %s# mod.npr dosyası oluşturur.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
	fmt.Printf("%s ➜ %spump version%s %s# Versiyonu gösterir.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
	fmt.Printf("%s ➜ %spump help%s %s# Yardım menüsünü gösterir.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
}

func installPackage(packageName string) bool {
	fmt.Printf("%s🔍 Paket indiriliyor: %s...%s\n", colorCyan, packageName, colorReset)

	cmd := exec.Command("npm", "install", packageName, "--silent")
	cmd.Stderr = nil
	cmd.Stdout = nil

	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s❌ %s kurulamadı.%s\n", colorRed, packageName, colorReset)
		return false
	}

	fmt.Printf("%s✅ %s başarıyla kuruldu.%s\n", colorGreen, packageName, colorReset)
	return true
}

func installFromModFile() {
	modFilePath := "mod.npr"

	if _, err := os.Stat(modFilePath); os.IsNotExist(err) {
		fmt.Printf("%s❌ Hata: mod.npr dosyası bulunamadı.%s\n", colorRed, colorReset)
		fmt.Printf("Şu komutla oluşturabilirsin: %spump init%s\n", colorCyan, colorReset)
		os.Exit(1)
	}

	file, err := os.Open(modFilePath)
	if err != nil {
		fmt.Printf("%s❌ Dosya okunamadı: %s%s\n", colorRed, err.Error(), colorReset)
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
		fmt.Printf("%s❌ Dosya okuma hatası: %s%s\n", colorRed, err.Error(), colorReset)
		os.Exit(1)
	}

	if len(modules) == 0 {
		fmt.Printf("%s⚠️ mod.npr içinde kuracak modül yok.%s\n", colorYellow, colorReset)
		os.Exit(0)
	}

	fmt.Printf("%s📦 %d modül bulundu. Kurulum başlıyor...%s\n", colorBlue, len(modules), colorReset)

	successCount := 0
	failCount := 0

	for _, moduleName := range modules {
		if installPackage(moduleName) {
			successCount++
		} else {
			failCount++
		}
	}

	fmt.Printf("\n%s📊 Kurulum Özeti:%s\n", colorBold, colorReset)
	fmt.Printf("%s✅ Başarılı: %d%s\n", colorGreen, successCount, colorReset)
	if failCount > 0 {
		fmt.Printf("%s❌ Başarısız: %d%s\n", colorRed, failCount, colorReset)
	}
}

func createModFile() {
	modFilePath := "mod.npr"

	if _, err := os.Stat(modFilePath); err == nil {
		fmt.Printf("%s⚠️ mod.npr zaten mevcut.%s\n", colorYellow, colorReset)
		return
	}

	file, err := os.Create(modFilePath)
	if err != nil {
		fmt.Printf("%s❌ Dosya oluşturulamadı: %s%s\n", colorRed, err.Error(), colorReset)
		os.Exit(1)
	}
	defer file.Close()

	content := `# Pump Modül Listesi
# Her satıra bir npm paketi ekle
# Örnek:
# react
# express
`

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("%s❌ Dosyaya yazılamadı: %s%s\n", colorRed, err.Error(), colorReset)
		os.Exit(1)
	}

	fmt.Printf("%s✅ mod.npr başarıyla oluşturuldu.%s\n", colorGreen, colorReset)
}
