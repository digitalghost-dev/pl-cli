package subcommands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//goland:noinspection ALL
func TestSetupStandingsFlagSet(t *testing.T) {
	standings, championsFlag, shortChampionsFlag := SetupStandingsFlagSet()

	assert.NotNil(t, standings, "standings flag set should not be nil")
	assert.Equal(t, "standings", standings.Name(), "standings flag set name should be 'standings'")

	assert.NotNil(t, championsFlag, " flag set should not be nil")
	assert.Equal(t, bool(false), *championsFlag, "champions flag name should be 'bool(false)'")
	assert.Equal(t, bool(false), *shortChampionsFlag, "champions flag name should be 'bool(false)'")
}

func TestChampionsFlag(t *testing.T) {
	mockRecords := [][]string{
		{"RANK", "TEAM", "GP", "WINS", "DRAWS", "LOSES", "RECENT FORM", "POINTS", "GF", "GA", "GD"},
		{"1", "Manchester City", "37", "27", "7", "3", "WWWWW", "88", "93", "33", "60"},
		{"2", "Arsenal", "37", "27", "5", "5", "WWWWW", "86", "89", "28", "61"},
		{"3", "Liverpool", "37", "23", "10", "4", "DWDLW", "79", "84", "41", "43"},
		{"4", "Aston Villa", "37", "20", "8", "9", "DLDWW", "68", "76", "56", "20"},
		{"5", "Tottenham", "37", "19", "6", "12", "LWLLL", "63", "71", "61", "10"},
	}

	err := ChampionsFlag(mockRecords)

	assert.Nil(t, err, "ChampionsFlag should not return an error")
}
