package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/rocketlaunchr/dataframe-go/imports"
)

var ctx = context.Background()

func main() {
	standingsFlag := flag.Bool("standings", false, "Prints current standings")
	shortStandingsFlag := flag.Bool("s", false, "Prints current standings")

	flag.Usage = func() {
		fmt.Println("Options:")
		fmt.Println("    -s                    Prints current standings")
		fmt.Println("    --standings           Prints current standings")
	}

	flag.Parse()

	if *standingsFlag || *shortStandingsFlag {
		fmt.Println("Printing standings...")

		csvFilePath := "standings.csv"

		// Open the CSV file
		file, err := os.Open(csvFilePath)
		if err != nil {
			panic(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				// Handle the error if needed
			}
		}(file)

		// Load the dataframe from the CSV file
		df, err := imports.LoadFromCSV(ctx, file, imports.CSVLoadOptions{InferDataTypes: true})
		if err != nil {
			panic(err)
		}

		fmt.Println(df.Table())
	} else {
		flag.Usage()
	}
}
