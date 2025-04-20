package commands

import (
	"fmt"
	"github.com/yigitkabak/pump/__TEST__/constants"
)

func PrintHelp() {
	fmt.Printf("\n%s%s📖 PUMP KOMUT REHBERİ%s\n\n", constants.ColorBold, constants.ColorBlue, constants.ColorReset)
	fmt.Printf("%s ➜ %spump install <paket>%s %s# Belirtilen npm paketini kurar.%s\n", 
		constants.ColorYellow, constants.ColorCyan, constants.ColorReset, constants.ColorGray, constants.ColorReset)
	// Diğer yardım satırları...
}

func PrintInvalidCommand(cmd string) {
	fmt.Printf("%s❌ Geçersiz komut \"%s\"! Yardım için: pump help%s\n", 
		constants.ColorRed, cmd, constants.ColorReset)
	os.Exit(1)
}
