package tmux

import (
	"reflect"
	"testing"
)

func TestEditorLayoutCommands(t *testing.T) {
	got := editorLayoutCommands("work", "/tmp/project")
	want := [][]string{
		{"send-keys", "-t", "work:0.0", "nvim .; tmux detach-client -s work", "C-m"},
		{"split-window", "-h", "-p", "35", "-t", "work:0", "-c", "/tmp/project"},
		{"select-pane", "-t", "work:0.0"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("editorLayoutCommands() = %#v, want %#v", got, want)
	}
}

func TestSessionOptions(t *testing.T) {
	got := sessionOptions("work")
	want := [][]string{
		{"set-option", "-t", "work", "status", "off"},
		{"set-option", "-t", "work", "mouse", "on"},
		{"set-option", "-t", "work", "pane-border-style", "fg=#45475a"},
		{"set-option", "-t", "work", "pane-active-border-style", "fg=#a6e3a1"},
		{"set-option", "-t", "work", "window-style", "bg=#1e1e2e"},
		{"set-option", "-t", "work", "window-active-style", "bg=#1e1e2e"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("sessionOptions() = %#v, want %#v", got, want)
	}
}
