package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/session"
	"github.com/mdipanjan/hive/internal/styles"
)

const (
	SessionPanelWidth  = 46
	maxVisibleSessions = 6
)

// sessionRowWidth is the usable width inside the panel, minus the 1-cell
// left accent-bar gutter.
func sessionRowWidth() int {
	return SessionPanelWidth - 2 /*border*/ - 4 /*padding*/ - 1 /*bar*/
}

func RenderSessions(sessions []session.Session, cursor int) string {
	title := styles.PanelTitle.Render("SESSIONS") + "\n\n"

	if len(sessions) == 0 {
		empty := styles.Dim.Render("No sessions yet — press n to create one")
		return styles.Panel.Width(SessionPanelWidth).Render(title + empty)
	}

	start, end := calculateWindow(len(sessions), cursor, maxVisibleSessions)

	var rows []string
	if start > 0 {
		rows = append(rows, styles.Dim.Render(fmt.Sprintf("  ▲ %d more", start)))
	}
	for index := start; index < end; index++ {
		rows = append(rows, renderSessionRow(sessions[index], index == cursor))
	}
	if end < len(sessions) {
		rows = append(rows, styles.Dim.Render(fmt.Sprintf("  ▼ %d more", len(sessions)-end)))
	}

	return styles.Panel.Width(SessionPanelWidth).Render(title + strings.Join(rows, "\n"))
}

func renderSessionRow(s session.Session, selected bool) string {
	width := sessionRowWidth()
	statusColor := GetStatusColor(s.Status)
	glyph := GetStatusIcon(s.Status) // already colored
	label := GetStatusLabel(s.Status)

	// name budget: width - glyph(1) - space(1) - gap(1) - label
	nameBudget := width - 3 - lipgloss.Width(label)
	if nameBudget < 4 {
		nameBudget = 4
	}
	name := TruncateMiddle(s.Name, nameBudget)

	pad := width - 2 - lipgloss.Width(name) - lipgloss.Width(label)
	if pad < 1 {
		pad = 1
	}

	if selected {
		bg := styles.ColorDim
		bar := lipgloss.NewStyle().Foreground(statusColor).Render("▎")
		g := lipgloss.NewStyle().Foreground(statusColor).Background(bg).Render(GetStatusGlyph(s.Status))
		n := lipgloss.NewStyle().Foreground(styles.ColorWhite).Background(bg).Bold(true).Render(" " + name)
		fill := lipgloss.NewStyle().Background(bg).Render(strings.Repeat(" ", pad))
		lab := lipgloss.NewStyle().Foreground(styles.ColorWhite).Background(bg).Render(label)
		return bar + g + n + fill + lab
	}

	n := styles.Normal.Render(" " + name)
	fill := strings.Repeat(" ", pad)
	lab := styles.Dim.Render(label)
	return " " + glyph + n + fill + lab
}

func calculateWindow(total, cursor, maxVisible int) (int, int) {
	if total <= maxVisible {
		return 0, total
	}

	half := maxVisible / 2
	start := cursor - half
	end := cursor + half + (maxVisible % 2)

	if start < 0 {
		start = 0
		end = maxVisible
	}
	if end > total {
		end = total
		start = total - maxVisible
	}

	return start, end
}
