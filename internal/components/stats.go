package components

import (
	"fmt"

	"github.com/mdipanjan/hive/internal/session"
	"github.com/mdipanjan/hive/internal/styles"
)

func RenderStats(sessions []session.Session) string {
	var running, waiting, idle int

	for _, s := range sessions {
		switch s.Status {
		case session.StatusRunning:
			running++
		case session.StatusWaiting:
			waiting++
		case session.StatusIdle:
			idle++
		}
	}

	total := len(sessions)

	stats := fmt.Sprintf("%s %d  %s %d  %s %d  │  %d sessions",
		styles.IconRunning, running,
		styles.IconWaiting, waiting,
		styles.IconIdle, idle,
		total,
	)

	return styles.Stats.Render(stats)
}
