package tui

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mdipanjan/hive/internal/components"
)

// dirPicker is a keyboard-driven, two-column directory chooser. Directories sort
// first, then files; hidden entries are skipped. enter selects the highlighted
// directory as the path; →/l descend, ←/h ascend.
type dirPicker struct {
	dir     string
	entries []components.DirEntry
	cursor  int
}

func newDirPicker(start string) dirPicker {
	if start == "" {
		start, _ = os.UserHomeDir()
	}
	p := dirPicker{dir: start}
	p.read()
	return p
}

func (p *dirPicker) read() {
	p.cursor = 0
	p.entries = nil

	ents, err := os.ReadDir(p.dir)
	if err != nil {
		return
	}

	var dirs, files []components.DirEntry
	for _, e := range ents {
		name := e.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		if e.IsDir() {
			dirs = append(dirs, components.DirEntry{Name: name, IsDir: true})
		} else {
			files = append(files, components.DirEntry{Name: name, IsDir: false})
		}
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name < dirs[j].Name })
	sort.Slice(files, func(i, j int) bool { return files[i].Name < files[j].Name })
	p.entries = append(dirs, files...)
}

func (p *dirPicker) moveUp() {
	if p.cursor > 0 {
		p.cursor--
	}
}

func (p *dirPicker) moveDown() {
	if p.cursor < len(p.entries)-1 {
		p.cursor++
	}
}

// moveColumn jumps a full column (DirPickerRows entries) left/right.
func (p *dirPicker) moveColumn(delta int) {
	next := p.cursor + delta*components.DirPickerRows
	if next >= 0 && next < len(p.entries) {
		p.cursor = next
	}
}

func (p *dirPicker) descend() {
	if len(p.entries) == 0 {
		return
	}
	e := p.entries[p.cursor]
	if !e.IsDir {
		return
	}
	p.dir = filepath.Join(p.dir, e.Name)
	p.read()
}

func (p *dirPicker) ascend() {
	parent := filepath.Dir(p.dir)
	if parent != p.dir {
		p.dir = parent
		p.read()
	}
}

// selected returns the highlighted directory path, if the highlight is a dir.
func (p *dirPicker) selected() (string, bool) {
	if len(p.entries) == 0 {
		return "", false
	}
	e := p.entries[p.cursor]
	if !e.IsDir {
		return "", false
	}
	return filepath.Join(p.dir, e.Name), true
}

func (p dirPicker) view() string {
	return components.RenderDirPicker(p.entries, p.cursor)
}
