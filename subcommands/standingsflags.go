// This file holds all the logic for creating flags relating to the standings subcommand.

package subcommands

import (
	"flag"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"os"
)

// SetupStandingsFlagSet - First creating a function to set up the Flag Set
func SetupStandingsFlagSet() (*flag.FlagSet, *bool, *bool) {

	standings := flag.NewFlagSet("standings", flag.ExitOnError)

	championsFlag := standings.Bool("champions", false, "Shows teams currently promoted to the Champions League.")
	shortChampionsFlag := standings.Bool("c", false, "Shows teams currently promoted to the Champions League.")

	standings.Usage = func() {
		fmt.Println("\nDisplays the current standings for the league.")

		fmt.Print("\n")
		fmt.Println(styleBold.Render("USAGE:"))
		fmt.Println("\t", "pl-cli", styleUnderline.Render("standings"), "[flag]")

		fmt.Print("\n")
		fmt.Println(styleBold.Render("AVAILABLE FLAGS:"))
		fmt.Println("\t", "-c, --champions", "\t", "Shows teams currently promoted to the Champions League.")
		fmt.Print("\n")
	}

	return standings, championsFlag, shortChampionsFlag
}

// ChampionsFlag Creating first flag: -c, --champions
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
