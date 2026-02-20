package components

import (
	"github.com/mdipanjan/hive-v0/internal/session"
	"github.com/mdipanjan/hive-v0/internal/styles"
)

// RenderSessions returns the sessions list panel
func RenderSessions(sessions []session.Session, cursor int) string {
	var s string

	// Sessions list
	for index, session := range sessions {
		icon := styles.IconIdle
		line := icon + " " + session.Name

		if cursor == index {
			line = styles.Selected.Render(line + " ←")
		} else {
			line = styles.Normal.Render(line)
		}
		s += line + "\n"
	}

	// Empty state
	if len(sessions) == 0 {
		s = "  No sessions. Press 'n' to create one.\n"
	}

	// Wrap in bordered panel
	return styles.Panel.Render(s)
}
