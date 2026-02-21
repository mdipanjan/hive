package components

import (
	"fmt"

	"github.com/mdipanjan/hive-v0/internal/session"
	"github.com/mdipanjan/hive-v0/internal/styles"
)

const (
	maxNameLength      = 20
	maxVisibleSessions = 6
)

func RenderSessions(sessions []session.Session, cursor int) string {
	title := styles.PanelTitle.Render("SESSIONS") + "\n\n"
	var s string

	if len(sessions) == 0 {
		s = "  No sessions. Press 'n' to create one.\n"
		return styles.Panel.Width(PanelWidth).Render(title + s)
	}

	start, end := calculateWindow(len(sessions), cursor, maxVisibleSessions)

	if start > 0 {
		s += styles.Dim.Render("  ▲ "+fmt.Sprintf("%d more", start)) + "\n"
	}

	for index := start; index < end; index++ {
		sessionItem := sessions[index]
		name := TruncateMiddle(sessionItem.Name, maxNameLength)

		if cursor == index {
			line := "● " + name
			s += styles.Logo.Render(line) + "\n"
		} else {
			icon := GetStatusIcon(sessionItem.Status)
			line := icon + " " + name
			s += styles.Normal.Render(line) + "\n"
		}
	}

	if end < len(sessions) {
		s += styles.Dim.Render("  ▼ "+fmt.Sprintf("%d more", len(sessions)-end)) + "\n"
	}

	return styles.Panel.Width(PanelWidth).Render(title + s)
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
