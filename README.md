# Hive

A lightweight TUI for managing tmux sessions, optimized for AI coding agents.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS-lightgrey)
![License](https://img.shields.io/badge/License-MIT-green)

<p align="center">
  <img src="docs/demo.gif" alt="Hive Demo" width="700">
</p>

## Why Hive?

Managing multiple AI coding agents (Claude, Pi, Cursor) across projects gets messy fast. Hive solves this.

**Problem:** Juggling tmux sessions manually - creating, naming, finding, switching.

**Solution:** A single interface to manage all your agent sessions.

## Features

| Feature | What it solves |
|---------|----------------|
| **Unified Dashboard** | See all sessions at a glance - status, tool, path |
| **Quick Search** | Find sessions instantly with `/` - no more `tmux ls` |
| **CLI + JSON Output** | AI agents can create/manage sessions programmatically |
| **12 Themes** | Match your terminal aesthetic |
| **Lightweight** | ~4MB binary, instant startup, no runtime deps |

## Requirements

- **tmux** - Terminal multiplexer
- **Go 1.21+** - For building from source

```bash
# Check if tmux is installed
tmux -V
```

## Installation

### From Source

```bash
git clone https://github.com/mdipanjan/hive-v0.git
cd hive-v0
make build
make install  # Installs to /usr/local/bin
```

### Manual

```bash
go build -o hive .
mv hive /usr/local/bin/
```

## Usage

### TUI Mode

```bash
hive
```

**Keyboard Shortcuts:**

| Key | Action |
|-----|--------|
| `n` | New session |
| `enter` | Attach to session |
| `d` | Delete session |
| `/` | Search sessions |
| `t` | Cycle themes |
| `?` | Help |
| `q` | Quit |

### CLI Mode (AI-Agent Compatible)

```bash
# List sessions
hive list              # Plain text
hive list --json       # JSON output

# Create session
hive create --tool pi --path /projects/myapp --name my-session

# Attach to session
hive attach my-session

# Delete session
hive delete my-session
```

**Aliases:**
- `list` / `ls`
- `create` / `new`
- `attach` / `a`
- `delete` / `rm`

**JSON Output Example:**

```json
[
  {
    "name": "hive-abc123",
    "tool": "pi",
    "path": "/projects/myapp",
    "status": "running"
  }
]
```

## Configuration

Config file: `~/.config/hive/config.json`

```json
{
  "theme": "tokyo-night"
}
```

## Themes

Press `t` to cycle through themes:

- Tokyo Night (default)
- Tokyo Storm
- Dracula
- Nord
- Gruvbox
- Catppuccin
- One Dark
- Solarized Dark
- GitHub Dark
- Rose Pine
- Monokai
- Zinc Dark

## Debug Mode

```bash
DEBUG=1 hive
```

Logs are written to `debug.log`.

## Project Structure

```
hive-v0/
  main.go
  internal/
    cli/          # CLI commands
    components/   # UI components
    config/       # Configuration
    logger/       # Logging
    session/      # Session types
    styles/       # Themes & styles
    tmux/         # tmux operations
    tui/          # TUI logic
  docs/
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.
