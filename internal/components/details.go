package components

import (
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/mdipanjan/hive-v0/internal/session"
	"github.com/mdipanjan/hive-v0/internal/styles"
)

// RenderDetails returns the selected session details panel
func RenderDetails(s session.Session) string {
	rows := []string{
		formatRow("Name", s.Name),
		formatRow("Tool", s.Tool),
		formatRow("Path", s.Path),
		formatRow("Status", formatStatus(s.Status)),
		formatRow("CreatedAt", humanize.Time(s.CreatedAt)),
		formatRow("LastActivity", humanize.Time(s.LastActivity)),
	}
	return styles.Panel.Render(strings.Join(rows, "\n"))
}

func formatRow(label, value string) string {
	formattedLabel := styles.Label.Render(label)
	formattedValue := styles.Value.Render(value)
	return formattedLabel + " " + formattedValue
}

func formatStatus(status session.Status) string {
	switch status {
	case session.StatusRunning:
		return styles.IconRunning + " running"
	case session.StatusWaiting:
		return styles.IconWaiting + " waiting"
	case session.StatusIdle:
		return styles.IconIdle + " idle"
	}
	return ""
}
