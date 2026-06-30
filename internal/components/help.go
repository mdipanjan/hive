package components

import (
	"strings"

	"github.com/mdipanjan/hive/internal/styles"
)

var Version = "dev"

// RenderHelp draws the dashboard footer hint bar: lowercase keys in the accent
// color, labels in muted, space-separated (DESIGN.md §2.2, §4.1).
func RenderHelp() string {
	themeName := styles.GetCurrentTheme().Name

	hint := func(key, label string) string {
		return styles.HelpKey.Render(key) + " " + styles.HelpDesc.Render(label)
	}

	items := []string{
		hint("n", "new"),
		hint("⏎", "attach"),
		hint("d", "delete"),
		hint("/", "search"),
		hint("t", "theme "+styles.Dim.Render("· "+strings.ToLower(themeName))),
		hint("?", "help"),
		hint("q", "quit"),
	}

	return strings.Join(items, "   ")
}
