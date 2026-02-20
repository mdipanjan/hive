package styles

import "github.com/charmbracelet/lipgloss"

// Tokyo Night color palette
var (
	ColorGreen  = lipgloss.Color("#9ece6a")
	ColorYellow = lipgloss.Color("#e0af68")
	ColorGray   = lipgloss.Color("#565f89")
	ColorCyan   = lipgloss.Color("#7aa2f7")
	ColorWhite  = lipgloss.Color("#c0caf5")
	ColorDim    = lipgloss.Color("#414868")
	ColorBg     = lipgloss.Color("#1a1b26")
	ColorBgDark = lipgloss.Color("#16161e")
)

// Status icons (pre-rendered)
var (
	IconRunning = lipgloss.NewStyle().Foreground(ColorGreen).Render("●")
	IconWaiting = lipgloss.NewStyle().Foreground(ColorYellow).Render("◐")
	IconIdle    = lipgloss.NewStyle().Foreground(ColorGray).Render("○")
)

// Panel styles
var (
	// Outer container
	OuterBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorGray).
			Padding(1, 2)

	// Panel with title (SESSIONS, SELECTED, etc.)
	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorGray).
		Padding(1, 2)

	// Panel title
	PanelTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorCyan)

	// Selected row in list
	Selected = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorWhite).
			Background(ColorDim)

	// Normal row in list
	Normal = lipgloss.NewStyle().
		Foreground(ColorWhite)

	// Dimmed/secondary text
	Dim = lipgloss.NewStyle().
		Foreground(ColorDim)

	// Logo style
	Logo = lipgloss.NewStyle().
		Foreground(ColorCyan).
		Bold(true)

	// Stats text (agents: 4, active: 2)
	Stats = lipgloss.NewStyle().
		Foreground(ColorGray)

	// Help bar
	Help = lipgloss.NewStyle().
		Foreground(ColorDim)

	// Help key
	HelpKey = lipgloss.NewStyle().
		Foreground(ColorCyan)

	// Help description
	HelpDesc = lipgloss.NewStyle().
			Foreground(ColorGray)

	// Label (NAME, TOOL, PATH, etc.)
	Label = lipgloss.NewStyle().
		Foreground(ColorGray).
		Width(10)

	// Value
	Value = lipgloss.NewStyle().
		Foreground(ColorWhite)
)
