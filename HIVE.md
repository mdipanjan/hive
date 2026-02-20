# 🐝 HIVE

A lightweight TUI for managing tmux sessions, optimized for AI coding agents.

---

## WHAT

Core features:

| Feature | Description |
|---------|-------------|
| List sessions | Show all tmux sessions with status |
| Create session | Launch pi/claude/gemini/bash in new tmux session |
| Attach session | Jump into a session |
| Delete session | Kill a session |
| Status detection | Running / Waiting / Idle indicators |
| Keyboard-driven | Fast navigation, no mouse needed |

---

## WHY

| Problem | Solution |
|---------|----------|
| tmux commands are verbose | Simple keyboard shortcuts |
| Hard to see what's running | Visual status indicators |
| Switching sessions is slow | One keypress to attach |
| No overview of all agents | Dashboard view |
| agent-deck is 90k lines | Minimal, focused alternative |
| Learn Go | Perfect small project |

---

## HOW

### Tech Stack

- **Go** - Single binary, fast
- **Bubble Tea** - TUI framework
- **Lipgloss** - Styling
- **tmux** - Session backend

### Project Structure

```
hive/
├── main.go              # Entry point, CLI args
├── internal/
│   ├── tmux/
│   │   └── tmux.go      # tmux commands (list, create, attach, kill)
│   ├── tui/
│   │   ├── model.go     # Bubble Tea model, Update()
│   │   ├── view.go      # Render UI
│   │   ├── keys.go      # Key bindings
│   │   └── styles.go    # Lipgloss styles
│   └── session/
│       └── session.go   # Session struct, status detection
└── go.mod
```

### Timeline

**Day 1:**

| # | Task | Est |
|---|------|-----|
| 1 | Project setup (go mod, deps) | 15 min |
| 2 | tmux.go - List/Create/Attach/Kill | 2 hrs |
| 3 | Basic TUI - List sessions | 2 hrs |
| 4 | Keyboard navigation | 1 hr |
| 5 | Create/Delete from TUI | 1 hr |

**Day 2:**

| # | Task | Est |
|---|------|-----|
| 6 | Status detection (running/idle/waiting) | 2 hrs |
| 7 | Styling (colors, borders, icons) | 1 hr |
| 8 | New session dialog (tool, path, name) | 2 hrs |
| 9 | Polish & testing | 1 hr |

---

## UI

### Main View

```
┌─ hive ──────────────────────────────────┐
│                                         │
│  ● pi-backend          running    2m    │
│  ◐ pi-frontend         waiting    5m  ← │
│  ○ claude-api          idle      12m    │
│                                         │
╰─────────────────────────────────────────╯
  n: new   enter: attach   d: delete   q: quit
```

### New Session Dialog

```
┌─ new session ───────────────────────────┐
│                                         │
│  Tool: [pi] claude  gemini  bash        │
│  Path: ~/projects/myapp                 │
│  Name: my-session                       │
│                                         │
│           [Create]  [Cancel]            │
╰─────────────────────────────────────────╯
```

---

## Key Bindings

| Key | Action |
|-----|--------|
| `↑/k` | Move up |
| `↓/j` | Move down |
| `Enter` | Attach to session |
| `n` | New session |
| `d` | Delete session |
| `r` | Rename session |
| `q` | Quit hive |
| `?` | Help |

---

## Status Icons

| Icon | Status | Color | Detection |
|------|--------|-------|-----------|
| `●` | Running | Green | Process active, output changing |
| `◐` | Waiting | Yellow | Prompt visible, awaiting input |
| `○` | Idle | Gray | No recent activity |

---

## Supported Tools

| Tool | Command |
|------|---------|
| pi | `pi` |
| claude | `claude` |
| gemini | `gemini` |
| opencode | `opencode` |
| bash | `bash` |

---

## Installation

```bash
# Build from source
go build -o hive .

# Or install
go install github.com/dipanjanmondal/hive@latest
```

## Usage

```bash
# Launch TUI
hive

# That's it!
```
