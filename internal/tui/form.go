package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdipanjan/hive/internal/lifecycle"
	"github.com/mdipanjan/hive/internal/logger"
)

var Tools = lifecycle.BuiltinTools

const (
	FocusTool = iota
	FocusPath
	FocusName
	FocusButtons
)

func (m Model) updateNewForm(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.String()

	switch key {
	case "esc":
		m.app.ReturnToSessionList()
		return m, nil

	case "tab", "down":
		if m.form.Focus < FocusButtons {
			m.form.Focus++
		}

	case "shift+tab", "up":
		if m.form.Focus > FocusTool {
			m.form.Focus--
		}

	case "left":
		switch m.form.Focus {
		case FocusTool:
			m.form.Tool = (m.form.Tool - 1 + len(Tools)) % len(Tools)
		case FocusButtons:
			m.form.Button = 0
		}

	case "right":
		switch m.form.Focus {
		case FocusTool:
			m.form.Tool = (m.form.Tool + 1) % len(Tools)
		case FocusButtons:
			m.form.Button = 1
		}

	case "enter":
		if m.form.Focus == FocusButtons {
			if m.form.Button == 1 {
				m.app.ReturnToSessionList()
				return m, nil
			}
			return m.createSession()
		}
		if m.form.Focus < FocusButtons {
			m.form.Focus++
		}

	case "backspace":
		if m.form.Focus == FocusPath && len(m.form.Path) > 0 {
			m.form.Path = m.form.Path[:len(m.form.Path)-1]
		} else if m.form.Focus == FocusName && len(m.form.Name) > 0 {
			m.form.Name = m.form.Name[:len(m.form.Name)-1]
		}

	default:
		if len(key) == 1 {
			switch m.form.Focus {
			case FocusPath:
				if key == "b" {
					m.app.PickPath()
					m.form.FilePicker = newFilePicker()
					return m, nil
				}
				m.form.Path += key
			case FocusName:
				m.form.Name += key
			}
		}
	}

	return m, nil
}

func newForm() NewSessionForm {
	return NewSessionForm{
		Tool:   3,
		Path:   getDefaultPath(),
		Name:   "",
		Focus:  FocusTool,
		Button: 0,
	}
}

func newFilePicker() dirPicker {
	return newDirPicker(getDefaultPath())
}

func (m Model) createSession() (tea.Model, tea.Cmd) {
	tool := Tools[m.form.Tool]
	path := m.form.Path
	name := m.form.Name

	logger.Log.Debug("creating session", "name", name, "tool", tool, "path", path)
	createdName, err := lifecycle.New().Create(lifecycle.CreateRequest{Name: name, Tool: tool, Path: path})
	if err != nil {
		m.err = err
		return m, nil
	}
	logger.Log.Debug("created session", "name", createdName)

	m.sessions, _ = lifecycle.New().List()
	m.app.ReturnToSessionList()
	return m, nil
}
