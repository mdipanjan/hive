package components

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/session"
	"github.com/mdipanjan/hive/internal/styles"
)

const (
	detailsLabelWidth = 10
	detailsValueWidth = SessionPanelWidth - 2 - 4 - detailsLabelWidth
	maxPathLength     = detailsValueWidth - 2
)

func RenderDetails(s session.Session) string {
	title := styles.PanelTitle.Render("DETAILS") + "\n\n"
	path := TruncateMiddle(CollapsePath(s.Path), maxPathLength)

	rows := []string{
		formatRow("name", TruncateMiddle(s.Name, detailsValueWidth)),
		formatRow("tool", s.Tool),
		formatRow("path", path),
		formatStatusRow(s.Status),
		formatRow("created", shortTime(s.CreatedAt)),
		formatRow("activity", shortTime(s.LastActivity)),
	}
	return styles.Panel.Width(SessionPanelWidth).Render(title + strings.Join(rows, "\n"))
}

func formatRow(label, value string) string {
	labelStyle := styles.Label.Width(detailsLabelWidth)
	valueStyle := styles.Value.Width(detailsValueWidth).Align(lipgloss.Right)
	return labelStyle.Render(label) + valueStyle.Render(value)
}

func formatStatusRow(status session.Status) string {
	labelStyle := styles.Label.Width(detailsLabelWidth)
	glyph := lipgloss.NewStyle().Foreground(GetStatusColor(status)).Render(GetStatusGlyph(status))
	value := glyph + " " + GetStatusLabel(status)
	valueStyle := lipgloss.NewStyle().Width(detailsValueWidth).Align(lipgloss.Right)
	return labelStyle.Render("status") + valueStyle.Render(value)
}

// shortTime renders a compact relative time: now, 5m ago, 3h ago, 2d ago.
func shortTime(t time.Time) string {
	if t.IsZero() {
		return "—"
	}
	d := time.Since(t)
	switch {
	case d < time.Minute:
		return "now"
	case d < time.Hour:
		return fmt.Sprintf("%dm ago", int(d.Minutes()))
	case d < 24*time.Hour:
		return fmt.Sprintf("%dh ago", int(d.Hours()))
	default:
		return fmt.Sprintf("%dd ago", int(d.Hours())/24)
	}
}
