package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/styles"
)

// RenderHelpPopup draws the keybinding reference, grouped by category for
// scannability, with a branding/close footer line (DESIGN.md §4.7).
func RenderHelpPopup() string {
	const inner = 50 - 2 - 4

	title := styles.PanelTitle.Render("HELP")
	keyStyle := lipgloss.NewStyle().Foreground(styles.ColorCyan).Width(14)
	descStyle := lipgloss.NewStyle().Foreground(styles.ColorWhite)
	groupStyle := lipgloss.NewStyle().Foreground(styles.ColorGray)

	type binding struct{ key, desc string }
	groups := []struct {
		name     string
		bindings []binding
	}{
		{"NAVIGATION", []binding{
			{"↑/k", "Move up"},
			{"↓/j", "Move down"},
		}},
		{"ACTIONS", []binding{
			{"n", "Create new session"},
			{"enter", "Attach to session"},
			{"d", "Delete session"},
			{"t", "Switch theme"},
		}},
		{"GLOBAL", []binding{
			{"?", "Toggle this help"},
			{"q", "Quit"},
		}},
	}

	var sections []string
	for _, g := range groups {
		lines := []string{groupStyle.Render(g.name)}
		for _, b := range g.bindings {
			lines = append(lines, keyStyle.Render(b.key)+descStyle.Render(b.desc))
		}
		sections = append(sections, strings.Join(lines, "\n"))
	}

	divider := styles.Dim.Render(strings.Repeat("─", inner))
	footer := styles.Dim.Render("hive v" + Version + " · press ? or esc to close")

	body := title + "\n\n" + strings.Join(sections, "\n\n") + "\n\n" + divider + "\n\n" + footer
	return styles.Panel.Width(50).Padding(1, 2).Render(body)
}
