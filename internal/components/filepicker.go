package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/styles"
)

// DirEntry is one row in the directory picker.
type DirEntry struct {
	Name  string
	IsDir bool
}

const (
	DirPickerRows  = 10 // visible rows per column
	DirPickerCols  = 2
	dirPickerWidth = 72
)

// RenderDirPicker draws the two-column directory chooser: directories in the
// foreground color, files muted, the highlighted entry prefixed "›" and tinted
// green, a hairline divider, and a dir/file legend (DESIGN.md §4.5).
func RenderDirPicker(entries []DirEntry, cursor int) string {
	title := styles.PanelTitle.Render("SELECT DIRECTORY") + "\n\n"
	inner := dirPickerWidth - 2 - 4
	colW := inner / DirPickerCols

	perPage := DirPickerRows * DirPickerCols
	start := 0
	if cursor >= perPage {
		start = (cursor / perPage) * perPage
	}
	end := start + perPage
	if end > len(entries) {
		end = len(entries)
	}
	page := entries[start:end]
	local := cursor - start

	var lines []string
	if len(page) == 0 {
		lines = append(lines, styles.Dim.Render("(empty)"))
	}
	for r := 0; r < DirPickerRows && len(page) > 0; r++ {
		left := dirCell(page, r, local, colW)
		right := dirCell(page, r+DirPickerRows, local, colW)
		lines = append(lines, left+right)
	}

	divider := styles.Dim.Render(strings.Repeat("─", inner))
	legend := dirLegend()
	content := title + strings.Join(lines, "\n") + "\n\n" + divider + "\n\n" + legend

	return styles.Panel.Width(dirPickerWidth).Padding(1, 2).Render(content)
}

func dirCell(page []DirEntry, idx, cursor, colW int) string {
	if idx < 0 || idx >= len(page) {
		return strings.Repeat(" ", colW)
	}
	e := page[idx]
	name := truncEnd(e.Name, colW-2)

	var s string
	switch {
	case idx == cursor:
		s = lipgloss.NewStyle().Foreground(styles.ColorGreen).Render("› " + name)
	case e.IsDir:
		s = lipgloss.NewStyle().Foreground(styles.ColorWhite).Render("  " + name)
	default:
		s = styles.Dim.Render("  " + name)
	}

	if pad := colW - lipgloss.Width(s); pad > 0 {
		s += strings.Repeat(" ", pad)
	}
	return s
}

func dirLegend() string {
	bar := func(c lipgloss.Color) string {
		return lipgloss.NewStyle().Foreground(c).Render("▎")
	}
	return bar(styles.ColorCyan) + " " + styles.Dim.Render("directory") +
		"   " + bar(styles.ColorGray) + " " + styles.Dim.Render("file")
}
