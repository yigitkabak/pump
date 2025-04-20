package main

import (
	"os"
	"strings"

	"pump/src"
)

func main() {
	args := os.Args[1:]

	// Termux fix
	if len(args) == 1 && strings.Contains(args[0], "/pump") {
		args = []string{}
	}

	if len(args) == 0 {
		src.PrintHelp()
		os.Exit(0)
	}

	switch args[0] {
	case "version":
		println("Pump v0.1.0")

	case "help":
		src.PrintHelp()

	case "install", "i":
		if len(args) < 2 {
			src.PrintError("Kurulacak paket ismini belirtmelisin.")
			os.Exit(1)
		}
		src.InstallPackage(args[1])

	case "mod":
		src.InstallFromModFile()

	case "init":
		src.CreateModFile()

	default:
		src.PrintError("Geçersiz komut \"" + args[0] + "\"! Yardım için: pump help")
		os.Exit(1)
	}
}
