package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdipanjan/hive-v0/internal/logger"
	"github.com/mdipanjan/hive-v0/internal/tmux"
)

func (m Model) updateSearch(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.isSearching = false
		return m, nil

	case "enter":
		if len(m.searchResults) > 0 {
			idx := m.searchResults[m.searchCursor]
			name := m.sessions[idx].Name
			logger.Log.Debug("attaching to session from search", "name", name)
			m.isSearching = false
			c := tmux.AttachCmd(name)
			return m, tea.ExecProcess(c, func(err error) tea.Msg {
				return sessionAttachedMsg{err}
			})
		}
		return m, nil

	case "up":
		if m.searchCursor > 0 {
			m.searchCursor--
		}
		return m, nil

	case "down":
		if m.searchCursor < len(m.searchResults)-1 {
			m.searchCursor++
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.searchInput, cmd = m.searchInput.Update(msg)
	m.searchResults = m.getIndices(m.searchInput.Value())
	if m.searchCursor >= len(m.searchResults) {
		m.searchCursor = max(0, len(m.searchResults)-1)
	}
	return m, cmd
}

func (m Model) getIndices(query string) []int {
	indices := []int{}
	query = strings.ToLower(query)
	for i, session := range m.sessions {
		if strings.Contains(strings.ToLower(session.Name), query) {
			indices = append(indices, i)
		}
	}
	return indices
}
