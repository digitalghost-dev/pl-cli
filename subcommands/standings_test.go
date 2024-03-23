package subcommands

import (
	"testing"
)

func TestFetchAndParseCSV(t *testing.T) {
	url := "https://storage.googleapis.com/premier_league_bucket/standings.csv"
	if _, err := FetchAndParseCSV(url); err != nil {
		t.Errorf("failed to fetch and parse CSV: %v", err)
	}
}

func TestSubcommandStandings(t *testing.T) {
	url := "https://storage.googleapis.com/premier_league_bucket/standings.csv"
	records, err := FetchAndParseCSV(url)
	if err != nil {
		t.Errorf("failed to fetch and parse CSV: %v", err)
		return
	}

	err = SubcommandStandings(records)
	if err != nil {
		t.Errorf("failed to generate standings: %v", err)
	}
}

func TestChampionsFlag(t *testing.T) {
	url := "https://storage.googleapis.com/premier_league_bucket/standings.csv"
	records, err := FetchAndParseCSV(url)
	if err != nil {
		t.Errorf("failed to fetch and parse CSV: %v", err)
		return
	}

	err = ChampionsFlag(records)
	if err != nil {
		t.Errorf("failed to identify ChampionsFlag teams: %v", err)
	}
}
