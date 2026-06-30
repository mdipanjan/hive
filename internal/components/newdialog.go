package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/styles"
)

type FormData struct {
	Tool   int
	Path   string
	Name   string
	Focus  int
	Button int
}

const (
	dialogWidth    = 70
	dialogLabelCol = 10
	dialogFieldW   = 48
)

func RenderNewDialog(tools []string, form FormData) string {
	label := func(text string) string {
		return styles.Dim.Width(dialogLabelCol).Render(text)
	}
	row := func(text, field string) string {
		return lipgloss.JoinHorizontal(lipgloss.Top, label(text), field)
	}

	var content strings.Builder
	content.WriteString(styles.PanelTitle.Render("NEW SESSION"))
	content.WriteString("\n\n")

	content.WriteString(row("tool", renderToolSelector(tools, form.Tool, form.Focus == 0)))
	content.WriteString("\n\n")

	content.WriteString(row("path", RenderTextInput(form.Path, form.Focus == 1, dialogFieldW)))
	content.WriteString("\n\n")

	content.WriteString(row("name", RenderTextInput(form.Name, form.Focus == 2, dialogFieldW)))
	content.WriteString("\n\n")

	buttons := RenderButtons([]string{"Create", "Cancel"}, form.Button, form.Focus == 3)
	content.WriteString(row("", buttons))

	return styles.Panel.Width(dialogWidth).Padding(1, 2).Render(content.String())
}

func renderToolSelector(tools []string, selected int, focused bool) string {
	var parts []string
	for i, tool := range tools {
		if i == selected {
			dot := lipgloss.NewStyle().Foreground(styles.ColorYellow).Render("●")
			name := lipgloss.NewStyle().Foreground(styles.ColorWhite).Bold(true).Render(tool)
			parts = append(parts, dot+" "+name)
		} else {
			dot := styles.Dim.Render("○")
			name := styles.Dim.Render(tool)
			parts = append(parts, dot+" "+name)
		}
	}
	_ = focused
	return strings.Join(parts, "   ")
}
