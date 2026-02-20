# AESTHETICS.md - Hive Visual Design

## Version History

| Version | Date | Description |
|---------|------|-------------|
| v0.1 | 2026-02-20 17:30 IST | Initial basic UI (list + help bar) |
| v0.2 | 2026-02-20 21:00 IST | Dashboard layout design planned |

---

## v0.2 - Dashboard Layout (Planned)

### Target Design

```
╭─────────────────────────────────────────────────────────────────────────────────────╮
│                                                                                     │
│   ██╗  ██╗██╗██╗   ██╗███████╗                                                      │
│   ██║  ██║██║██║   ██║██╔════╝         ╭─ SESSIONS ─────────────────────────────╮   │
│   ███████║██║██║   ██║█████╗           │                                        │   │
│   ██╔══██║██║╚██╗ ██╔╝██╔══╝           │  ● pi-backend          running    2m   │   │
│   ██║  ██║██║ ╚████╔╝ ███████╗         │  ◐ claude-api          waiting    5m ← │   │
│   ╚═╝  ╚═╝╚═╝  ╚═══╝  ╚══════╝         │  ○ gemini-test         idle      12m   │   │
│                                        │  ○ opencode-pr         idle      45m   │   │
│      ⬡ ⬡ ⬡                             │                                        │   │
│     ⬡ ⬡ ⬡ ⬡                            ╰────────────────────────────────────────╯   │
│      ⬡ ⬡ ⬡                                                                          │
│                                        ╭─ SELECTED ─────────────────────────────╮   │
│   agents: 4    active: 2               │  NAME     claude-api                   │   │
│                                        │  TOOL     claude                       │   │
│                                        │  PATH     ~/projects/backend           │   │
│                                        │  STATUS   ◐ waiting for input          │   │
│                                        │  UPTIME   5m 32s                       │   │
│                                        ╰────────────────────────────────────────╯   │
│                                                                                     │
│   ╭─ ACTIVITY ──────────────────────────────────────────────────────────────────╮   │
│   │  ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁▁▂▃▄▅▆▇█▇▆▅▄▃▂▁▁▂▃▄▅▆▇█▇▆▅▄▃▂▁▁▂▃▄▅▃▂▁▁▂▃▄▅▆▇█▇▆▅▄▃▂▁     │   │
│   │  cpu                                                              12%       │   │
│   ╰─────────────────────────────────────────────────────────────────────────────╯   │
│                                                                                     │
╰─────────────────────────────────────────────────────────────────────────────────────╯
  n new   enter attach   d delete   r rename   ? help                          q quit
```

### Color Scheme (Tokyo Night)

| Element | Color | Hex |
|---------|-------|-----|
| Background | Dark blue | `#1a1b26` |
| Borders | Gray | `#565f89` |
| Logo/Titles | Cyan | `#7aa2f7` |
| Running status | Green | `#9ece6a` |
| Waiting status | Yellow | `#e0af68` |
| Idle status | Gray | `#565f89` |
| Selected row bg | Dark gray | `#414868` |
| Text | Light gray | `#c0caf5` |
| Dim text | Dim gray | `#414868` |

### Layout Structure

```
OUTER BOX
├── TOP SECTION (JoinHorizontal)
│   ├── LEFT COLUMN (width: 30%)
│   │   ├── Logo (ASCII art)
│   │   ├── Honeycomb (decorative)
│   │   └── Stats (agents count)
│   └── RIGHT COLUMN (width: 70%)
│       ├── Sessions Panel (bordered)
│       └── Details Panel (bordered)
├── ACTIVITY PANEL (full width, bordered)
└── HELP BAR (outside main box)
```

### Components

| Component | File | Description |
|-----------|------|-------------|
| Logo | `components.go` | ASCII art "HIVE" |
| Honeycomb | `components.go` | ⬡ pattern decoration |
| Stats | `components.go` | Agent counts |
| Sessions Panel | `components.go` | Scrollable session list |
| Details Panel | `components.go` | Selected session info |
| Activity Panel | `components.go` | CPU/activity graph |
| Help Bar | `components.go` | Key bindings |

### Implementation Order

1. [ ] Update `styles.go` - Add panel styles, title styles
2. [ ] Create `components.go` - Individual render functions
3. [ ] Update `view.go` - Compose layout
4. [ ] Update `model.go` - Add width/height tracking
5. [ ] Test & adjust spacing

### Dependencies

- `lipgloss` - Layout & styling (have)
- `bubbles/viewport` - Scrollable panels (may need)

---

## v0.1 - Initial UI (Current)

Basic list view with:
- Title
- Session list with status icons
- Help bar

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
