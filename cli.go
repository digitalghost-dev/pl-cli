package main

import (
	"flag"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/digitalghost-dev/pl-cli/subcommands"
	"os"
)

var validCommands = map[string]bool{
	"standings": true,
}

var red = lipgloss.Color("#F2055C")
var styleBold = lipgloss.NewStyle().Bold(true)
var errorColor = lipgloss.NewStyle().Foreground(red)

func main() {

	flag.Usage = func() {
		fmt.Println("\nWelcome! This tool displays statistics and data for English Premier League in the terminal.")
		fmt.Println("Data is currently updated every day at 06:00 PST.")

		fmt.Println(styleBold.Render("\nUSAGE:"))
		fmt.Println("\t", "pl-cli [flag]")
		fmt.Println("\t", "pl-cli [subcommand] [flag]")

		fmt.Println(styleBold.Render("\nAVAILABLE COMMANDS:"))
		fmt.Println("\t", "standings:", "\t", "Renders the current standings of the English Premier League.")

		fmt.Println(styleBold.Render("\nGLOBAL FLAGS:"))
		fmt.Println("\t", "-h, --help", "\t", "Shows the help menu")
		fmt.Print("\n")
	}

	flag.Parse()

	// Check the length of provided arguments
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	} else if len(os.Args) > 3 {
		fmt.Println(errorColor.Render("error: too many arguments\n"))
		os.Exit(1)
	}

	// Check whether the provided command is valid
	if !validCommands[os.Args[1]] {
		flag.Usage()
		fmt.Println(errorColor.Render("error: not a valid command\n"))
		os.Exit(1)
	}

	if os.Args[1] == "standings" {
		subcommands.StandingsCommand()
	}
}
