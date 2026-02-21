package components

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"

	"github.com/mdipanjan/hive-v0/internal/styles"
)

var bars = []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

func RenderActivity(width int, cpuHistory []int) string {
	innerWidth := width - 6
	if innerWidth < 20 {
		innerWidth = 20
	}

	var data []int
	if len(cpuHistory) > 0 {
		data = cpuHistory
		if len(data) < innerWidth {
			padding := make([]int, innerWidth-len(data))
			data = append(padding, data...)
		} else if len(data) > innerWidth {
			data = data[len(data)-innerWidth:]
		}
	} else {
		data = fakeData(innerWidth)
	}

	var sb strings.Builder
	for _, value := range data {
		index := value * 7 / 100
		if index > 7 {
			index = 7
		}
		if index < 0 {
			index = 0
		}
		sb.WriteRune(bars[index])
	}
	sparkline := styles.Value.Render(sb.String())

	cpuValue := data[len(data)-1]
	cpuText := "cpu"
	cpuPercent := fmt.Sprintf("%d%%", cpuValue)
	padding := innerWidth - len(cpuText) - len(cpuPercent)
	if padding < 1 {
		padding = 1
	}
	cpuLabel := styles.Label.Render(cpuText) + strings.Repeat(" ", padding) + styles.Stats.Render(cpuPercent)

	content := sparkline + "\n" + cpuLabel
	return styles.Panel.Width(width).Render(content)
}

func fakeData(length int) []int {
	data := make([]int, length)
	for i := range data {
		data[i] = 50 + int(30*math.Sin(float64(i)*0.3)) + rand.Intn(20)
	}
	return data
}

func GetCPUPercent() int {
	percent, err := cpu.Percent(0, false)
	if err != nil || len(percent) == 0 {
		return 0
	}
	return int(percent[0])
}
