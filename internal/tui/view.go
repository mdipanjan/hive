package tui

import "github.com/mdipanjan/hive-v0/internal/styles"

func RenderView(m Model) string {
	var s string
	// title
	s += styles.PanelTitle.Render("hive") + "\n"

	for index, session := range m.sessions {
		// Status icon
		icon := styles.IconIdle
		// Session name
		line := icon + " " + session.Name
		// Highlight if selected
		if m.cursor == index {
			line = styles.Selected.Render(line + " ←")
		} else {
			line = styles.Normal.Render(line)
		}
		s += line + "\n"
	}
	// 3. Empty state
	if len(m.sessions) == 0 {
		s += "  No sessions. Press 'n' to create one.\n"
	}

	// 4. Help bar
	s += styles.Help.Render("\n  n: new   enter: attach   d: delete   q: quit")

	return styles.OuterBox.Render(s)
}
