package components

import (
	"fmt"

	"github.com/mdipanjan/hive/internal/session"
	"github.com/mdipanjan/hive/internal/styles"
)

func RenderStats(sessions []session.Session) string {
	var active, running, ready, done, failed int

	for _, s := range sessions {
		switch s.Status {
		case session.StatusActive:
			active++
		case session.StatusRunning:
			running++
		case session.StatusReady:
			ready++
		case session.StatusCompleted:
			done++
		case session.StatusFailed:
			failed++
		}
	}

	total := len(sessions)

	stats := fmt.Sprintf("%s %d  %s %d  %s %d  %s %d  %s %d  │  %d sessions",
		styles.IconActive, active,
		styles.IconRunning, running,
		styles.IconReady, ready,
		styles.IconCompleted, done,
		styles.IconFailed, failed,
		total,
	)

	return styles.Stats.Render(stats)
}
