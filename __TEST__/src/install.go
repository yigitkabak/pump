package src

import (
	"fmt"
	"os/exec"
)

func InstallPackage(packageName string) bool {
	fmt.Printf("%s🔍 Paket indiriliyor: %s...%s\n", Cyan, packageName, Reset)

	cmd := exec.Command("npm", "install", packageName, "--silent")
	cmd.Stderr = nil
	cmd.Stdout = nil

	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s❌ %s kurulamadı.%s\n", Red, packageName, Reset)
		return false
	}

	fmt.Printf("%s✅ %s başarıyla kuruldu.%s\n", Green, packageName, Reset)
	return true
}
