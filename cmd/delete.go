package cmd

import (
	"fmt"
	"os"
)

func DeleteFile(fileName string) error {
	// Check if the file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return fmt.Errorf("File %s does not exist\n", fileName)
	}

	// Attempt to remove the file
	err := os.Remove(fileName)
	if err != nil {
		return fmt.Errorf("error deleting file: %w", err)
	}

	fmt.Printf("File %s was deleted\n", fileName)
	return nil
}
