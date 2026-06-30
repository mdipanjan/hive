package styles

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Key    string
	Name   string
	Green  lipgloss.Color
	Yellow lipgloss.Color
	Red    lipgloss.Color
	Gray   lipgloss.Color
	Cyan   lipgloss.Color
	White  lipgloss.Color
	Dim    lipgloss.Color
	Bg     lipgloss.Color
	BgDark lipgloss.Color
}

var (
	TokyoNight = Theme{
		Key:    "tokyo-night",
		Name:   "Tokyo Night",
		Green:  lipgloss.Color("#9ece6a"),
		Yellow: lipgloss.Color("#e0af68"),
		Red:    lipgloss.Color("#f7768e"),
		Gray:   lipgloss.Color("#565f89"),
		Cyan:   lipgloss.Color("#7aa2f7"),
		White:  lipgloss.Color("#c0caf5"),
		Dim:    lipgloss.Color("#414868"),
		Bg:     lipgloss.Color("#1a1b26"),
		BgDark: lipgloss.Color("#16161e"),
	}

	Dracula = Theme{
		Key:    "dracula",
		Name:   "Dracula",
		Green:  lipgloss.Color("#50fa7b"),
		Yellow: lipgloss.Color("#f1fa8c"),
		Red:    lipgloss.Color("#ff5555"),
		Gray:   lipgloss.Color("#6272a4"),
		Cyan:   lipgloss.Color("#8be9fd"),
		White:  lipgloss.Color("#f8f8f2"),
		Dim:    lipgloss.Color("#44475a"),
		Bg:     lipgloss.Color("#282a36"),
		BgDark: lipgloss.Color("#21222c"),
	}

	Nord = Theme{
		Key:    "nord",
		Name:   "Nord",
		Green:  lipgloss.Color("#a3be8c"),
		Yellow: lipgloss.Color("#ebcb8b"),
		Red:    lipgloss.Color("#bf616a"),
		Gray:   lipgloss.Color("#4c566a"),
		Cyan:   lipgloss.Color("#88c0d0"),
		White:  lipgloss.Color("#eceff4"),
		Dim:    lipgloss.Color("#3b4252"),
		Bg:     lipgloss.Color("#2e3440"),
		BgDark: lipgloss.Color("#242933"),
	}

	Gruvbox = Theme{
		Key:    "gruvbox",
		Name:   "Gruvbox",
		Green:  lipgloss.Color("#b8bb26"),
		Yellow: lipgloss.Color("#fabd2f"),
		Red:    lipgloss.Color("#fb4934"),
		Gray:   lipgloss.Color("#928374"),
		Cyan:   lipgloss.Color("#83a598"),
		White:  lipgloss.Color("#ebdbb2"),
		Dim:    lipgloss.Color("#504945"),
		Bg:     lipgloss.Color("#282828"),
		BgDark: lipgloss.Color("#1d2021"),
	}

	Catppuccin = Theme{
		Key:    "catppuccin",
		Name:   "Catppuccin",
		Green:  lipgloss.Color("#a6e3a1"),
		Yellow: lipgloss.Color("#f9e2af"),
		Red:    lipgloss.Color("#f38ba8"),
		Gray:   lipgloss.Color("#6c7086"),
		Cyan:   lipgloss.Color("#89dceb"),
		White:  lipgloss.Color("#cdd6f4"),
		Dim:    lipgloss.Color("#45475a"),
		Bg:     lipgloss.Color("#1e1e2e"),
		BgDark: lipgloss.Color("#181825"),
	}

	TokyoStorm = Theme{
		Key:    "tokyo-storm",
		Name:   "Tokyo Storm",
		Green:  lipgloss.Color("#9ece6a"),
		Yellow: lipgloss.Color("#e0af68"),
		Red:    lipgloss.Color("#f7768e"),
		Gray:   lipgloss.Color("#565f89"),
		Cyan:   lipgloss.Color("#7aa2f7"),
		White:  lipgloss.Color("#a9b1d6"),
		Dim:    lipgloss.Color("#3b4261"),
		Bg:     lipgloss.Color("#24283b"),
		BgDark: lipgloss.Color("#1f2335"),
	}

	OneDark = Theme{
		Key:    "one-dark",
		Name:   "One Dark",
		Green:  lipgloss.Color("#98c379"),
		Yellow: lipgloss.Color("#e5c07b"),
		Red:    lipgloss.Color("#e06c75"),
		Gray:   lipgloss.Color("#5c6370"),
		Cyan:   lipgloss.Color("#56b6c2"),
		White:  lipgloss.Color("#abb2bf"),
		Dim:    lipgloss.Color("#3e4451"),
		Bg:     lipgloss.Color("#282c34"),
		BgDark: lipgloss.Color("#21252b"),
	}

	SolarizedDark = Theme{
		Key:    "solarized-dark",
		Name:   "Solarized Dark",
		Green:  lipgloss.Color("#859900"),
		Yellow: lipgloss.Color("#b58900"),
		Red:    lipgloss.Color("#dc322f"),
		Gray:   lipgloss.Color("#586e75"),
		Cyan:   lipgloss.Color("#2aa198"),
		White:  lipgloss.Color("#93a1a1"),
		Dim:    lipgloss.Color("#073642"),
		Bg:     lipgloss.Color("#002b36"),
		BgDark: lipgloss.Color("#001e26"),
	}

	GitHubDark = Theme{
		Key:    "github-dark",
		Name:   "GitHub Dark",
		Green:  lipgloss.Color("#3fb950"),
		Yellow: lipgloss.Color("#d29922"),
		Red:    lipgloss.Color("#f85149"),
		Gray:   lipgloss.Color("#484f58"),
		Cyan:   lipgloss.Color("#58a6ff"),
		White:  lipgloss.Color("#c9d1d9"),
		Dim:    lipgloss.Color("#30363d"),
		Bg:     lipgloss.Color("#0d1117"),
		BgDark: lipgloss.Color("#010409"),
	}

	RosePine = Theme{
		Key:    "rose-pine",
		Name:   "Rosé Pine",
		Green:  lipgloss.Color("#9ccfd8"),
		Yellow: lipgloss.Color("#f6c177"),
		Red:    lipgloss.Color("#eb6f92"),
		Gray:   lipgloss.Color("#6e6a86"),
		Cyan:   lipgloss.Color("#c4a7e7"),
		White:  lipgloss.Color("#e0def4"),
		Dim:    lipgloss.Color("#26233a"),
		Bg:     lipgloss.Color("#191724"),
		BgDark: lipgloss.Color("#1f1d2e"),
	}

	Monokai = Theme{
		Key:    "monokai",
		Name:   "Monokai",
		Green:  lipgloss.Color("#a6e22e"),
		Yellow: lipgloss.Color("#e6db74"),
		Red:    lipgloss.Color("#f92672"),
		Gray:   lipgloss.Color("#75715e"),
		Cyan:   lipgloss.Color("#66d9ef"),
		White:  lipgloss.Color("#f8f8f2"),
		Dim:    lipgloss.Color("#3e3d32"),
		Bg:     lipgloss.Color("#272822"),
		BgDark: lipgloss.Color("#1e1f1c"),
	}

	ZincDark = Theme{
		Key:    "zinc-dark",
		Name:   "Zinc Dark",
		Green:  lipgloss.Color("#4ade80"),
		Yellow: lipgloss.Color("#facc15"),
		Red:    lipgloss.Color("#ef4444"),
		Gray:   lipgloss.Color("#71717a"),
		Cyan:   lipgloss.Color("#22d3ee"),
		White:  lipgloss.Color("#fafafa"),
		Dim:    lipgloss.Color("#3f3f46"),
		Bg:     lipgloss.Color("#18181b"),
		BgDark: lipgloss.Color("#09090b"),
	}
)

var Themes = []Theme{
	TokyoNight,
	TokyoStorm,
	Dracula,
	Nord,
	Gruvbox,
	Catppuccin,
	OneDark,
	SolarizedDark,
	GitHubDark,
	RosePine,
	Monokai,
	ZincDark,
}

var CurrentThemeIndex = 0

func GetCurrentTheme() Theme {
	return Themes[CurrentThemeIndex]
}

func NextTheme() Theme {
	CurrentThemeIndex = (CurrentThemeIndex + 1) % len(Themes)
	ApplyTheme(Themes[CurrentThemeIndex])
	return Themes[CurrentThemeIndex]
}

func GetThemeByKey(key string) Theme {
	for index, theme := range Themes {
		if theme.Key == key {
			CurrentThemeIndex = index
			return theme
		}
	}
	return TokyoNight
}
