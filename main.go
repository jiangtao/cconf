package main

import (
	"fmt"
	"os"

	"github.com/jiangtao/ccconfig/cmd"
	"github.com/jiangtao/ccconfig/pkg/config"
)

// Version information (set by ldflags)
var (
	Version   = "dev"
	BuildTime = "unknown"
)

func main() {
	// Add version flag
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Printf("ccconfig version %s (built: %s)\n", Version, BuildTime)
		os.Exit(0)
	}

	// Initialize configuration
	if err := config.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing config: %v\n", err)
		os.Exit(1)
	}

	// Execute CLI
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
