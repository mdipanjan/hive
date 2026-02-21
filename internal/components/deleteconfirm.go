package components

import (
	"github.com/mdipanjan/hive-v0/internal/styles"
)

func RenderDeleteConfirm(sessionName string, selectedButton int) string {
	title := styles.PanelTitle.Render("DELETE SESSION")
	message := styles.Normal.Render("Delete \"" + sessionName + "\"?")
	buttons := RenderButtons([]string{"Yes", "No"}, selectedButton, true, 40)

	content := title + "\n\n" + message + "\n\n" + buttons

	return styles.Panel.Width(44).Padding(1, 2).Render(content)
}
