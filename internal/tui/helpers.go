package tui

import (
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/styles"
)

func newSearchInput() textinput.Model {
	in := textinput.New()
	in.Placeholder = "Search..."
	in.Focus()
	in.CharLimit = 30
	in.Width = 30
	in.PromptStyle = lipgloss.NewStyle().Foreground(styles.ColorCyan)
	in.TextStyle = lipgloss.NewStyle().Foreground(styles.ColorWhite)
	return in
}

func getDefaultPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return "~"
	}
	return dir
}

func appendCapped(history []int, value, capacity int) []int {
	history = append(history, value)
	if len(history) > capacity {
		history = history[len(history)-capacity:]
	}
	return history
}

func cpuTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return cpuTickMsg(t)
	})
}
