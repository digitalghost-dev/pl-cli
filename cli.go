package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/digitalghost-dev/pl-cli/subcommands"
)

func main() {
	// Standings subcommand
	standings := flag.NewFlagSet("standings", flag.ExitOnError)

	championsFlag := standings.Bool("champions", false, "Shows teams currently promoted to the Champions League.")
	shortChampionsFlag := standings.Bool("c", false, "Shows teams currently promoted to the Champions League.")

	standings.Usage = func() {
		fmt.Println("Displays the current standings for the league.")

		fmt.Print("\n")
		fmt.Println("Usage:")
		fmt.Println("  pl-cli standings [flag]")

		fmt.Print("\n")
		fmt.Println("Available Flags:")
		fmt.Println("\t", "-c, --champions", "\t", "Shows teams currently promoted to the Champions League.")
		fmt.Print("\n")
	}

	flag.Usage = func() {
		fmt.Println("Welcome! This tool displays statistics and data for English Premier League in the terminal.")
		fmt.Println("Data is currently updated every day at 6:00 PST.")

		fmt.Print("\n")
		fmt.Println("Usage:")
		fmt.Println("  pl-cli [subcommand] [flag]")

		fmt.Print("\n")
		fmt.Println("Available Commands:")
		fmt.Println("\t", "standings:", "\t", "Renders the current standings of the English Premier League.")
		fmt.Print("\n")
	}

	flag.Parse()

	if len(os.Args) > 3 {
		flag.Usage()
		fmt.Println("Too many arguments")
		os.Exit(1)
	}

	// Standings subcommand
	if os.Args[1] == "standings" {
		err := standings.Parse(os.Args[2:])
		if err != nil {
			fmt.Printf("error parsing flags: %v", err)
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
