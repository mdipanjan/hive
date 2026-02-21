package tui

import (
	"crypto/rand"
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

func generateID(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	rand.Read(b)
	for i := range b {
		b[i] = chars[b[i]%byte(len(chars))]
	}
	return string(b)
}

func cpuTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return cpuTickMsg(t)
	})
}
