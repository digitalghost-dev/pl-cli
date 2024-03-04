package cmd

import (
	"net/http"
	"fmt"
	"os"
	"testing"
)

func TestUpdateFile(t *testing.T) {

	fileName := "testfile.csv"
	url := "https://storage.googleapis.com/premier_league_bucket/standings.csv"

	response, err := http.Get(url)
	if err != nil {
		t.Errorf("error reaching bucket: %v", err)
		return
	}
	
	defer response.Body.Close()
	
	switch response.StatusCode {
	case http.StatusOK:
		fmt.Println("Connection to bucket successful")
	case http.StatusNotFound:
		fmt.Println("404 - Not Found")
	default:
		t.Errorf("error reaching bucket: %v", err)
	}

    defer func() {
        _ = os.Remove(fileName)
    }()

    if err := UpdateFile(fileName, url); err != nil {
        t.Errorf("UpdateFile() failed: %v", err)
    }

    _, err = os.Stat(fileName)
    if os.IsNotExist(err) {
        t.Errorf("UpdateFile() did not create the file: %v", err)
    }
}