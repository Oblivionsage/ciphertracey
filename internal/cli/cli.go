package cli

import (
	"fmt"
	"strings"

	"github.com/Oblivionsage/ciphertracey/internal/log"
	"github.com/Oblivionsage/ciphertracey/internal/version"
)

// Run executes the CLI with the given arguments and returns an exit code.
func Run(args []string) int {
	if len(args) < 2 {
		printUsage()
		return 0
	}

	cmd := args[1]

	// Handle global flags
	if cmd == "--version" || cmd == "-v" {
		fmt.Println(version.String())
		return 0
	}

	if cmd == "help" || cmd == "--help" || cmd == "-h" {
		printUsage()
		return 0
	}

	// Route subcommands
	switch cmd {
	case "ingest":
		return cmdIngest(args[2:])
	case "score":
		return cmdScore(args[2:])
	case "graph":
		return cmdGraph(args[2:])
	case "export":
		return cmdExport(args[2:])
	case "run":
		return cmdRun(args[2:])
	default:
		log.Error("unknown command: %s", cmd)
		printUsage()
		return 1
	}
}

func printUsage() {
	usage := `CipherTracey - Local AML Intelligence & Analytics

Usage:
  ciphertracey [command] [flags]

Available Commands:
  ingest      Import blockchain transaction data
  score       Calculate risk scores for addresses
  graph       Visualize transaction flows
  export      Export analysis results
  run         Start interactive analysis server

Flags:
  --version, -v    Show version information
  --help, -h       Show this help message

Use "ciphertracey [command] --help" for more information about a command.
`
	fmt.Print(usage)
}

// Stub command handlers
func cmdIngest(args []string) int {
	log.Info("ingest called with args: %s", strings.Join(args, " "))
	return 0
}

func cmdScore(args []string) int {
	log.Info("score called with args: %s", strings.Join(args, " "))
	return 0
}

func cmdGraph(args []string) int {
	log.Info("graph called with args: %s", strings.Join(args, " "))
	return 0
}

func cmdExport(args []string) int {
	log.Info("export called with args: %s", strings.Join(args, " "))
	return 0
}

func cmdRun(args []string) int {
	log.Info("run called with args: %s", strings.Join(args, " "))
	return 0
}
