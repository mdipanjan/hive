package components

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/styles"
)

// RenderDeleteConfirm draws the destructive confirmation dialog. "No" is the
// safe default focus; when "Yes" is selected it takes a red fill, and an
// attached session adds an explicit warning (DESIGN.md §4.6).
func RenderDeleteConfirm(sessionName string, attached bool, selectedButton int) string {
	title := styles.PanelTitle.Render("DELETE SESSION")
	message := styles.Normal.Render(`Delete "` + sessionName + `"?`)

	body := title + "\n\n" + message
	if attached {
		warn := styles.Dim.Render("It's currently attached and will be killed.")
		body += "\n" + warn
	}

	var opts []ButtonOption
	if selectedButton == 0 {
		opts = append(opts, Destructive())
	}
	buttons := RenderButtons([]string{"Yes", "No"}, selectedButton, true, opts...)

	content := body + "\n\n" + buttons
	dialog := styles.Panel.Width(48).Padding(1, 2).Render(content)
	return lipgloss.PlaceHorizontal(lipgloss.Width(dialog), lipgloss.Left, dialog)
}
