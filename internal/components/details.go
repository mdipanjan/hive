package components

import (
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/mdipanjan/hive-v0/internal/session"
	"github.com/mdipanjan/hive-v0/internal/styles"
)

const maxPathLength = 20

func RenderDetails(s session.Session) string {
	title := styles.PanelTitle.Render("DETAILS") + "\n\n"
	path := TruncateMiddle(CollapsePath(s.Path), maxPathLength)

	rows := []string{
		formatRow("Name", s.Name),
		formatRow("Tool", s.Tool),
		formatRow("Path", path),
		formatRow("Status", GetStatusText(s.Status)),
		formatRow("CreatedAt", humanize.Time(s.CreatedAt)),
		formatRow("LastActivity", humanize.Time(s.LastActivity)),
	}
	return styles.Panel.Copy().Width(PanelWidth).Render(title + strings.Join(rows, "\n"))
}

func formatRow(label, value string) string {
	return styles.Label.Render(label) + " " + styles.Value.Render(value)
}
