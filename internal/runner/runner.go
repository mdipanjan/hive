package runner

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/mdipanjan/hive/internal/state"
)

type Options struct {
	Name string
	Tool string
	Path string
}

func Run(opts Options) error {
	if opts.Name == "" {
		return fmt.Errorf("session name is required")
	}
	if opts.Tool == "" {
		return fmt.Errorf("tool is required")
	}
	if opts.Path == "" {
		return fmt.Errorf("path is required")
	}

	if err := os.Chdir(opts.Path); err != nil {
		markFailed(opts.Name, time.Now(), 1)
		return err
	}

	if opts.Tool == "nvim" || opts.Tool == "vim" {
		return runEditor(opts)
	}

	return runOnce(opts)
}

func runOnce(opts Options) error {
	startedAt := time.Now().UTC()
	if err := state.WriteRuntime(opts.Name, state.Runtime{State: state.StateRunning, StartedAt: startedAt}); err != nil {
		return err
	}

	exitCode, _ := runCommand(opts.Tool, nil)
	endedAt := time.Now().UTC()

	runtimeState := state.StateCompleted
	if exitCode != 0 {
		runtimeState = state.StateFailed
	}

	writeErr := state.WriteRuntime(opts.Name, state.Runtime{
		State:     runtimeState,
		StartedAt: startedAt,
		EndedAt:   &endedAt,
		ExitCode:  &exitCode,
	})
	if writeErr != nil {
		return writeErr
	}

	return runShell()
}

func runEditor(opts Options) error {
	for {
		startedAt := time.Now().UTC()
		if err := state.WriteRuntime(opts.Name, state.Runtime{State: state.StateRunning, StartedAt: startedAt}); err != nil {
			return err
		}

		exitCode, err := runCommand("nvim", []string{"."})
		endedAt := time.Now().UTC()

		runtimeState := state.StateReady
		if exitCode != 0 {
			runtimeState = state.StateFailed
		}

		if writeErr := state.WriteRuntime(opts.Name, state.Runtime{
			State:     runtimeState,
			StartedAt: startedAt,
			EndedAt:   &endedAt,
			ExitCode:  &exitCode,
		}); writeErr != nil {
			return writeErr
		}

		detachClient(opts.Name)

		if err != nil {
			return err
		}
	}
}

func runCommand(name string, args []string) (int, error) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err == nil {
		return 0, nil
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr.ExitCode(), err
	}

	return 1, err
}

func runShell() error {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	_, err := runCommand(shell, nil)
	return err
}

func detachClient(name string) {
	_ = exec.Command("tmux", "detach-client", "-s", name).Run()
}

func markFailed(name string, startedAt time.Time, exitCode int) {
	endedAt := time.Now().UTC()
	_ = state.WriteRuntime(name, state.Runtime{
		State:     state.StateFailed,
		StartedAt: startedAt.UTC(),
		EndedAt:   &endedAt,
		ExitCode:  &exitCode,
	})
}
