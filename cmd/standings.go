package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/rocketlaunchr/dataframe-go/imports"
)

var ctx = context.Background()

func GetStandings(fileName string) {
	fmt.Println("Printing standings...")

	// Open the CSV file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Errorf("error closing file: %w", err)
		}
	}(file)

	// Load the dataframe from the CSV file
	df, err := imports.LoadFromCSV(ctx, file, imports.CSVLoadOptions{InferDataTypes: true})
	if err != nil {
		panic(err)
	}

	fmt.Println(df.Table())
}
