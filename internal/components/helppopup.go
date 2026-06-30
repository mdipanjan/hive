package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/styles"
)

// RenderHelpPopup draws the keybinding reference, grouped by category for
// scannability (DESIGN.md §4.7).
func RenderHelpPopup() string {
	title := styles.PanelTitle.Render("HELP")

	keyStyle := lipgloss.NewStyle().Foreground(styles.ColorCyan).Width(10)
	descStyle := lipgloss.NewStyle().Foreground(styles.ColorWhite)
	groupStyle := styles.Dim

	type binding struct{ key, desc string }
	groups := []struct {
		name     string
		bindings []binding
	}{
		{"navigation", []binding{
			{"↑/k", "Move up"},
			{"↓/j", "Move down"},
			{"/", "Search sessions"},
		}},
		{"actions", []binding{
			{"n", "Create new session"},
			{"enter", "Attach to session"},
			{"d", "Delete session"},
			{"t", "Switch theme"},
		}},
		{"global", []binding{
			{"?", "Toggle this help"},
			{"q", "Quit"},
		}},
	}

	var sections []string
	for _, g := range groups {
		lines := []string{groupStyle.Render("  " + g.name)}
		for _, b := range g.bindings {
			lines = append(lines, "    "+keyStyle.Render(b.key)+descStyle.Render(b.desc))
		}
		sections = append(sections, strings.Join(lines, "\n"))
	}

	body := title + "\n\n" + strings.Join(sections, "\n\n")
	return styles.Panel.Width(44).Padding(1, 2).Render(body)
}
