package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/digitalghost-dev/pl-cli/cmd"
)

func main() {
	standingsFlag := flag.Bool("standings", false, "Prints current standings")
	shortStandingsFlag := flag.Bool("s", false, "Prints current standings")

	deleteFlag := flag.Bool("delete", false, "Deletes standings.csv file")
	shortDeleteFlag := flag.Bool("d", false, "Deletes standings.csv file")

	updateFlag := flag.Bool("update", false, "Updates standings.csv file")
	shortUpdateFlag := flag.Bool("u", false, "Updates standings.csv file")

	flag.Usage = func() {
		fmt.Println("Welcome. This tool simply prints the current standings of the English Premier League.")
		fmt.Println("Options:")
		fmt.Printf("    %-25s%s\n", "-d, --delete", "Deletes standings.csv file")
		fmt.Printf("    %-25s%s\n", "-s, --standings", "Prints current standings")
	}

	flag.Parse()

	if *standingsFlag || *shortStandingsFlag {
		ctx := context.Background()
		err := cmd.DisplayStandings("standings.csv", ctx)
		if err != nil {
			return
		}
	} else if *deleteFlag || *shortDeleteFlag {
		err := cmd.DeleteFile("standings.csv")
		if err != nil {
			return
		}
	} else if *updateFlag || *shortUpdateFlag {
		err := cmd.UpdateFile("standings.csv", "https://storage.googleapis.com/premier_league_bucket/standings.csv")
		if err != nil {
			return
		}
	} else {
		flag.Usage()
	}
}
