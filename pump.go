package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Color codes for terminal output
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

	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	switch args[0] {
	case "version":
		fmt.Println("Pump v0.1.0")

	case "help":
		printHelp()

	case "install", "i":
		if len(args) < 2 {
			fmt.Printf("%s❌ Error: Please specify the package name to install.%s\n", colorRed, colorReset)
			os.Exit(1)
		}
		packageToInstall := args[1]
		installPackage(packageToInstall)

	case "mod":
		installFromModFile()

	case "init":
		createModFile()

	default:
		fmt.Printf("%s❌ Error: Invalid command! Use \"pump help\" for usage information.%s\n", colorRed, colorReset)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf("%sUsage:%s pump install <package-name>\n\n", colorYellow, colorReset)
	fmt.Printf("%sPump - A simple npm package installer.%s\n", colorBold, colorReset)
}

func printHelp() {
	fmt.Printf("\n%s%s📖 PUMP COMMAND GUIDE%s\n\n", colorBold, colorBlue, colorReset)
	fmt.Printf("%s ➜ %spump install <module>%s %s# Installs a new module.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
	fmt.Printf("%s ➜ %spump i <module>%s %s# Short for install.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
	fmt.Printf("%s ➜ %spump mod%s %s# Installs all modules from mod.npr file.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
	fmt.Printf("%s ➜ %spump init%s %s# Creates an empty mod.npr file.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
	fmt.Printf("%s ➜ %spump version%s %s# Displays the current version.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
	fmt.Printf("%s ➜ %spump help%s %s# Shows this help menu.%s\n", colorYellow, colorCyan, colorReset, colorGray, colorReset)
}

func installPackage(packageName string) bool {
	fmt.Printf("%s🔍 Downloading package: %s...%s\n", colorCyan, packageName, colorReset)

	cmd := exec.Command("npm", "install", packageName, "--silent")
	cmd.Stderr = nil // Suppress error output
	cmd.Stdout = nil // Suppress standard output

	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s❌ Installation of %s failed%s\n", colorRed, packageName, colorReset)
		return false
	}

	fmt.Printf("%s✅ %s was successfully installed.%s\n", colorGreen, packageName, colorReset)
	return true
}

func installFromModFile() {
	modFilePath := "mod.npr"
	
	// Check if mod.npr file exists
	if _, err := os.Stat(modFilePath); os.IsNotExist(err) {
		fmt.Printf("%s❌ Error: mod.npr file not found in the current directory.%s\n", colorRed, colorReset)
		fmt.Printf("Create one with: %spump init%s\n", colorCyan, colorReset)
		os.Exit(1)
	}

	// Open and read the file
	file, err := os.Open(modFilePath)
	if err != nil {
		fmt.Printf("%s❌ Error reading mod.npr file: %s%s\n", colorRed, err.Error(), colorReset)
		os.Exit(1)
	}
	defer file.Close()

	// Read file line by line
	var modules []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			modules = append(modules, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("%s❌ Error scanning mod.npr file: %s%s\n", colorRed, err.Error(), colorReset)
		os.Exit(1)
	}

	if len(modules) == 0 {
		fmt.Printf("%s⚠️ No modules found in mod.npr file.%s\n", colorYellow, colorReset)
		os.Exit(0)
	}

	fmt.Printf("%s📦 Found %d module(s) in mod.npr file. Starting installation...%s\n", colorBlue, len(modules), colorReset)

	// Install each module one by one
	successCount := 0
	failCount := 0

	for _, moduleName := range modules {
		if installPackage(moduleName) {
			successCount++
		} else {
			failCount++
		}
	}

	// Show installation summary
	fmt.Printf("\n%s📊 Installation Summary:%s\n", colorBold, colorReset)
	fmt.Printf("%s✅ Successfully installed: %d module(s)%s\n", colorGreen, successCount, colorReset)
	if failCount > 0 {
		fmt.Printf("%s❌ Failed to install: %d module(s)%s\n", colorRed, failCount, colorReset)
	}
}

func createModFile() {
	modFilePath := "mod.npr"
	
	// Check if mod.npr file already exists
	if _, err := os.Stat(modFilePath); err == nil {
		fmt.Printf("%s⚠️ mod.npr file already exists%s\n", colorYellow, colorReset)
		return
	}

	// Create the file
	file, err := os.Create(modFilePath)
	if err != nil {
		fmt.Printf("%s❌ Error creating mod.npr file: %s%s\n", colorRed, err.Error(), colorReset)
		os.Exit(1)
	}
	defer file.Close()

	// Write initial content
	content := `# Pump Module List
# Add npm packages below, one per line
# Example:
# react
# express
`
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("%s❌ Error writing to mod.npr file: %s%s\n", colorRed, err.Error(), colorReset)
		os.Exit(1)
	}

	fmt.Printf("%s✅ Created mod.npr file successfully%s\n", colorGreen, colorReset)
}
