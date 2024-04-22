package main

import (
	"flag"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"os"
	"strings"

	"github.com/digitalghost-dev/pl-cli/subcommands"
)

var red = lipgloss.Color("#F2055C")
var styleUnderline = lipgloss.NewStyle().Underline(true)
var styleBold = lipgloss.NewStyle().Bold(true)

func main() {

	var errorColor = lipgloss.NewStyle().Foreground(red)
	standings, championsFlag, shortChampionsFlag := setupStandingsFlagSet()

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
	if os.Args[1] != "standings" {
		flag.Usage()
		fmt.Println(errorColor.Render("error: not a valid command\n"))
		os.Exit(1)
	}

	// Standings subcommand
	if os.Args[1] == "standings" {
		err := standings.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("error parsing flags: %v", err)
		}

		if len(os.Args) > 2 && !strings.HasPrefix(os.Args[2], "-") {
			fmt.Println(errorColor.Render("error: only flags are allowed after the standings subcommand\n"))
			os.Exit(1)
		}

		if *championsFlag || *shortChampionsFlag {
			records, err := subcommands.FetchAndParseCSV("https://storage.googleapis.com/premier_league_bucket/standings.csv")
			if err != nil {
				fmt.Printf("error fetching CSV: %v", err)
				return
			}

			if err := subcommands.ChampionsFlag(records); err != nil {
				fmt.Printf("error processing champions: %v", err)
				return
			}
			return
		}

		records, err := subcommands.FetchAndParseCSV("https://storage.googleapis.com/premier_league_bucket/standings.csv")
		if err != nil {
			fmt.Printf("error fetching CSV: %v", err)
			return
		}

		if err := subcommands.SubcommandStandings(records); err != nil {
			fmt.Printf("error displaying standings: %v", err)
		}
	}
}

func setupStandingsFlagSet() (*flag.FlagSet, *bool, *bool) {
	//var colorSwitch = lipgloss.AdaptiveColor{Light: "#38003C", Dark: "#CCCCCC"}
	//var colorSwitchStyle = lipgloss.NewStyle().Foreground(colorSwitch)

	standings := flag.NewFlagSet("standings", flag.ExitOnError)

	championsFlag := standings.Bool("champions", false, "Shows teams currently promoted to the Champions League.")
	shortChampionsFlag := standings.Bool("c", false, "Shows teams currently promoted to the Champions League.")

	standings.Usage = func() {
		fmt.Println("\nDisplays the current standings for the league.")

		fmt.Print("\n")
		fmt.Println(styleBold.Render("USAGE:"))
		fmt.Println("\t", "pl-cli", styleUnderline.Render("standings"), "[flag]")

		fmt.Print("\n")
		fmt.Println(styleBold.Render("AVAILABLE FLAGS:"))
		fmt.Println("\t", "-c, --champions", "\t", "Shows teams currently promoted to the Champions League.")
		fmt.Print("\n")
	}

	return standings, championsFlag, shortChampionsFlag
}
