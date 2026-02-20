package tui

func RenderView(m Model) string {
	var s string
	// title
	s += TitleStyle.Render("hive") + "\n"

	for index, session := range m.sessions {
		// Status icon
		icon := StatusIdle
		// Session name
		line := icon + " " + session.Name
		// Highlight if selected
		if m.cursor == index {
			line = SelectedStyle.Render(line + " ←")
		} else {
			line = NormalStyle.Render(line)
		}
		s += line + "\n"
	}
	// 3. Empty state
	if len(m.sessions) == 0 {
		s += "  No sessions. Press 'n' to create one.\n"
	}

	// 4. Help bar
	s += HelpStyle.Render("\n  n: new   enter: attach   d: delete   q: quit")

	return BoxStyle.Render(s)
}
