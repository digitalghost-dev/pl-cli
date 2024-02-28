package main

import (
	"flag"
	"fmt"
	"github.com/digitalghost-dev/premier-league-cli/cmd"
)

func main() {
	standingsFlag := flag.Bool("standings", false, "Prints current standings")
	shortStandingsFlag := flag.Bool("s", false, "Prints current standings")

	deleteFlag := flag.Bool("delete", false, "Deletes standings.csv file")
	shortDeleteFlag := flag.Bool("d", false, "Deletes standings.csv file")

	flag.Usage = func() {
		fmt.Println("Welcome. This tool simply prints the current standings of the English Premier League.")
		fmt.Println("Options:")
		fmt.Printf("    %-25s%s\n", "-d, --delete", "Deletes standings.csv file")
		fmt.Printf("    %-25s%s\n", "-s, --standings", "Prints current standings")
	}

	flag.Parse()

	if *standingsFlag || *shortStandingsFlag {
		err := cmd.DisplayStandings("standings.csv")
		if err != nil {
			return
		}
	} else if *deleteFlag || *shortDeleteFlag {
		err := cmd.DeleteFile("standings.csv")
		if err != nil {
			return
		}
	} else {
		flag.Usage()
	}
}
