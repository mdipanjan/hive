package tui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors (Tokyo Night theme)
	colorGreen  = lipgloss.Color("#9ece6a")
	colorYellow = lipgloss.Color("#e0af68")
	colorGray   = lipgloss.Color("#565f89")
	colorCyan   = lipgloss.Color("#7aa2f7")
	colorWhite  = lipgloss.Color("#c0caf5")
	colorDim    = lipgloss.Color("#414868")
	colorBg     = lipgloss.Color("#1a1b26")

	// Title
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorCyan).
			MarginBottom(1)

	// Box border
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorGray).
			Padding(1, 2)

	// Session list items
	SelectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorWhite).
			Background(lipgloss.Color("236"))

	NormalStyle = lipgloss.NewStyle().
			Foreground(colorWhite)

	// Status icons
	StatusRunning = lipgloss.NewStyle().
			Foreground(colorGreen).
			Render("●")

	StatusWaiting = lipgloss.NewStyle().
			Foreground(colorYellow).
			Render("◐")

	StatusIdle = lipgloss.NewStyle().
			Foreground(colorGray).
			Render("○")

	// Help bar
	HelpStyle = lipgloss.NewStyle().
			Foreground(colorDim).
			MarginTop(1)
)
