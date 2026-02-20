package components

import (
	"strings"

	"github.com/mdipanjan/hive-v0/internal/styles"
)

// RenderHoneycomb returns a decorative honeycomb pattern
// rows: number of rows, cols: base number of hexagons per row
func RenderHoneycomb(rows, cols int) string {
	var lines []string

	for i := 0; i < rows; i++ {
		var line string

		if i%2 == 0 {
			// Odd rows (0, 2, 4...) - more offset, n hexagons
			line = "  "
			for j := 0; j < cols; j++ {
				line += "⬡ "
			}
		} else {
			// Even rows (1, 3, 5...) - less offset, n+1 hexagons
			line = " "
			for j := 0; j < cols+1; j++ {
				line += "⬡ "
			}
		}

		lines = append(lines, strings.TrimRight(line, " "))
	}

	return styles.Logo.Render(strings.Join(lines, "\n"))
}
