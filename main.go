package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdipanjan/hive-v0/internal/tui"
)

func main() {
	m := tui.New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
