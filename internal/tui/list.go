package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive-v0/internal/config"
	"github.com/mdipanjan/hive-v0/internal/logger"
	"github.com/mdipanjan/hive-v0/internal/styles"
	"github.com/mdipanjan/hive-v0/internal/tmux"
)

func (m Model) updateList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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
			logger.Log.Debug("attaching to session", "name", name)
			c := tmux.AttachCmd(name)
			return m, tea.ExecProcess(c, func(err error) tea.Msg {
				return sessionAttachedMsg{err}
			})
		}

	case "n":
		m.viewMode = "new"
		m.form = newForm()
		m.form.FilePicker = newFilePicker()

	case "d":
		if len(m.sessions) > 0 {
			m.isConfirmingDelete = true
			m.deleteButton = 1
		}

	case "t":
		theme := styles.NextTheme()
		logger.Log.Debug("theme switched", "theme", theme.Name)
		config.Save(config.Config{Theme: theme.Key})

	case "?":
		m.isShowingHelp = true

	case "/":
		m.searchInput = textinput.New()
		m.searchInput.Placeholder = "Search..."
		m.searchInput.Focus()
		m.searchInput.CharLimit = 30
		m.searchInput.Width = 30
		m.searchInput.PromptStyle = lipgloss.NewStyle().Foreground(styles.ColorCyan)
		m.searchInput.TextStyle = lipgloss.NewStyle().Foreground(styles.ColorWhite)
		m.isSearching = true
		m.searchResults = m.getIndices("")
		m.searchCursor = 0
		return m, textinput.Blink
	}
	return m, nil
}
