package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdipanjan/hive-v0/internal/logger"
	"github.com/mdipanjan/hive-v0/internal/tmux"
)

func (m Model) updateDeleteConfirm(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "n":
		m.isConfirmingDelete = false

	case "left", "right", "h", "l":
		m.deleteButton = 1 - m.deleteButton

	case "y":
		m.deleteButton = 0
		return m.executeDelete()

	case "enter":
		if m.deleteButton == 0 {
			return m.executeDelete()
		}
		m.isConfirmingDelete = false
	}
	return m, nil
}

func (m Model) executeDelete() (tea.Model, tea.Cmd) {
	if len(m.sessions) > 0 {
		name := m.sessions[m.cursor].Name
		logger.Log.Debug("deleting session", "name", name)
		tmux.Kill(name)
		m.sessions, _ = tmux.List()
		if m.cursor >= len(m.sessions) {
			m.cursor = max(0, len(m.sessions)-1)
		}
	}
	m.isConfirmingDelete = false
	return m, nil
}
