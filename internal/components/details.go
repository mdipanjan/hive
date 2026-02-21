package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"github.com/mdipanjan/hive-v0/internal/session"
	"github.com/mdipanjan/hive-v0/internal/styles"
)

const (
	maxPathLength = 20
	labelWidth    = 10
	valueWidth    = 18
)

func RenderDetails(s session.Session) string {
	title := styles.PanelTitle.Render("DETAILS") + "\n\n"
	path := TruncateMiddle(CollapsePath(s.Path), maxPathLength)

	rows := []string{
		formatRow("Name", TruncateMiddle(s.Name, valueWidth)),
		formatRow("Tool", s.Tool),
		formatRow("Path", path),
		formatRow("Status", GetStatusText(s.Status)),
		formatRow("Created", humanize.Time(s.CreatedAt)),
		formatRow("Activity", humanize.Time(s.LastActivity)),
	}
	return styles.Panel.Width(PanelWidth).Render(title + strings.Join(rows, "\n"))
}

func formatRow(label, value string) string {
	labelStyle := styles.Label.Width(labelWidth)
	valueStyle := styles.Value.Width(valueWidth).Align(lipgloss.Right)
	return labelStyle.Render(label) + valueStyle.Render(value)
}
