package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/session"
	"github.com/mdipanjan/hive/internal/styles"
)

const statsWidth = 40

// RenderStats draws the labeled status legend, a hairline divider, and the
// total session count (DESIGN.md §4.1).
func RenderStats(sessions []session.Session) string {
	var attached, detached, idle, done, dead int
	for _, s := range sessions {
		switch s.Status {
		case session.StatusActive:
			attached++
		case session.StatusRunning:
			detached++
		case session.StatusReady:
			idle++
		case session.StatusCompleted:
			done++
		case session.StatusFailed:
			dead++
		default:
			idle++
		}
	}

	item := func(glyph string, color lipgloss.Color, count int, label string) string {
		g := lipgloss.NewStyle().Foreground(color).Render(glyph)
		c := lipgloss.NewStyle().Foreground(styles.ColorWhite).Bold(true).Render(fmt.Sprintf("%d", count))
		l := styles.Dim.Render(label)
		return g + " " + c + " " + l
	}

	line1 := item("●", styles.ColorGreen, attached, "attached") + "   " +
		item("■", styles.ColorYellow, detached, "detached")
	line2 := item("◌", styles.ColorCyan, idle, "idle") + "   " +
		item("✓", styles.ColorGreen, done, "done") + "   " +
		item("✗", styles.ColorRed, dead, "dead")

	divider := styles.Dim.Render(strings.Repeat("─", statsWidth))
	total := styles.Dim.Render(fmt.Sprintf("%d sessions", len(sessions)))

	return lipgloss.JoinVertical(lipgloss.Left,
		line1,
		line2,
		"",
		divider,
		"",
		total,
	)
}
