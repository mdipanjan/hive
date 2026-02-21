package components

import (
	"fmt"

	"github.com/mdipanjan/hive-v0/internal/session"
	"github.com/mdipanjan/hive-v0/internal/styles"
)

func RenderSearchPopup(inputView string, query string, sessions []session.Session, results []int, cursor int) string {
	title := styles.PanelTitle.Render("SEARCH")
	input := "› " + inputView

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
				icon := GetStatusIcon(sessions[idx].Status)
				line := icon + " " + name

				if i == cursor {
					line = styles.Logo.Render("● " + name)
				} else {
					line = styles.Normal.Render(line)
				}
				resultsList += line + "\n"
			} else {
				resultsList += "\n"
			}
		}
	}

	footer := styles.Dim.Render(fmt.Sprintf("%d / %d", len(results), len(sessions)))
	content := title + "\n\n" + input + "\n\n" + resultsList + footer

	return styles.Panel.Width(40).Padding(1, 2).Render(content)
}
