package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive-v0/internal/styles"
)

type FormData struct {
	Tool   int
	Path   string
	Name   string
	Focus  int
	Button int
}

func RenderNewDialog(tools []string, form FormData) string {
	var content strings.Builder

	content.WriteString(styles.PanelTitle.Render("NEW SESSION"))
	content.WriteString("\n\n")

	toolLabel := styles.Label.Render("TOOL")
	toolValue := renderToolSelector(tools, form.Tool, form.Focus == 0)
	content.WriteString(toolLabel + toolValue + "\n\n")

	pathLabel := styles.Label.Render("PATH")
	pathValue := renderTextInput(form.Path, form.Focus == 1)
	content.WriteString(pathLabel + pathValue + "\n\n")

	nameLabel := styles.Label.Render("NAME")
	namePlaceholder := form.Name
	if namePlaceholder == "" {
		namePlaceholder = "(auto-generate)"
	}
	nameValue := renderTextInput(namePlaceholder, form.Focus == 2)
	content.WriteString(nameLabel + nameValue + "\n\n")

	content.WriteString(renderButtons(form.Button, form.Focus == 3))

	dialogStyle := styles.Panel.Copy().Width(70).Padding(1, 2)
	return dialogStyle.Render(content.String())
}

func renderToolSelector(tools []string, selected int, focused bool) string {
	var parts []string

	for i, tool := range tools {
		var icon string
		if i == selected {
			icon = styles.IconRunning + " "
		} else {
			icon = styles.IconIdle + " "
		}

		if i == selected && focused {
			parts = append(parts, styles.Selected.Render(icon+tool))
		} else if i == selected {
			parts = append(parts, styles.Value.Render(icon+tool))
		} else {
			parts = append(parts, styles.Dim.Render(icon+tool))
		}
	}

	return strings.Join(parts, "  ")
}

func renderTextInput(value string, focused bool) string {
	width := 30
	displayValue := value

	if len(displayValue) < width {
		displayValue = displayValue + strings.Repeat(" ", width-len(displayValue))
	} else if len(displayValue) > width {
		displayValue = displayValue[:width]
	}

	inputStyle := lipgloss.NewStyle().Foreground(styles.ColorWhite)

	if focused {
		inputStyle = inputStyle.Background(styles.ColorDim).Bold(true)
		if len(value) < width {
			displayValue = value + "█" + strings.Repeat(" ", width-len(value)-1)
		}
	}

	text := inputStyle.Render(displayValue)
	underline := styles.Dim.Render(strings.Repeat("─", width))

	return text + "\n" + strings.Repeat(" ", 10) + underline
}

func renderButtons(selected int, focused bool) string {
	createBtn := "  Create  "
	cancelBtn := "  Cancel  "

	createStyle := lipgloss.NewStyle().
		Foreground(styles.ColorWhite).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.ColorGray)

	cancelStyle := createStyle.Copy()

	if focused {
		if selected == 0 {
			createStyle = createStyle.
				Background(styles.ColorCyan).
				BorderForeground(styles.ColorCyan).
				Bold(true)
		} else {
			cancelStyle = cancelStyle.
				Background(styles.ColorGray).
				BorderForeground(styles.ColorWhite).
				Bold(true)
		}
	}

	create := createStyle.Render(createBtn)
	cancel := cancelStyle.Render(cancelBtn)

	buttons := lipgloss.JoinHorizontal(lipgloss.Center, create, "  ", cancel)
	return lipgloss.PlaceHorizontal(70, lipgloss.Center, buttons)
}
