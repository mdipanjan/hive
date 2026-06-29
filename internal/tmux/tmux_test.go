package tmux

import (
	"reflect"
	"strings"
	"testing"
)

func TestRunnerCommand(t *testing.T) {
	got := runnerCommand("work", "pi", "/tmp/project")
	wantParts := []string{
		" run-session",
		"--name 'work'",
		"--tool 'pi'",
		"--path '/tmp/project'",
	}

	for _, part := range wantParts {
		if !strings.Contains(got, part) {
			t.Fatalf("runnerCommand() = %q, expected to contain %q", got, part)
		}
	}
}

func TestRunnerCommandQuotesArgs(t *testing.T) {
	got := runnerCommand("work's", "nvim", "/tmp/my project")
	wantParts := []string{
		"--name 'work'\\''s'",
		"--tool 'nvim'",
		"--path '/tmp/my project'",
	}

	for _, part := range wantParts {
		if !strings.Contains(got, part) {
			t.Fatalf("runnerCommand() = %q, expected to contain %q", got, part)
		}
	}
}

func TestEditorLayoutCommands(t *testing.T) {
	got := editorLayoutCommands("work", "/tmp/project")
	want := [][]string{
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
