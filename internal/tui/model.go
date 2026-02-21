package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdipanjan/hive-v0/internal/components"
	"github.com/mdipanjan/hive-v0/internal/logger"
	"github.com/mdipanjan/hive-v0/internal/session"
	"github.com/mdipanjan/hive-v0/internal/tmux"
)

type Model struct {
	width              int
	height             int
	sessions           []session.Session
	cursor             int
	viewMode           string
	form               NewSessionForm
	isPickingPath      bool
	isShowingHelp      bool
	isSearching        bool
	searchInput        textinput.Model
	searchResults      []int
	searchCursor       int
	isConfirmingDelete bool
	deleteButton       int
	cpuUsageHistory    []int
	err                error
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
	if m.isPickingPath {
		return m.updateFilePicker(msg)
	}

	if m.isSearching {
		if _, ok := msg.(tea.KeyMsg); !ok {
			var cmd tea.Cmd
			m.searchInput, cmd = m.searchInput.Update(msg)
			return m, cmd
		}
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
		if m.isSearching {
			return m.updateSearch(msg)
		}
		if m.isShowingHelp {
			if msg.String() == "?" || msg.String() == "esc" {
				m.isShowingHelp = false
			}
			return m, nil
		}
		if m.isConfirmingDelete {
			return m.updateDeleteConfirm(msg)
		}
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

func (m Model) View() string {
	return RenderView(m)
}
