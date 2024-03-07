package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/rocketlaunchr/dataframe-go/imports"
)

func DisplayStandings(fileName string, ctx context.Context) error {
	fmt.Println("Printing standings...")

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("File does not exist. Use the -u flag to download a fresh copy.")
	}

	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	df, err := imports.LoadFromCSV(ctx, file, imports.CSVLoadOptions{InferDataTypes: true})
	if err != nil {
		return fmt.Errorf("error loading dataframe from CSV: %w", err)
	}

	fmt.Println(df.Table())
	return nil
}
