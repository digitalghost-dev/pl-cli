package cmd

import (
	"context"
	"os"
	"testing"
)

func TestDisplayStandings(t *testing.T) {
	// create temp file
	file, err := os.CreateTemp("", "test_csv.csv")
	if err != nil {
		t.Errorf("Error creating temp file: %v", err)
	}

	// check if file exists
	if _, err := os.Stat(file.Name()); os.IsNotExist(err) {
		t.Errorf("File %s does not exist", file.Name())
	}

	// write to file
	_, err = file.Write([]byte("Team ID,Rank,Team,Games Played,Won,Draws,Loses,Recent Form,Points,Goals For,Goals Against,Goal Difference\n1,1,Manchester United,38,24,11,3,WWLDW,83,73,32,41\n2,2,Manchester City,38,23,13,2,WLWWD,82,79,32,47\n3,3,Liverpool,38,23,12,3,WLWDD,81,67,30,37\n4,4,Chelsea,38,22,10,6,WLLDW,76,58,36,22"))
	if err != nil {
		t.Errorf("Error writing to file: %v", err)
	}

	// call the function
	ctx := context.Background()
	err = DisplayStandings(file.Name(), ctx)
	if err != nil {
		t.Errorf("Error displaying standings: %v", err)
	}

	// remove file
	err = os.Remove(file.Name())
	if err != nil {
		t.Errorf("Error removing file: %v", err)
	}

	// check if file was deleted
	if _, err := os.Stat(file.Name()); !os.IsNotExist(err) {
		t.Errorf("File %s was not deleted", file.Name())
	}
}
