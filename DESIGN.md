# System Design

## Components & Roles

```
┌─────────────────────────────────────────────────────────────┐
│                        main.go                              │
│                    (Entry Point)                            │
│         - Parse CLI args                                    │
│         - Initialize TUI                                    │
│         - Start Bubble Tea program                          │
└─────────────────────┬───────────────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────────────┐
│                     internal/tui                            │
│                   (UI Layer)                                │
│                                                             │
│  model.go  - App state (sessions, cursor, view mode)        │
│  view.go   - Render UI to string                            │
│  keys.go   - Handle keyboard input                          │
│  styles.go - Colors, borders, layout                        │
└──────────┬─────────────────────────────────────────┬────────┘
           │                                         │
           │ "get sessions"                          │ "create/delete"
           │ "get status"                            │ "attach"
           ▼                                         ▼
┌─────────────────────────┐    ┌─────────────────────────────┐
│   internal/session      │◄───│      internal/tmux          │
│   (Data Layer)          │    │   (System Layer)            │
│                         │    │                             │
│ - Session struct        │    │ - List()    → tmux ls       │
│ - Status enum           │    │ - Create()  → tmux new      │
│ - Status detection      │    │ - Attach()  → tmux attach   │
│   logic                 │    │ - Kill()    → tmux kill     │
│                         │    │ - Capture() → tmux capture  │
└─────────────────────────┘    └─────────────────────────────┘
                                           │
                                           ▼
                                    ┌─────────────┐
                                    │    tmux     │
                                    │  (system)   │
                                    └─────────────┘
```

---

## Data Flow

### 1. Startup

```
main.go → tui.New() → tmux.List() → session.Parse() → render UI
```

### 2. User presses `n` (new session)

```
keys.go (capture 'n') 
    → model.go (switch to "new dialog" view)
    → view.go (render dialog)
    → user fills form, presses Enter
    → tmux.Create(name, tool, path)
    → tmux.List() (refresh)
    → view.go (render updated list)
```

### 3. User presses `Enter` (attach)

```
keys.go (capture 'Enter')
    → model.go (get selected session)
    → tmux.Attach(sessionName)
    → (hive suspends, user is in tmux session)
    → (user detaches with Ctrl+b d)
    → (hive resumes, refreshes list)
```

### 4. Status polling (background)

```
Every 2 seconds:
    → tmux.Capture(session) (get terminal output)
    → session.DetectStatus(output) 
    → model.go (update session status)
    → view.go (re-render)
```

---

## Key Interactions

| From | To | What |
|------|----|------|
| `main.go` | `tui/` | Starts the TUI program |
| `tui/model.go` | `tmux/` | Requests session operations |
| `tui/model.go` | `session/` | Gets parsed session data |
| `tmux/` | `session/` | Returns raw data, session parses it |
| `tui/view.go` | `tui/styles.go` | Uses styles to render |
| `tui/keys.go` | `tui/model.go` | Updates state based on input |

---

## Separation of Concerns

| Layer | Responsibility | Knows about |
|-------|----------------|-------------|
| `main.go` | Bootstrap | tui |
| `tui/` | UI, user interaction | session, tmux |
| `session/` | Data structures, status logic | nothing |
| `tmux/` | System commands | session (returns Session structs) |

---

## Core Data Structures

### Session

```go
type Session struct {
    Name         string        // "pi-backend"
    Tool         string        // "pi", "claude", "bash"
    Path         string        // "~/projects/myapp"
    Status       Status        // Running, Waiting, Idle
    CreatedAt    time.Time     
    LastActivity time.Time  
}
```

### Status

```go
type Status int

const (
    StatusIdle Status = iota
    StatusRunning
    StatusWaiting
)
```

---

## File Structure

```
hive-v0/
├── main.go                  # Entry point
├── go.mod
├── go.sum
├── internal/
│   ├── tmux/
│   │   └── tmux.go          # tmux commands
│   ├── tui/
│   │   ├── model.go         # Bubble Tea model
│   │   ├── view.go          # Render UI
│   │   ├── keys.go          # Key bindings
│   │   └── styles.go        # Lipgloss styles
│   └── session/
│       └── session.go       # Session struct, status detection
├── HIVE.md                  # Project overview
└── DESIGN.md                # This file
```
