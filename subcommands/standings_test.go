package subcommands

import (
	"testing"
)

func TestSubcommandStandings(t *testing.T) {
	StandingsCommand()
}

func TestFetchAndParseCSV(t *testing.T) {
	url := "https://storage.googleapis.com/premier_league_bucket/standings.csv"
	if _, err := FetchAndParseCSV(url); err != nil {
		t.Errorf("failed to fetch and parse CSV: %v", err)
	}
}
