package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdipanjan/hive-v0/internal/logger"
	"github.com/mdipanjan/hive-v0/internal/session"
	"github.com/mdipanjan/hive-v0/internal/tmux"
)

type Model struct {
	width    int
	height   int
	sessions []session.Session
	cursor   int
	viewMode string
	err      error
}

type sessionAttachedMsg struct {
	err error
}

func New() Model {
	logger.Log.Debug("initializing model")
	sessions, err := tmux.List()
	if err != nil {
		logger.Log.Error("failed to list sessions", "err", err)
	}
	logger.Log.Debug("found sessions", "count", len(sessions))

	return Model{
		sessions: sessions,
		cursor:   0,
		viewMode: "list",
		err:      err,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case sessionAttachedMsg:
		logger.Log.Debug("returned from attach", "err", msg.err)
		sessions, err := tmux.List()
		if err != nil {
			logger.Log.Error("failed to list sessions after attach", "err", err)
		}
		m.sessions = sessions
		return m, tea.ClearScreen
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		logger.Log.Debug("key pressed", "key", msg.String())
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			m.cursor = max(0, m.cursor-1)

		case "down", "j":
			m.cursor = min(len(m.sessions)-1, m.cursor+1)

		case "enter":
			if len(m.sessions) > 0 {
				name := m.sessions[m.cursor].Name
				logger.Log.Debug("attaching to session", "name", name, "insideTmux", tmux.IsInsideTmux())
				c := tmux.AttachCmd(name)
				return m, tea.ExecProcess(c, func(err error) tea.Msg {
					return sessionAttachedMsg{err}
				})
			}

		case "n":
			name := fmt.Sprintf("hive-%d", time.Now().Unix())
			logger.Log.Debug("creating session", "name", name)
			if err := tmux.Create(name, "bash", "~"); err != nil {
				logger.Log.Error("failed to create session", "err", err)
			}
			m.sessions, _ = tmux.List()

		case "d":
			if len(m.sessions) > 0 {
				name := m.sessions[m.cursor].Name
				logger.Log.Debug("deleting session", "name", name)
				if err := tmux.Kill(name); err != nil {
					logger.Log.Error("failed to kill session", "err", err)
				}
				m.sessions, _ = tmux.List()
				if m.cursor >= len(m.sessions) {
					m.cursor = max(0, len(m.sessions)-1)
				}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	return RenderView(m)
}
