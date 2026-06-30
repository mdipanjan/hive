# Hive TUI Screens

Reference snapshots of every screen in the Hive TUI (theme: Tokyo Night).

The `.txt` files are raw ANSI snapshots rendered straight from `RenderView`, so
they stay in sync with the code and diff cleanly in review. PNGs are generated
artifacts (gitignored) — render them locally when you need images.

| Screen                                          | Snapshot               |
| ----------------------------------------------- | ---------------------- |
| Dashboard (session list + details + CPU)        | `01-dashboard.txt`     |
| Search overlay                                  | `02-search.txt`        |
| Session switcher (`prefix + h` / `hive switch`) | `03-switch.txt`        |
| New session form                                | `04-new-form.txt`      |
| Directory picker                                | `05-filepicker.txt`    |
| Delete confirmation                             | `06-delete-confirm.txt`|
| Help overlay                                    | `07-help.txt`          |

## Regenerating

```bash
# 1. Render each screen to an ANSI .txt snapshot
HIVE_SCREENSHOTS=1 go test ./internal/tui -run TestGenerateScreenshots

# 2. (optional) Convert snapshots to PNGs with charmbracelet/freeze
cd screenshots
for f in *.txt; do freeze --language ansi "$f" -o "${f%.txt}.png"; done
```

Install freeze with `brew install charmbracelet/tap/freeze` or
`go install github.com/charmbracelet/freeze@latest`.

> Status icons use Nerd Font glyphs; install a Nerd Font and pass
> `--font.family "<name>"` to freeze for pixel-perfect icons.
