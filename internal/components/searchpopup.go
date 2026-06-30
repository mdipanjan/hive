package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/session"
	"github.com/mdipanjan/hive/internal/styles"
)

func RenderSearchPopup(inputView string, query string, sessions []session.Session, results []int, cursor int) string {
	return RenderSearchPopupTitled("SEARCH", inputView, query, sessions, results, cursor)
}

func RenderSearchPopupTitled(titleText, inputView, query string, sessions []session.Session, results []int, cursor int) string {
	title := styles.PanelTitle.Render(titleText)
	// Single prompt marker only (DESIGN.md §4.2): the textinput already renders
	// its own "❯ " prompt, so don't prepend another.
	input := inputView

	maxResults := 6
	var resultsList string

	if len(results) == 0 && query != "" {
		resultsList = styles.Dim.Render("  No results") + "\n"
		for i := 1; i < maxResults; i++ {
			resultsList += "\n"
		}
	} else {
		for i := 0; i < maxResults; i++ {
			if i < len(results) {
				idx := results[i]
				name := TruncateMiddle(sessions[idx].Name, 28)
				icon := GetStatusIcon(sessions[idx].Status) // real status glyph, always

				nameStyled := styles.Normal.Render(name)
				if i == cursor {
					nameStyled = lipgloss.NewStyle().Foreground(styles.ColorCyan).Bold(true).Render(name)
				}
				resultsList += icon + " " + nameStyled + "\n"
			} else {
				resultsList += "\n"
			}
		}
	}

	footer := styles.Dim.Render(fmt.Sprintf("%d / %d", len(results), len(sessions)))
	content := title + "\n\n" + input + "\n\n" + resultsList + footer

	return styles.Panel.Width(40).Padding(1, 2).Render(content)
}
