package tmux

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/mdipanjan/hive/internal/logger"
	"github.com/mdipanjan/hive/internal/session"
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

		status := session.StatusIdle
		if attached > 0 {
			status = session.StatusRunning
		} else {
			lastActivity := time.Unix(activity, 0)
			if time.Since(lastActivity) < 5*time.Minute {
				status = session.StatusWaiting
			}
		}

		sessions = append(sessions, session.Session{
			Name:         name,
			Tool:         tool,
			Path:         path,
			Status:       status,
			CreatedAt:    time.Unix(created, 0),
			LastActivity: time.Unix(activity, 0),
		})
	}

	logger.Log.Debug("parsed sessions", "count", len(sessions))
	return sessions, nil
}

func Create(name, tool, path string) error {
	logger.Log.Debug("creating tmux session", "name", name, "tool", tool, "path", path)

	if tool == "nvim" || tool == "vim" {
		return createEditorSession(name, path)
	}

	err := exec.Command("tmux", "new-session", "-d", "-s", name, "-c", path, tool).Run()
	if err != nil {
		logger.Log.Error("tmux create failed", "err", err)
		return err
	}

	configureSession(name)
	return nil
}

func createEditorSession(name, path string) error {
	err := exec.Command("tmux", "new-session", "-d", "-s", name, "-c", path).Run()
	if err != nil {
		logger.Log.Error("tmux editor session create failed", "err", err)
		return err
	}

	commands := [][]string{
		{"send-keys", "-t", name + ":0.0", "nvim .; tmux detach-client -s " + name, "C-m"},
		{"split-window", "-h", "-p", "35", "-t", name + ":0", "-c", path},
		{"select-pane", "-t", name + ":0.0"},
	}

	for _, args := range commands {
		if err := exec.Command("tmux", args...).Run(); err != nil {
			logger.Log.Error("tmux editor layout failed", "args", args, "err", err)
			return err
		}
	}

	configureSession(name)
	return nil
}

func configureSession(name string) {
	options := [][]string{
		{"set-option", "-t", name, "status", "off"},
		{"set-option", "-t", name, "mouse", "on"},
		{"set-option", "-t", name, "pane-border-style", "fg=#45475a"},
		{"set-option", "-t", name, "pane-active-border-style", "fg=#a6e3a1"},
		{"set-option", "-t", name, "window-style", "bg=#1e1e2e"},
		{"set-option", "-t", name, "window-active-style", "bg=#1e1e2e"},
	}

	for _, args := range options {
		exec.Command("tmux", args...).Run()
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
	return err
}
