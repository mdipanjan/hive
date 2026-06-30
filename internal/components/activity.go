package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	"github.com/mdipanjan/hive/internal/styles"
)

var bars = []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

const sparkWidth = 16

// RenderActivity draws inline cpu + mem sparklines with load-colored values
// (DESIGN.md §4.1). No surrounding panel — it's a bottom status strip.
func RenderActivity(cpuHistory, memHistory []int) string {
	cpuMetric := metric("cpu", cpuHistory)
	memMetric := metric("mem", memHistory)
	return cpuMetric + "    " + memMetric
}

func metric(label string, history []int) string {
	data := fit(history, sparkWidth)

	var sb strings.Builder
	for _, value := range data {
		idx := value * (len(bars) - 1) / 100
		if idx > len(bars)-1 {
			idx = len(bars) - 1
		}
		if idx < 0 {
			idx = 0
		}
		sb.WriteRune(bars[idx])
	}

	current := 0
	if len(data) > 0 {
		current = data[len(data)-1]
	}

	spark := lipgloss.NewStyle().Foreground(styles.ColorCyan).Render(sb.String())
	value := lipgloss.NewStyle().Foreground(loadColor(current)).Render(fmt.Sprintf("%d%%", current))
	return styles.Dim.Render(label) + " " + spark + " " + value
}

func fit(history []int, width int) []int {
	if len(history) == 0 {
		return make([]int, width)
	}
	if len(history) < width {
		return append(make([]int, width-len(history)), history...)
	}
	return history[len(history)-width:]
}

func loadColor(pct int) lipgloss.Color {
	switch {
	case pct >= 80:
		return styles.ColorRed
	case pct >= 50:
		return styles.ColorYellow
	default:
		return styles.ColorGreen
	}
}

func GetCPUPercent() int {
	percent, err := cpu.Percent(0, false)
	if err != nil || len(percent) == 0 {
		return 0
	}
	return int(percent[0])
}

func GetMemPercent() int {
	vm, err := mem.VirtualMemory()
	if err != nil || vm == nil {
		return 0
	}
	return int(vm.UsedPercent)
}
