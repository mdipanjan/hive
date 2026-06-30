package tui

import (
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func getDefaultPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return "~"
	}
	return dir
}

func cpuTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return cpuTickMsg(t)
	})
}
