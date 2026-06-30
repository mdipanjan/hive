package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/styles"
)

type HelpItem struct {
	Key  string
	Desc string
}

func RenderHelpBar(items []HelpItem) string {
	var parts []string
	for _, item := range items {
		part := styles.HelpKey.Render(item.Key) + styles.HelpDesc.Render(": "+item.Desc)
		parts = append(parts, part)
	}
	return styles.Help.Render("  " + strings.Join(parts, "   "))
}

// RenderHints draws a footer hint bar in the dashboard style: accent key +
// muted label, space-separated, no colon (DESIGN.md §2.2).
func RenderHints(items []HelpItem) string {
	var parts []string
	for _, item := range items {
		parts = append(parts, styles.HelpKey.Render(item.Key)+" "+styles.HelpDesc.Render(item.Desc))
	}
	return strings.Join(parts, "   ")
}

// RenderButtons draws left-aligned buttons. The focused button takes a
// selection fill + accent border; others are neutral/muted (DESIGN.md §3).
// destructive marks the focused button red instead of accent (e.g. Delete).
func RenderButtons(labels []string, selected int, focused bool, opts ...ButtonOption) string {
	cfg := buttonConfig{}
	for _, o := range opts {
		o(&cfg)
	}

	accent := styles.ColorCyan
	if cfg.destructive {
		accent = styles.ColorRed
	}

	var buttons []string
	for i, label := range labels {
		btnStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Foreground(styles.ColorGray).
			BorderForeground(styles.ColorGray)

		if i == selected {
			if focused {
				btnStyle = btnStyle.
					Foreground(styles.ColorWhite).
					Background(styles.ColorDim).
					BorderForeground(accent).
					Bold(true)
			} else {
				btnStyle = btnStyle.Foreground(styles.ColorWhite)
			}
		}
		buttons = append(buttons, btnStyle.Render("  "+label+"  "))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, intersperse(buttons, "  ")...)
}

type buttonConfig struct{ destructive bool }

type ButtonOption func(*buttonConfig)

// Destructive marks the focused button red (used by the delete confirm dialog).
func Destructive() ButtonOption {
	return func(c *buttonConfig) { c.destructive = true }
}

func intersperse(items []string, sep string) []string {
	if len(items) <= 1 {
		return items
	}
	out := make([]string, 0, len(items)*2-1)
	for i, it := range items {
		if i > 0 {
			out = append(out, sep)
		}
		out = append(out, it)
	}
	return out
}

// RenderTextInput draws a single-line field over a hairline underline. When
// focused the value gets a selection highlight followed by a block cursor; the
// rest of the line stays plain (DESIGN.md §3 "Form field").
func RenderTextInput(value string, focused bool, width int) string {
	underline := styles.Dim.Render(strings.Repeat("─", width))

	if focused {
		shown := truncEnd(value, width-1) // leave room for cursor
		val := lipgloss.NewStyle().Foreground(styles.ColorWhite).Background(styles.ColorDim).Bold(true).Render(shown)
		cursor := lipgloss.NewStyle().Background(styles.ColorCyan).Render(" ")
		pad := width - lipgloss.Width(shown) - 1
		if pad < 0 {
			pad = 0
		}
		return val + cursor + strings.Repeat(" ", pad) + "\n" + underline
	}

	shown := truncEnd(value, width)
	val := lipgloss.NewStyle().Foreground(styles.ColorWhite).Render(shown)
	pad := width - lipgloss.Width(shown)
	if pad < 0 {
		pad = 0
	}
	return val + strings.Repeat(" ", pad) + "\n" + underline
}

// truncEnd keeps the start of s and trims the tail to fit max display cells.
func truncEnd(s string, max int) string {
	if max <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= max {
		return s
	}
	return string(runes[:max])
}
