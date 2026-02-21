package components

import (
	"os"
	"strings"

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
	case session.StatusRunning:
		return styles.IconRunning
	case session.StatusWaiting:
		return styles.IconWaiting
	default:
		return styles.IconIdle
	}
}

func GetStatusText(status session.Status) string {
	switch status {
	case session.StatusRunning:
		return styles.IconRunning + " running"
	case session.StatusWaiting:
		return styles.IconWaiting + " waiting"
	default:
		return styles.IconIdle + " idle"
	}
}
