package tui

import (
	"crypto/rand"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive-v0/internal/components"
	"github.com/mdipanjan/hive-v0/internal/logger"
	"github.com/mdipanjan/hive-v0/internal/session"
	"github.com/mdipanjan/hive-v0/internal/styles"
	"github.com/mdipanjan/hive-v0/internal/tmux"
)

var Tools = []string{"pi", "claude", "gemini", "opencode", "bash"}

const (
	FocusTool = iota
	FocusPath
	FocusName
	FocusButtons
)

type NewSessionForm struct {
	Tool       int
	FilePicker filepicker.Model
	Path       string
	Name       string
	Focus      int
	Button     int
}

type Model struct {
	width           int
	height          int
	sessions        []session.Session
	cursor          int
	viewMode        string
	form            NewSessionForm
	PickingPath     bool
	cpuUsageHistory []int
	err             error
}

type cpuTickMsg time.Time
type sessionAttachedMsg struct{ err error }

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
	return cpuTick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.PickingPath {
		return m.updateFilePicker(msg)
	}

	switch msg := msg.(type) {
	case sessionAttachedMsg:
		logger.Log.Debug("returned from attach", "err", msg.err)
		sessions, _ := tmux.List()
		m.sessions = sessions
		return m, tea.ClearScreen

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		if m.viewMode == "new" {
			return m.updateNewForm(msg)
		}
		return m.updateList(msg)

	case cpuTickMsg:
		cpuValue := components.GetCPUPercent()
		m.cpuUsageHistory = append(m.cpuUsageHistory, cpuValue)
		if len(m.cpuUsageHistory) > 60 {
			m.cpuUsageHistory = m.cpuUsageHistory[1:]
		}
		return m, cpuTick()
	}
	return m, nil
}

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
			name := m.sessions[m.cursor].Name
			logger.Log.Debug("deleting session", "name", name)
			tmux.Kill(name)
			m.sessions, _ = tmux.List()
			if m.cursor >= len(m.sessions) {
				m.cursor = max(0, len(m.sessions)-1)
			}
		}

	case "t":
		theme := styles.NextTheme()
		logger.Log.Debug("theme switched", "theme", theme.Name)
	}
	return m, nil
}

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
		if m.form.Focus == FocusTool {
			m.form.Tool = (m.form.Tool - 1 + len(Tools)) % len(Tools)
		} else if m.form.Focus == FocusButtons {
			m.form.Button = 0
		}

	case "right":
		if m.form.Focus == FocusTool {
			m.form.Tool = (m.form.Tool + 1) % len(Tools)
		} else if m.form.Focus == FocusButtons {
			m.form.Button = 1
		}

	case "b":
		if m.form.Focus == FocusPath {
			m.PickingPath = true
			return m, m.form.FilePicker.Init()
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
			if m.form.Focus == FocusPath {
				m.form.Path += key
			} else if m.form.Focus == FocusName {
				m.form.Name += key
			}
		}
	}

	return m, nil
}

func (m Model) updateFilePicker(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "esc" {
		m.PickingPath = false
		return m, nil
	}

	var cmd tea.Cmd
	m.form.FilePicker, cmd = m.form.FilePicker.Update(msg)

	if didSelect, path := m.form.FilePicker.DidSelectFile(msg); didSelect {
		m.form.Path = path
		m.PickingPath = false
	}

	return m, cmd
}

func newForm() NewSessionForm {
	return NewSessionForm{
		Tool:   4,
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

func (m Model) View() string {
	return RenderView(m)
}

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
