package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive-v0/internal/components"
	"github.com/mdipanjan/hive-v0/internal/styles"
)

func RenderView(m Model) string {
	// LEFT COLUMN: Logo + Honeycomb (stacked vertically)
	leftColumn := lipgloss.JoinVertical(lipgloss.Left,
		components.RenderLogo(),
		"",
		components.RenderHoneycomb(3, 3),
	)

	// RIGHT COLUMN: Sessions panel
	rightColumn := lipgloss.JoinVertical(lipgloss.Left,
		components.RenderSessions(m.sessions, m.cursor),
		components.RenderDetails(m.sessions[m.cursor]),
	)

	// COMBINE: Left + Right side by side
	mainContent := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, "  ", rightColumn)

	// Measure width for box
	contentWidth := maxLineWidth(mainContent)
	boxWidth := contentWidth + 6

	// Clamp to terminal width
	if m.width > 0 && boxWidth > m.width-2 {
		boxWidth = m.width - 2
	}

	boxStyle := styles.OuterBox.Width(boxWidth)

	// HELP BAR: Outside the main box
	help := styles.Help.Render("  n: new   enter: attach   d: delete   q: quit")

	return boxStyle.Render(mainContent) + "\n" + help
}

// maxLineWidth returns the width of the longest line in a string
func maxLineWidth(s string) int {
	maxWidth := 0
	for _, line := range strings.Split(s, "\n") {
		w := lipgloss.Width(line)
		if w > maxWidth {
			maxWidth = w
		}
	}
	return maxWidth
}
