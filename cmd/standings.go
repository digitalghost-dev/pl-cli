package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"log"
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

func DisplayStandings(url string) error {

	// Perform HTTP GET request to fetch the CSV file
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to fetch CSV file: %v", err)
	}
	defer response.Body.Close()

	// Parse the CSV file
	reader := csv.NewReader(response.Body)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to parse CSV: %v", err)
	}

	records = records[1:]

	re := lipgloss.NewRenderer(os.Stdout)

	headerStyle := re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
	BorderStyle := lipgloss.NewStyle().Foreground(purple)
	positionColors := map[string]lipgloss.Style{
		"championsLeague": lipgloss.NewStyle().Foreground(lipgloss.Color(green)),
		"europaLeague":    lipgloss.NewStyle().Foreground(lipgloss.Color(blue)),
		"relegation":      lipgloss.NewStyle().Foreground(lipgloss.Color(red)),
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

			// Wins
			if col == 3 {
				style = style.Copy().Width(6).Align(lipgloss.Center)
			}

			// Draws
			if col == 4 {
				style = style.Copy().Width(7).Align(lipgloss.Center)
			}

			// Loses
			if col == 5 {
				style = style.Copy().Width(7).Align(lipgloss.Center)
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
