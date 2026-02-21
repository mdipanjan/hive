package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) updateFilePicker(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "esc" {
		m.isPickingPath = false
		return m, nil
	}

	var cmd tea.Cmd
	m.form.FilePicker, cmd = m.form.FilePicker.Update(msg)

	if didSelect, path := m.form.FilePicker.DidSelectFile(msg); didSelect {
		m.form.Path = path
		m.isPickingPath = false
	}

	return m, cmd
}
