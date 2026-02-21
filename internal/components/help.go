package components

import (
	"fmt"

	"github.com/mdipanjan/hive-v0/internal/styles"
)

func RenderHelp() string {
	themeName := styles.GetCurrentTheme().Name
	helpText := fmt.Sprintf("  n: new   enter: attach   d: delete   t: theme [%s]   q: quit", themeName)
	return styles.Help.Render(helpText)
}
