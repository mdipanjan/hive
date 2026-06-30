package components

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/styles"
)

// RenderTitleBar draws faux window chrome: three traffic-light dots followed by
// a muted "hive · session manager" caption (DESIGN.md §4.1).
func RenderTitleBar() string {
	dot := func(c lipgloss.Color) string {
		return lipgloss.NewStyle().Foreground(c).Render("●")
	}
	lights := dot(styles.ColorRed) + " " + dot(styles.ColorYellow) + " " + dot(styles.ColorGreen)
	caption := styles.Dim.Render("hive · session manager")
	return lights + "   " + caption
}
