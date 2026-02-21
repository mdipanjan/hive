package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdipanjan/hive-v0/internal/components"
	"github.com/mdipanjan/hive-v0/internal/config"
	"github.com/mdipanjan/hive-v0/internal/styles"
	"github.com/mdipanjan/hive-v0/internal/tui"
)

var Version = "dev"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version", "version":
			fmt.Println("hive", Version)
			return
		case "-h", "--help", "help":
			printHelp()
			return
		}
	}

	components.Version = Version

	cfg := config.Load()
	theme := styles.GetThemeByKey(cfg.Theme)
	styles.ApplyTheme(theme)

	p := tea.NewProgram(tui.New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`hive - lightweight TUI for managing tmux sessions

Usage:
  hive              Start the TUI
  hive -v           Show version
  hive -h           Show this help

Keys:
  n         New session
  enter     Attach to session
  d         Delete session
  t         Switch theme
  q         Quit`)
}
