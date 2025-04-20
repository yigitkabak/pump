package src

import "fmt"

func PrintHelp() {
	fmt.Printf("\n%s%s📖 PUMP KOMUT REHBERİ -%s\n\n", Bold, Blue, Reset)
	fmt.Printf("%s ➜ %spump install <paket>%s %s# NPM paketi kurar.%s\n", Yellow, Cyan, Reset, Gray, Reset)
	fmt.Printf("%s ➜ %spump i <paket>%s %s# Kısa yazım.%s\n", Yellow, Cyan, Reset, Gray, Reset)
	fmt.Printf("%s ➜ %spump mod%s %s# mod.pmp dosyasındaki tüm modülleri kurar.%s\n", Yellow, Cyan, Reset, Gray, Reset)
	fmt.Printf("%s ➜ %spump init%s %s# mod.pmp dosyası oluşturur.%s\n", Yellow, Cyan, Reset, Gray, Reset)
	fmt.Printf("%s ➜ %spump version%s %s# Versiyonu gösterir.%s\n", Yellow, Cyan, Reset, Gray, Reset)
	fmt.Printf("%s ➜ %spump help%s %s# Yardım menüsü.%s\n", Yellow, Cyan, Reset, Gray, Reset)
}
