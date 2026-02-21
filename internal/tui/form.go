package tui

import (
	"os"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/components"
	"github.com/mdipanjan/hive/internal/logger"
	"github.com/mdipanjan/hive/internal/styles"
	"github.com/mdipanjan/hive/internal/tmux"
)

var Tools = []string{"pi", "claude", "bash"}

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
		m.viewMode = "list"
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
				m.viewMode = "list"
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
					m.isPickingPath = true
					return m, m.form.FilePicker.Init()
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
		Tool:   2,
		Path:   getDefaultPath(),
		Name:   "",
		Focus:  FocusTool,
		Button: 0,
	}
}

func newFilePicker() filepicker.Model {
	fp := filepicker.New()
	fp.DirAllowed = true
	fp.FileAllowed = false
	fp.ShowPermissions = false
	fp.ShowSize = false
	fp.ShowHidden = false
	fp.Height = 15
	fp.CurrentDirectory, _ = os.UserHomeDir()

	fp.Styles.Cursor = lipgloss.NewStyle().Foreground(styles.ColorCyan)
	fp.Styles.Directory = lipgloss.NewStyle().Foreground(styles.ColorCyan)
	fp.Styles.File = lipgloss.NewStyle().Foreground(styles.ColorWhite)
	fp.Styles.Selected = lipgloss.NewStyle().Foreground(styles.ColorGreen).Bold(true)

	return fp
}

func (m Model) createSession() (tea.Model, tea.Cmd) {
	tool := Tools[m.form.Tool]
	path := components.ExpandPath(m.form.Path)
	name := m.form.Name

	if name == "" {
		name = "hive-" + generateID(6)
	}

	logger.Log.Debug("creating session", "name", name, "tool", tool, "path", path)
	tmux.Create(name, tool, path)

	m.sessions, _ = tmux.List()
	m.viewMode = "list"
	return m, nil
}
