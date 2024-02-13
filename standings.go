package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rocketlaunchr/dataframe-go/imports"
)

var ctx = context.Background()

func standings() {
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
			// Handle error if needed
		}
	}(file)

	// Load the dataframe from the CSV file
	df, err := imports.LoadFromCSV(ctx, file, imports.CSVLoadOptions{InferDataTypes: true})
	if err != nil {
		panic(err)
	}

	fmt.Println(df.Table())
}
