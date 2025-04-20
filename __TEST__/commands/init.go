package commands

import (
	"fmt"
	"os"
	"github.com/yigitkabak/pump/_TEST__/constants"
	"github.com/yigitkabak/pump/__TEST__/utils"
)

func HandleInit() {
	if utils.ModFileExists() {
		fmt.Printf("%s⚠️ mod.npr zaten mevcut.%s\n", constants.ColorYellow, constants.ColorReset)
		return
	}

	if err := utils.CreateModFile(); err != nil {
		fmt.Printf("%s❌ Hata: %v%s\n", constants.ColorRed, err, constants.ColorReset)
		os.Exit(1)
	}

	fmt.Printf("%s✅ mod.npr başarıyla oluşturuldu.%s\n", constants.ColorGreen, constants.ColorReset)
}