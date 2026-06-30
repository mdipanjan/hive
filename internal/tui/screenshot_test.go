package tui

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mdipanjan/hive/internal/session"
	"github.com/mdipanjan/hive/internal/styles"
	"github.com/muesli/termenv"
)

// TestGenerateScreenshots renders every TUI screen to an ANSI .txt file under
// ./screenshots so they can be converted to images (e.g. with `freeze`).
//
// It is skipped during normal test runs; enable with:
//
//	HIVE_SCREENSHOTS=1 go test ./internal/tui -run TestGenerateScreenshots
func TestGenerateScreenshots(t *testing.T) {
	if os.Getenv("HIVE_SCREENSHOTS") == "" {
		t.Skip("set HIVE_SCREENSHOTS=1 to generate screenshots")
	}

	// go test has no TTY, so force a color profile before rendering or
	// lipgloss strips all theme colors (monochrome output).
	lipgloss.SetColorProfile(termenv.TrueColor)
	styles.ApplyTheme(styles.GetThemeByKey("tokyo-night"))

	outDir := filepath.Join("..", "..", "screenshots")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		t.Fatalf("mkdir screenshots: %v", err)
	}

	const w, h = 120, 38

	base := func() Model {
		now := time.Now()
		return Model{
			mode:   ModeDashboard,
			width:  w,
			height: h,
			sessions: []session.Session{
				{Name: "hive-dev", Tool: "nvim", Path: "~/2026/projects/hive", Status: session.StatusActive, CreatedAt: now.Add(-3 * time.Hour), LastActivity: now},
				{Name: "mail-work", Tool: "pi", Path: "~/projects/kalgan/mail", Status: session.StatusRunning, CreatedAt: now.Add(-2 * time.Hour), LastActivity: now.Add(-5 * time.Minute)},
				{Name: "seldon", Tool: "claude", Path: "~/projects/seldon", Status: session.StatusReady, CreatedAt: now.Add(-time.Hour), LastActivity: now.Add(-time.Minute)},
				{Name: "file-watch", Tool: "bash", Path: "~/projects/dotfiles", Status: session.StatusCompleted, CreatedAt: now.Add(-30 * time.Minute), LastActivity: now.Add(-10 * time.Minute)},
			},
			cursor:          0,
			app:             NewAppState(),
			cpuUsageHistory: []int{4, 9, 12, 7, 18, 22, 15, 30, 24, 19, 11, 8, 14, 26, 33, 20},
		}
	}

	type shot struct {
		name  string
		build func() Model
	}

	shots := []shot{
		{"01-dashboard", func() Model { return base() }},
		{"02-search", func() Model {
			m := base()
			m.searchInput = newSearchInput()
			m.searchInput.SetValue("ma")
			m.app.Search()
			m.searchResults = m.getIndices("ma")
			m.searchCursor = 0
			return m
		}},
		{"03-switch", func() Model {
			m := base()
			m.mode = ModeSwitch
			m.searchInput = newSearchInput()
			m.app.Search()
			m.searchResults = m.getIndices("")
			m.searchCursor = 1
			return m
		}},
		{"04-new-form", func() Model {
			m := base()
			m.app.StartNewSession()
			m.form = newForm()
			m.form.Name = "api-worker"
			m.form.Focus = FocusName
			return m
		}},
		{"05-filepicker", func() Model {
			m := base()
			m.app.StartNewSession()
			m.form = newForm()
			fp := newFilePicker()
			fp, _ = fp.Update(tea.WindowSizeMsg{Width: w, Height: h})
			if cmd := fp.Init(); cmd != nil {
				if msg := cmd(); msg != nil {
					fp, _ = fp.Update(msg)
				}
			}
			m.form.FilePicker = fp
			m.app.PickPath()
			return m
		}},
		{"06-delete-confirm", func() Model {
			m := base()
			m.cursor = 1
			m.app.ConfirmDelete()
			m.deleteButton = 1
			return m
		}},
		{"07-help", func() Model {
			m := base()
			m.app.ShowHelp()
			return m
		}},
	}

	for _, s := range shots {
		out := RenderView(s.build())
		path := filepath.Join(outDir, s.name+".txt")
		if err := os.WriteFile(path, []byte(out+"\n"), 0o644); err != nil {
			t.Fatalf("write %s: %v", path, err)
		}
		t.Logf("wrote %s", path)
	}
}
