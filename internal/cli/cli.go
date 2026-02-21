package cli

import (
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/mdipanjan/hive/internal/tmux"
)

type SessionOutput struct {
	Name   string `json:"name"`
	Tool   string `json:"tool"`
	Path   string `json:"path"`
	Status string `json:"status"`
}

func Run(args []string) bool {
	if len(args) < 2 {
		return false
	}

	switch args[1] {
	case "list", "ls":
		return runList(args[2:])
	case "create", "new":
		return runCreate(args[2:])
	case "attach", "a":
		return runAttach(args[2:])
	case "delete", "rm":
		return runDelete(args[2:])
	}

	return false
}

func runList(args []string) bool {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	jsonOutput := fs.Bool("json", false, "Output as JSON")
	fs.Parse(args)

	sessions, err := tmux.List()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if *jsonOutput {
		output := make([]SessionOutput, len(sessions))
		for i, s := range sessions {
			output[i] = SessionOutput{
				Name:   s.Name,
				Tool:   s.Tool,
				Path:   s.Path,
				Status: s.Status.String(),
			}
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(output)
	} else {
		if len(sessions) == 0 {
			fmt.Println("No sessions")
			return true
		}
		for _, s := range sessions {
			fmt.Printf("%s\t%s\t%s\t%s\n", s.Name, s.Status.String(), s.Tool, s.Path)
		}
	}

	return true
}

func runCreate(args []string) bool {
	fs := flag.NewFlagSet("create", flag.ExitOnError)
	tool := fs.String("tool", "bash", "Tool to use (pi, claude, bash)")
	path := fs.String("path", ".", "Working directory")
	name := fs.String("name", "", "Session name (auto-generated if empty)")
	fs.Parse(args)

	if *name == "" {
		*name = "hive-" + randomID(6)
	}

	if *path == "." {
		*path, _ = os.Getwd()
	}

	err := tmux.Create(*name, *tool, *path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created session: %s\n", *name)
	return true
}

func runAttach(args []string) bool {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: hive attach <session-name>\n")
		os.Exit(1)
	}

	name := args[0]
	cmd := tmux.AttachCmd(name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	return true
}

func runDelete(args []string) bool {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: hive delete <session-name>\n")
		os.Exit(1)
	}

	name := args[0]
	err := tmux.Kill(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Deleted session: %s\n", name)
	return true
}

func randomID(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	rand.Read(b)
	for i := range b {
		b[i] = chars[b[i]%byte(len(chars))]
	}
	return string(b)
}
