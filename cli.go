package main

import (
	"flag"
	"fmt"
)

func main() {
	standingsFlag := flag.Bool("standings", false, "Prints current standings")
	shortStandingsFlag := flag.Bool("s", false, "Prints current standings")

	flag.Usage = func() {
		fmt.Println("Welcome. This tool simply prints the current standings of the English Premier League.")
		fmt.Println("Options:")
		fmt.Printf("    %-25s%s\n", "-s, --standings", "Prints current standings")
	}

	flag.Parse()

	if *standingsFlag || *shortStandingsFlag {
		standings()
	} else {
		flag.Usage()
	}
}
