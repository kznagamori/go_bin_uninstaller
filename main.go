package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define flags
	listFlag := flag.Bool("list", false, "List all installed Go binaries")
	listFlagShort := flag.Bool("l", false, "List all installed Go binaries (shorthand)")
	flag.Parse()

	args := flag.Args()

	// Display help if no arguments are provided
	if len(args) == 0 && !*listFlag && !*listFlagShort {
		printHelp()
		return
	}

	// Determine the Go bin path
	gobinPath := getGoBinPath()

	if *listFlag || *listFlagShort {
		// List all binaries in the Go bin path
		files, err := os.ReadDir(gobinPath)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			return
		}

		for _, file := range files {
			if !file.IsDir() {
				fmt.Println(file.Name())
			}
		}
	} else {
		// Remove specified binaries
		for _, binary := range args {
			if binary == "help" || binary == "-h" || binary == "--help" {
				printHelp()
				return
			}
			binaryPath := filepath.Join(gobinPath, binary)
			err := os.Remove(binaryPath)
			if err != nil {
				fmt.Printf("Error removing file %s: %v\n", binary, err)
			} else {
				fmt.Printf("Removed %s\n", binary)
			}
		}
	}
}

func getGoBinPath() string {
	gobinPath := os.Getenv("GOBIN")
	if gobinPath == "" {
		gobinPath = filepath.Join(os.Getenv("GOPATH"), "bin")
		if _, err := os.Stat(gobinPath); os.IsNotExist(err) {
			ex, err := os.Executable()
			if err != nil {
				fmt.Printf("Error getting executable path: %v\n", err)
				os.Exit(1)
			}
			gobinPath = filepath.Dir(ex)
		}
	}
	return gobinPath
}

func printHelp() {
	fmt.Println("A tool to list and uninstall Go binaries")
	fmt.Println("\nUsage:")
	fmt.Println("  go_bin_uninstaller [binaries...] [flags]")
	fmt.Println("\nFlags:")
	fmt.Println("  -h, --help   Show help")
	fmt.Println("  -l, --list   List all installed Go binaries")
}
