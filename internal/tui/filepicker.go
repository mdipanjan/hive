package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) updateFilePicker(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch keyMsg.String() {
	case "esc":
		m.app.CloseOverlay()
	case "up", "k":
		m.form.FilePicker.moveUp()
	case "down", "j":
		m.form.FilePicker.moveDown()
	case "right", "l":
		m.form.FilePicker.descend()
	case "left", "h", "backspace":
		m.form.FilePicker.ascend()
	case "enter":
		if path, ok := m.form.FilePicker.selected(); ok {
			m.form.Path = path
			m.app.CloseOverlay()
		}
	}

	return m, nil
}
