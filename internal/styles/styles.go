package styles

import "github.com/charmbracelet/lipgloss"

var (
	ColorGreen  lipgloss.Color
	ColorYellow lipgloss.Color
	ColorGray   lipgloss.Color
	ColorCyan   lipgloss.Color
	ColorWhite  lipgloss.Color
	ColorDim    lipgloss.Color
	ColorBg     lipgloss.Color
	ColorBgDark lipgloss.Color
)

var (
	IconRunning string
	IconWaiting string
	IconIdle    string
)

var (
	OuterBox   lipgloss.Style
	Panel      lipgloss.Style
	PanelTitle lipgloss.Style
	Selected   lipgloss.Style
	Normal     lipgloss.Style
	Dim        lipgloss.Style
	Logo       lipgloss.Style
	Stats      lipgloss.Style
	Help       lipgloss.Style
	HelpKey    lipgloss.Style
	HelpDesc   lipgloss.Style
	Label      lipgloss.Style
	Value      lipgloss.Style
)

func ApplyTheme(t Theme) {
	ColorGreen = t.Green
	ColorYellow = t.Yellow
	ColorGray = t.Gray
	ColorCyan = t.Cyan
	ColorWhite = t.White
	ColorDim = t.Dim
	ColorBg = t.Bg
	ColorBgDark = t.BgDark

	IconRunning = lipgloss.NewStyle().Foreground(ColorGreen).Render("●")
	IconWaiting = lipgloss.NewStyle().Foreground(ColorYellow).Render("◐")
	IconIdle = lipgloss.NewStyle().Foreground(ColorGray).Render("○")

	OuterBox = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorGray).
		Padding(1, 2)

	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorGray).
		Padding(1, 2)

	PanelTitle = lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorCyan)

	Selected = lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorWhite).
		Background(ColorDim)

	Normal = lipgloss.NewStyle().
		Foreground(ColorWhite)

	Dim = lipgloss.NewStyle().
		Foreground(ColorDim)

	Logo = lipgloss.NewStyle().
		Foreground(ColorCyan).
		Bold(true)

	Stats = lipgloss.NewStyle().
		Foreground(ColorGray)

	Help = lipgloss.NewStyle().
		Foreground(ColorDim)

	HelpKey = lipgloss.NewStyle().
		Foreground(ColorCyan)

	HelpDesc = lipgloss.NewStyle().
		Foreground(ColorGray)

	Label = lipgloss.NewStyle().
		Foreground(ColorGray).
		Width(10)

	Value = lipgloss.NewStyle().
		Foreground(ColorWhite)
}

func init() {
	ApplyTheme(TokyoNight)
}
