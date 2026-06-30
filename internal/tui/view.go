package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/components"
	"github.com/mdipanjan/hive/internal/session"
	"github.com/mdipanjan/hive/internal/styles"
)

func RenderView(m Model) string {
	if m.app.PickingPath() {
		return renderFilePickerView(m)
	}
	if m.app.ShowingHelp() {
		return renderHelpView(m)
	}
	if m.app.ConfirmingDelete() {
		return renderDeleteConfirmView(m)
	}
	if m.app.CreatingSession() {
		return renderNewView(m)
	}
	if m.app.Searching() {
		return renderSearchView(m)
	}
	return renderListView(m)
}

func renderDeleteConfirmView(m Model) string {
	sessionName := ""
	attached := false
	if len(m.sessions) > 0 {
		s := m.sessions[m.cursor]
		sessionName = s.Name
		attached = s.Status == session.StatusActive
	}

	dialog := components.RenderDeleteConfirm(sessionName, attached, m.deleteButton)
	footer := components.RenderHints([]components.HelpItem{
		{Key: "←→", Desc: "select"},
		{Key: "⏎", Desc: "confirm"},
		{Key: "y", Desc: "yes"},
		{Key: "n/esc", Desc: "cancel"},
	})
	return renderChrome(m, dialog, footer)
}

func renderHelpView(m Model) string {
	popup := components.RenderHelpPopup()
	help := components.RenderHelp()

	if m.width > 0 && m.height > 0 {
		popup = lipgloss.Place(m.width, m.height-2, lipgloss.Center, lipgloss.Center, popup)
		help = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, help)
	}

	return popup + "\n" + help
}

func renderFilePickerView(m Model) string {
	picker := m.form.FilePicker.view()
	footer := components.RenderHints([]components.HelpItem{
		{Key: "↑↓", Desc: "navigate"},
		{Key: "→", Desc: "open"},
		{Key: "⏎", Desc: "select"},
		{Key: "esc", Desc: "cancel"},
	})
	return renderChrome(m, picker, footer)
}

func renderSearchView(m Model) string {
	title := "SEARCH"
	action := "attach"
	if m.mode == ModeSwitch {
		title = "SWITCH SESSION"
		action = "switch"
	}

	popup := components.RenderSearchPopupTitled(title, m.searchInput.View(), m.searchInput.Value(), m.sessions, m.searchResults, m.searchCursor)
	footer := components.RenderHints([]components.HelpItem{
		{Key: "↑↓", Desc: "select"},
		{Key: "⏎", Desc: action},
		{Key: "esc", Desc: "close"},
	})
	return renderChrome(m, popup, footer)
}
func renderListView(m Model) string {
	cursor := m.cursor
	if cursor >= len(m.sessions) {
		cursor = max(0, len(m.sessions)-1)
	}

	subtitle := styles.Dim.Render("session manager · v" + components.Version)
	leftColumn := lipgloss.JoinVertical(lipgloss.Left,
		components.RenderLogo(),
		"",
		subtitle,
		"",
		"",
		components.RenderStats(m.sessions),
	)

	var rightColumn string
	if len(m.sessions) > 0 {
		rightColumn = lipgloss.JoinVertical(lipgloss.Left,
			components.RenderSessions(m.sessions, cursor),
			"",
			components.RenderDetails(m.sessions[cursor]),
		)
	} else {
		rightColumn = components.RenderSessions(m.sessions, cursor)
	}

	mainContent := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, "    ", rightColumn)

	activity := components.RenderActivity(m.cpuUsageHistory, m.memUsageHistory)
	footer := components.RenderHelp()

	fullContent := lipgloss.JoinVertical(lipgloss.Left,
		components.RenderTitleBar(),
		"",
		mainContent,
		"",
		activity,
		"",
		footer,
	)

	boxWidth := maxLineWidth(fullContent) + 6
	if m.width > 0 && boxWidth > m.width-2 {
		boxWidth = m.width - 2
	}

	box := styles.OuterBox.Width(boxWidth).Render(fullContent)
	if m.width > 0 && m.height > 0 {
		box = lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
	}

	return box
}

func renderNewView(m Model) string {
	formData := components.FormData{
		Tool:   m.form.Tool,
		Path:   m.form.Path,
		Name:   m.form.Name,
		Focus:  m.form.Focus,
		Button: m.form.Button,
	}

	dialog := components.RenderNewDialog(Tools, formData)
	footer := components.RenderHints([]components.HelpItem{
		{Key: "tab", Desc: "next"},
		{Key: "←→", Desc: "select"},
		{Key: "⏎", Desc: "confirm"},
		{Key: "esc", Desc: "cancel"},
	})
	return renderChrome(m, dialog, footer)
}

// renderChrome wraps a centered body in the shared window frame: faux title bar
// at the top, the body vertically centered, and a footer hint bar at the bottom
// (DESIGN.md §4). Used by modal screens so they match the dashboard chrome.
func renderChrome(m Model, body, footer string) string {
	if m.width <= 0 || m.height <= 0 {
		return body + "\n\n" + footer
	}

	boxW := m.width - 2
	boxH := m.height - 2
	innerW := boxW - 2 - 4 // border + padding(1,2)
	innerH := boxH - 2 - 2
	bodyH := innerH - 2 // title line + footer line
	if bodyH < 1 {
		bodyH = 1
	}

	title := components.RenderTitleBar()
	placed := lipgloss.Place(innerW, bodyH, lipgloss.Center, lipgloss.Center, body)
	foot := lipgloss.PlaceHorizontal(innerW, lipgloss.Center, footer)
	content := lipgloss.JoinVertical(lipgloss.Left, title, placed, foot)

	box := styles.OuterBox.Width(boxW).Height(boxH).Render(content)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}

func maxLineWidth(s string) int {
	maxWidth := 0
	for _, line := range strings.Split(s, "\n") {
		if w := lipgloss.Width(line); w > maxWidth {
			maxWidth = w
		}
	}
	return maxWidth
}
