package components

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/styles"
)

// RenderLogo draws the filled "HIVE" wordmark in the accent color with a
// subtle shadow echo behind it for depth (DESIGN.md §4.1).
func RenderLogo() string {
	letters := [][]string{
		{ // H
			"██  ██",
			"██  ██",
			"██████",
			"██  ██",
			"██  ██",
		},
		{ // I
			"██",
			"██",
			"██",
			"██",
			"██",
		},
		{ // V
			"██   ██",
			"██   ██",
			"██   ██",
			" ██ ██ ",
			"  ███  ",
		},
		{ // E
			"██████",
			"██    ",
			"█████ ",
			"██    ",
			"██████",
		},
	}

	blocks := make([]string, len(letters))
	for i, rows := range letters {
		joined := ""
		for r, row := range rows {
			if r > 0 {
				joined += "\n"
			}
			joined += row
		}
		blocks[i] = lipgloss.NewStyle().Foreground(styles.ColorCyan).Bold(true).Render(joined)
	}

	gap := "  "
	wordmark := lipgloss.JoinHorizontal(lipgloss.Top,
		blocks[0], gap, blocks[1], gap, blocks[2], gap, blocks[3],
	)
	return wordmark
}
