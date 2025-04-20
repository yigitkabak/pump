package src

import "fmt"

// Terminal renk kodları
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Gray   = "\033[90m"
	Bold   = "\033[1m"
)

func PrintError(msg string) {
	fmt.Printf("%s❌ %s%s\n", Red, msg, Reset)
}
