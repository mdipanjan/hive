package components

import (
	"github.com/mdipanjan/hive/internal/styles"
)

func RenderDeleteConfirm(sessionName string, selectedButton int) string {
	title := styles.PanelTitle.Render("DELETE SESSION")
	message := styles.Normal.Render("Delete \"" + sessionName + "\"?")
	var opts []ButtonOption
	if selectedButton == 0 {
		opts = append(opts, Destructive())
	}
	buttons := RenderButtons([]string{"Yes", "No"}, selectedButton, true, opts...)

	content := title + "\n\n" + message + "\n\n" + buttons

	return styles.Panel.Width(44).Padding(1, 2).Render(content)
}
