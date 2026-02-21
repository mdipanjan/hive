package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive-v0/internal/components"
	"github.com/mdipanjan/hive-v0/internal/styles"
)

func RenderView(m Model) string {
	if m.PickingPath {
		return renderFilePickerView(m)
	}
	if m.viewMode == "new" {
		return renderNewView(m)
	}
	return renderListView(m)
}

func renderFilePickerView(m Model) string {
	title := styles.PanelTitle.Render("SELECT DIRECTORY") + "\n\n"
	picker := m.form.FilePicker.View()

	panelStyle := styles.Panel.Copy().Width(50).Padding(1, 2)
	panel := panelStyle.Render(title + picker)

	if m.width > 0 && m.height > 0 {
		panel = lipgloss.Place(m.width, m.height-2, lipgloss.Center, lipgloss.Center, panel)
	}

	help := styles.Help.Render("  ↑↓: navigate   enter: select/open   esc: cancel")
	return panel + "\n" + help
}

func renderListView(m Model) string {
	leftColumn := lipgloss.JoinVertical(lipgloss.Left,
		components.RenderLogo(),
		"",
		components.RenderHoneycomb(3, 3),
		"",
		components.RenderStats(m.sessions),
	)

	var rightColumn string
	if len(m.sessions) > 0 {
		rightColumn = lipgloss.JoinVertical(lipgloss.Left,
			components.RenderSessions(m.sessions, m.cursor),
			components.RenderDetails(m.sessions[m.cursor]),
		)
	} else {
		rightColumn = components.RenderSessions(m.sessions, m.cursor)
	}

	mainContent := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, "  ", rightColumn)

	contentWidth := maxLineWidth(mainContent)
	boxWidth := contentWidth + 6

	if m.width > 0 && boxWidth > m.width-2 {
		boxWidth = m.width - 2
	}

	activity := components.RenderActivity(contentWidth, m.cpuUsageHistory)
	fullContent := lipgloss.JoinVertical(lipgloss.Right, mainContent, "", activity)

	boxStyle := styles.OuterBox.Width(boxWidth)
	help := components.RenderHelp()

	return boxStyle.Render(fullContent) + "\n" + help
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

	if m.width > 0 && m.height > 0 {
		dialog = lipgloss.Place(m.width, m.height-2, lipgloss.Center, lipgloss.Center, dialog)
	}

	help := styles.Help.Render("  tab: next   ←→: select   enter: confirm   esc: cancel")
	return dialog + "\n" + help
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
