package cmd

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDisplayStandings(t *testing.T) {
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

}
