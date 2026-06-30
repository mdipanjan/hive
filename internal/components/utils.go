package components

import (
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/session"
	"github.com/mdipanjan/hive/internal/styles"
)

const PanelWidth = 35

func TruncateMiddle(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}

	ellipsis := "..."
	remaining := maxLen - len(ellipsis)
	startLen := (remaining * 2) / 3
	endLen := remaining - startLen

	return s[:startLen] + ellipsis + s[len(s)-endLen:]
}

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~") {
		if home, err := os.UserHomeDir(); err == nil {
			path = strings.Replace(path, "~", home, 1)
		}
	}
	return path
}

func CollapsePath(path string) string {
	if home, err := os.UserHomeDir(); err == nil {
		path = strings.Replace(path, home, "~", 1)
	}
	return path
}

func GetStatusIcon(status session.Status) string {
	switch status {
	case session.StatusActive:
		return styles.IconActive
	case session.StatusRunning:
		return styles.IconRunning
	case session.StatusReady:
		return styles.IconReady
	case session.StatusCompleted:
		return styles.IconCompleted
	case session.StatusFailed:
		return styles.IconFailed
	default:
		return styles.IconIdle
	}
}

func GetStatusText(status session.Status) string {
	return GetStatusIcon(status) + " " + GetStatusLabel(status)
}

// GetStatusGlyph returns the raw (uncolored) status glyph so callers can apply
// their own foreground/background (e.g. on a selected row).
func GetStatusGlyph(status session.Status) string {
	switch status {
	case session.StatusActive:
		return "●"
	case session.StatusRunning:
		return "■"
	case session.StatusCompleted:
		return "✓"
	case session.StatusFailed:
		return "✗"
	default:
		return "◌"
	}
}

// GetStatusLabel returns the design status vocabulary (DESIGN.md §2.3).
func GetStatusLabel(status session.Status) string {
	switch status {
	case session.StatusActive:
		return "attached"
	case session.StatusRunning:
		return "detached"
	case session.StatusReady:
		return "idle"
	case session.StatusCompleted:
		return "done"
	case session.StatusFailed:
		return "dead"
	default:
		return "idle"
	}
}

// GetStatusColor returns the role color paired with a status glyph.
func GetStatusColor(status session.Status) lipgloss.Color {
	switch status {
	case session.StatusActive, session.StatusCompleted:
		return styles.ColorGreen
	case session.StatusRunning:
		return styles.ColorYellow
	case session.StatusFailed:
		return styles.ColorRed
	default:
		return styles.ColorCyan
	}
}
