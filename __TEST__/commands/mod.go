package commands

import (
	"fmt"
	"os"
	"github.com/yigitkabak/pump/__TEST__/constants"
	"github.com/yigitkabak/pump/__TEST__/utils"
)

func HandleMod() {
	modules := utils.ReadModFile()
	if len(modules) == 0 {
		return
	}

	fmt.Printf("%s📦 %d modül bulundu. Kurulum başlıyor...%s\n", constants.ColorBlue, len(modules), constants.ColorReset)

	successCount := 0
	failCount := 0

	for _, module := range modules {
		if installPackage(module) {
			successCount++
		} else {
			failCount++
		}
	}

	printSummary(successCount, failCount)
}

func printSummary(success, fail int) {
	fmt.Printf("\n%s📊 Kurulum Özeti:%s\n", constants.ColorBold, constants.ColorReset)
	fmt.Printf("%s✅ Başarılı: %d%s\n", constants.ColorGreen, success, constants.ColorReset)
	if fail > 0 {
		fmt.Printf("%s❌ Başarısız: %d%s\n", constants.ColorRed, fail, constants.ColorReset)
	}
}