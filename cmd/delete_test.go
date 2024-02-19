package cmd

import (
	"os"
	"testing"
)

func TestDeleteFile(t *testing.T) {
	// create temp file
	file, err := os.CreateTemp("", "tempfile")
	if err != nil {
		t.Errorf("Error creating temp file: %v", err)
	}
	defer os.Remove(file.Name())

	// write to file
	_, err = file.Write([]byte{1, 2, 3, 4})
	if err != nil {
		t.Errorf("Error writing to temp file: %v", err)
	}

	// check if file exists
	if _, err := os.Stat(file.Name()); os.IsNotExist(err) {
		t.Errorf("File %s does not exist", file.Name())
	}

	// call the function
	err = DeleteFile(file.Name())
	if err != nil {
		t.Errorf("Error deleting file: %v", err)
	}

	// check if file was deleted
	if _, err := os.Stat(file.Name()); !os.IsNotExist(err) {
		t.Errorf("File %s was not deleted", file.Name())
	}
}
