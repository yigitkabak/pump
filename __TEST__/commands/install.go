package commands

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/yigitkabak/pump/__TEST__/constants"
)

func HandleInstall(args []string) {
	if len(args) < 2 {
		fmt.Printf("%s❌ Hata: Kurulacak paket ismini belirtmelisin.%s\n", constants.ColorRed, constants.ColorReset)
		os.Exit(1)
	}
	packageToInstall := args[1]
	success := installPackage(packageToInstall)
	
	if !success {
		os.Exit(1)
	}
}

func installPackage(packageName string) bool {
	fmt.Printf("%s🔍 Paket indiriliyor: %s...%s\n", constants.ColorCyan, packageName, constants.ColorReset)

	cmd := exec.Command("npm", "install", packageName, "--silent")
	cmd.Stderr = nil
	cmd.Stdout = nil

	if err := cmd.Run(); err != nil {
		fmt.Printf("%s❌ %s kurulamadı.%s\n", constants.ColorRed, packageName, constants.ColorReset)
		return false
	}

	fmt.Printf("%s✅ %s başarıyla kuruldu.%s\n", constants.ColorGreen, packageName, constants.ColorReset)
	return true
}