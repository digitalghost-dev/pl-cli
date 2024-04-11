package subcommands

import (
	"encoding/csv"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"io"
	"net/http"
	"os"
)

const (
	purple = lipgloss.Color("#38003C")
	green  = lipgloss.Color("#05F26C")
	blue   = lipgloss.Color("#07F2F2")
	red    = lipgloss.Color("#F2055C")
)

var colorSwitch = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#CCCCCC"}

func FetchAndParseCSV(url string) ([][]string, error) {
	// Perform HTTP GET request to fetch the CSV file
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch CSV file: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("failed to close response body: %v", err)
		}
	}(response.Body)

	// Parse the CSV file
	reader := csv.NewReader(response.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse CSV: %v", err)
	}

	return records, nil
}

func SubcommandStandings(records [][]string) error {

	records = records[1:]

	re := lipgloss.NewRenderer(os.Stdout)

	headerStyle := re.NewStyle().Foreground(colorSwitch).Bold(true).Align(lipgloss.Center)
	BorderStyle := lipgloss.NewStyle().Foreground(purple)
	positionColors := map[string]lipgloss.Style{
		"championsLeague": lipgloss.NewStyle().Foreground(green),
		"europaLeague":    lipgloss.NewStyle().Foreground(blue),
		"relegation":      lipgloss.NewStyle().Foreground(red),
	}

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		Headers(headerStyle.String()).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			if row == 0 {
				style = headerStyle
			} else if row <= 4 {
				style = positionColors["championsLeague"]
			} else if row == 5 {
				style = positionColors["europaLeague"]
			} else if row >= 18 {
				style = positionColors["relegation"]
			} else {
				style = lipgloss.NewStyle().Foreground(colorSwitch)
			}

			// Rank
			if col == 0 {
				style = style.Copy().Width(7).Align(lipgloss.Center)
			}

			// Team
			if col == 1 {
				style = style.Copy().Width(20)
			}

			// Games Played
			if col == 2 {
				style = style.Copy().Width(7).Align(lipgloss.Center)
			}

			// Wins, Draws, Loses
			if col >= 3 && col <= 5 {
				style = style.Copy().Width(6).Align(lipgloss.Center)
			}

			// Recent Form
			if col == 6 {
				style = style.Copy().Width(12).Align(lipgloss.Center)
			}

			// Points
			if col == 7 {
				style = style.Copy().Width(8).Align(lipgloss.Center)
			}

			// Goals For, Against, Difference
			if col >= 8 {
				style = style.Copy().Width(5).Align(lipgloss.Center)
			}

			return style
		}).
		Headers("RANK", "TEAM", "GP", "WINS", "DRAWS", "LOSES", "RECENT FORM", "POINTS", "GF", "GA", "GD").
		Rows(records...)

	fmt.Println(t)
	return nil
}

func ChampionsFlag(records [][]string) error {
	fmt.Println("Teams currently qualified for the Champions League")

	championsRecords := records[1:5]
	testSlices := make([][]string, 0, 4)

	for position := 0; position < 4; position++ {
		soloSlice := []string{championsRecords[position][1], championsRecords[position][7]}
		testSlices = append(testSlices, soloSlice)
	}

	re := lipgloss.NewRenderer(os.Stdout)

	headerStyle := re.NewStyle().Foreground(colorSwitch).Bold(true).Align(lipgloss.Center)
	BorderStyle := lipgloss.NewStyle().Foreground(green)

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		Headers(headerStyle.String()).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style
			// Team
			if col == 0 {
				style = style.Copy().Width(20)
			}

			// Points
			if col == 1 {
				style = style.Copy().Width(8).Align(lipgloss.Center)
			}
			return style
		}).
		Headers("TEAM", "POINTS").
		Rows(testSlices...)

	fmt.Println(t)
	return nil
}
