package main

import (
	"os"
	"strings"
	"github.com/yigitkabak/pump/__TEST__/commands"
)

func main() {
	args := handleTermuxArgs(os.Args[1:])

	if len(args) == 0 {
		commands.PrintHelp()
		os.Exit(0)
	}

	switch args[0] {
	case "version":
		commands.HandleVersion()
	case "help":
		commands.PrintHelp()
	case "install", "i":
		commands.HandleInstall(args)
	case "mod":
		commands.HandleMod()
	case "init":
		commands.HandleInit()
	default:
		commands.PrintInvalidCommand(args[0])
	}
}

func handleTermuxArgs(args []string) []string {
	if len(args) == 1 && strings.Contains(args[0], "/pump") {
		return []string{}
	}
	return args
}
