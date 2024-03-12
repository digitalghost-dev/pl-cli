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

	flag.Usage = func() {
		fmt.Println("Welcome! This tool simply prints the current standings of the English Premier League.")

		fmt.Print("\n")
		fmt.Println("Usage:")
		fmt.Println("  pl-cli [subcommand]")

		fmt.Print("\n")
		fmt.Println("Available Commands:")
		fmt.Println("\t", "standings:", "\t", "Renders the current standings of the English Premier League.")
		fmt.Print("\n")
	}

	flag.Parse()

	if len(os.Args) > 2 {
		flag.Usage()
		fmt.Println("Error: Too many arguments.")
		os.Exit(1)
	}

	if os.Args[1] == "standings" {
		if err := standings.Parse(os.Args); err != nil {
			fmt.Printf("error parsing standings: %v\n", err)
		}

		if err := cmd.DisplayStandings("https://storage.googleapis.com/premier_league_bucket/standings.csv"); err != nil {
			fmt.Printf("error displaying standings: %v", err)
		}
	}

}
