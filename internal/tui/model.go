package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdipanjan/hive/internal/components"
	"github.com/mdipanjan/hive/internal/lifecycle"
	"github.com/mdipanjan/hive/internal/logger"
	"github.com/mdipanjan/hive/internal/session"
)

type Model struct {
	width           int
	height          int
	sessions        []session.Session
	cursor          int
	app             AppState
	form            NewSessionForm
	searchInput     textinput.Model
	searchResults   []int
	searchCursor    int
	deleteButton    int
	cpuUsageHistory []int
	err             error
}

type NewSessionForm struct {
	Tool       int
	FilePicker filepicker.Model
	Path       string
	Name       string
	Focus      int
	Button     int
}

type cpuTickMsg time.Time
type sessionAttachedMsg struct{ err error }

func New() Model {
	logger.Log.Debug("initializing model")
	sessions, err := lifecycle.New().List()
	if err != nil {
		logger.Log.Error("failed to list sessions", "err", err)
	}
	logger.Log.Debug("found sessions", "count", len(sessions))

	return Model{
		sessions: sessions,
		cursor:   0,
		app:      NewAppState(),
		err:      err,
	}
}

func (m Model) Init() tea.Cmd {
	return cpuTick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.app.PickingPath() {
		return m.updateFilePicker(msg)
	}

	if m.app.Searching() {
		if _, ok := msg.(tea.KeyMsg); !ok {
			var cmd tea.Cmd
			m.searchInput, cmd = m.searchInput.Update(msg)
			return m, cmd
		}
	}

	switch msg := msg.(type) {
	case sessionAttachedMsg:
		logger.Log.Debug("returned from attach", "err", msg.err)
		sessions, _ := lifecycle.New().List()
		m.sessions = sessions
		return m, tea.ClearScreen

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		if m.app.Searching() {
			return m.updateSearch(msg)
		}
		if m.app.ShowingHelp() {
			if msg.String() == "?" || msg.String() == "esc" {
				m.app.CloseOverlay()
			}
			return m, nil
		}
		if m.app.ConfirmingDelete() {
			return m.updateDeleteConfirm(msg)
		}
		if m.app.CreatingSession() {
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

func (m Model) View() string {
	return RenderView(m)
}
