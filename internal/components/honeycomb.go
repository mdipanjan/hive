package components

import (
	"strings"

	"github.com/mdipanjan/hive/internal/styles"
)

func RenderHoneycomb(rows, cols int) string {
	var lines []string

	for i := 0; i < rows; i++ {
		var line string

		if i%2 == 0 {
			line = "  "
			for j := 0; j < cols; j++ {
				line += "⬡ "
			}
		} else {
			line = " "
			for j := 0; j < cols+1; j++ {
				line += "⬡ "
			}
		}

		lines = append(lines, strings.TrimRight(line, " "))
	}

	return styles.Logo.Render(strings.Join(lines, "\n"))
}
