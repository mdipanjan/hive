package tmux

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/mdipanjan/hive/internal/logger"
	"github.com/mdipanjan/hive/internal/session"
	"github.com/mdipanjan/hive/internal/state"
)

func List() ([]session.Session, error) {
	cmd := exec.Command("tmux", "list-sessions", "-F", "#{session_name}|#{session_created}|#{session_activity}|#{session_attached}|#{session_path}|#{pane_current_command}")
	output, err := cmd.Output()

	logger.Log.Debug("tmux list-sessions", "output", string(output), "err", err)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	sessions := make([]session.Session, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) < 4 {
			continue
		}

		name := parts[0]
		created, _ := strconv.ParseInt(parts[1], 10, 64)
		activity, _ := strconv.ParseInt(parts[2], 10, 64)
		attached, _ := strconv.ParseInt(parts[3], 10, 64)

		path := ""
		tool := ""
		if len(parts) >= 5 {
			path = parts[4]
		}
		if len(parts) >= 6 {
			tool = parts[5]
		}

		createdAt := time.Unix(created, 0)
		lastActivity := time.Unix(activity, 0)

		meta, metaErr := state.ReadMetadata(name)
		if metaErr == nil {
			tool = meta.Tool
			path = meta.Path
			createdAt = meta.CreatedAt
		}

		runtime, runtimeErr := state.ReadRuntime(name)
		status := statusFromRuntime(runtime, runtimeErr, attached, lastActivity)

		sessions = append(sessions, session.Session{
			Name:         name,
			Tool:         tool,
			Path:         path,
			Status:       status,
			CreatedAt:    createdAt,
			LastActivity: lastActivity,
		})
	}

	logger.Log.Debug("parsed sessions", "count", len(sessions))
	return sessions, nil
}

func statusFromRuntime(runtime state.Runtime, err error, attached int64, lastActivity time.Time) session.Status {
	if attached > 0 {
		return session.StatusActive
	}

	if err != nil {
		if time.Since(lastActivity) < 5*time.Minute {
			return session.StatusRunning
		}
		return session.StatusIdle
	}

	switch runtime.State {
	case state.StateStarting, state.StateRunning:
		return session.StatusRunning
	case state.StateReady:
		return session.StatusReady
	case state.StateCompleted:
		return session.StatusCompleted
	case state.StateFailed:
		return session.StatusFailed
	case state.StateUnknown:
		return session.StatusUnknown
	default:
		return session.StatusIdle
	}
}

func Create(name, tool, path string) error {
	logger.Log.Debug("creating tmux session", "name", name, "tool", tool, "path", path)

	if err := state.WriteMetadata(state.Metadata{Name: name, Tool: tool, Path: path, CreatedAt: time.Now().UTC()}); err != nil {
		return err
	}

	if err := state.WriteRuntime(name, state.Runtime{State: state.StateStarting, StartedAt: time.Now().UTC()}); err != nil {
		return err
	}

	if tool == "nvim" || tool == "vim" {
		return createEditorSession(name, tool, path)
	}

	err := exec.Command("tmux", "new-session", "-d", "-s", name, "-c", path, runnerCommand(name, tool, path)).Run()
	if err != nil {
		logger.Log.Error("tmux create failed", "err", err)
		return err
	}

	configureSession(name)
	return nil
}

func createEditorSession(name, tool, path string) error {
	err := exec.Command("tmux", "new-session", "-d", "-s", name, "-c", path, runnerCommand(name, tool, path)).Run()
	if err != nil {
		logger.Log.Error("tmux editor session create failed", "err", err)
		return err
	}

	for _, args := range editorLayoutCommands(name, path) {
		if err := exec.Command("tmux", args...).Run(); err != nil {
			logger.Log.Error("tmux editor layout failed", "args", args, "err", err)
			return err
		}
	}

	configureSession(name)
	return nil
}

func runnerCommand(name, tool, path string) string {
	executable, err := os.Executable()
	if err != nil {
		executable = "hive"
	}

	command := shellQuote(executable) + " run-session --name " + shellQuote(name) + " --tool " + shellQuote(tool) + " --path " + shellQuote(path)
	if stateDir := os.Getenv("HIVE_STATE_DIR"); stateDir != "" {
		command = "HIVE_STATE_DIR=" + shellQuote(stateDir) + " " + command
	}
	return command
}

func shellQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "'\\''") + "'"
}

func editorLayoutCommands(name, path string) [][]string {
	return [][]string{
		{"split-window", "-h", "-p", "35", "-t", name + ":0", "-c", path},
		{"select-pane", "-t", name + ":0.0"},
	}
}

func configureSession(name string) {
	for _, args := range sessionOptions(name) {
		exec.Command("tmux", args...).Run()
	}
}

func sessionOptions(name string) [][]string {
	return [][]string{
		{"set-option", "-t", name, "status", "off"},
		{"set-option", "-t", name, "mouse", "on"},
		{"set-option", "-t", name, "pane-border-style", "fg=#45475a"},
		{"set-option", "-t", name, "pane-active-border-style", "fg=#a6e3a1"},
		{"set-option", "-t", name, "window-style", "bg=#1e1e2e"},
		{"set-option", "-t", name, "window-active-style", "bg=#1e1e2e"},
	}
}

func Attach(name string) error {
	cmd := exec.Command("tmux", "attach", "-t", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func IsInsideTmux() bool {
	return os.Getenv("TMUX") != ""
}

func AttachCmd(name string) *exec.Cmd {
	if IsInsideTmux() {
		logger.Log.Debug("using switch-client (inside tmux)", "target", name)
		return exec.Command("tmux", "switch-client", "-t", name)
	}
	logger.Log.Debug("using attach (outside tmux)", "target", name)
	return exec.Command("tmux", "attach", "-t", name)
}

func Kill(name string) error {
	logger.Log.Debug("killing tmux session", "name", name)
	err := exec.Command("tmux", "kill-session", "-t", name).Run()
	if err != nil {
		logger.Log.Error("tmux kill failed", "err", err)
	}
	if stateErr := state.DeleteSessionState(name); stateErr != nil {
		logger.Log.Error("session state delete failed", "err", stateErr)
		if err == nil {
			err = stateErr
		}
	}
	return err
}
