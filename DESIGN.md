# Hive — Design

> Design language for **Hive**, a keyboard-first terminal session manager.
> Last updated: 2026-06-30 · Status: living document

---

## 1. Philosophy

Hive is a full-screen TUI that you live inside for seconds at a time — open it, find a session, attach, gone. Every design decision serves that loop. Three principles drive the whole interface:

1. **Minimal, not bare.** One accent color, one structural border weight, generous negative space. Chrome earns its place or it's cut. "Too cluttered" is the failure mode we design against.
2. **Keyboard is the only input.** Nothing requires a mouse. Every screen advertises its keys in a footer hint bar, and bindings are consistent across screens (`enter` always confirms, `esc` always backs out).
3. **Legible at a glance.** Status is encoded in **both shape and color** so it survives colorblindness, theme switches, and low-contrast terminals. Text truncates rather than wraps. The thing you most likely want is pre-selected.

Hive should feel like a native part of the terminal — same fonts, same darkness — not a web app cosplaying as one.

---

## 2. Foundations

### 2.1 Color — Tokyo Night

The default (and reference) theme is **Tokyo Night (Night)**. Themes are swappable at runtime via `t`; every theme must supply the full token set below. Colors are referenced by **role token**, never by raw hex, so a new theme is a values-only change.

| Token          | Hex       | Role                                                                      |
| -------------- | --------- | ------------------------------------------------------------------------- |
| `bg`           | `#16161e` | Screen background (deepened from canonical `#1a1b26` for higher contrast) |
| `bg.titlebar`  | `#0f0f16` | Faux window chrome                                                        |
| `fg`           | `#c0caf5` | Primary text                                                              |
| `muted`        | `#565f89` | Labels, hints, secondary text                                             |
| `border`       | `#2a2e3f` | Panel / divider hairlines                                                 |
| `border.focus` | `#3b4261` | Modal borders, emphasis                                                   |
| `selection`    | `#283457` | Selected-row / focused-control fill                                       |
| `accent`       | `#7aa2f7` | Blue — headers, prompts, keys, primary focus                              |
| `cyan`         | `#7dcfff` | Tool names, secondary info                                                |
| `green`        | `#9ece6a` | Attached / done / success                                                 |
| `yellow`       | `#e0af68` | Detached / warning                                                        |
| `red`          | `#f7768e` | Dead / destructive                                                        |
| `magenta`      | `#bb9af7` | Reserved (unused today)                                                   |

**Color budget:** at most one accent (`accent`) plus the status ramp per screen. If a screen needs more than two non-status colors, the design is wrong.

### 2.2 Typography

- **One family.** A monospace stack — `JetBrains Mono`, falling back to the terminal's font. The interface assumes a fixed grid; proportional fonts break alignment.
- **No icon fonts.** We do **not** depend on Nerd Fonts or PUA glyphs. The full icon vocabulary is restricted to the BMP set in §2.3, which renders in any reasonable monospace font. (This is a hard rule — the original build shipped tofu boxes where Nerd glyphs were missing.)
- **Casing:**
  - Screen / panel titles → `UPPERCASE` with `0.18em` letter-spacing (`SESSIONS`, `DELETE SESSION`).
  - Body and labels → sentence case (`Delete "mail-work"?`).
  - Footer hint keys → lowercase (`enter: confirm`).

### 2.3 Iconography — the status system

This is the load-bearing convention. **One state vocabulary, identical across every screen.** A session that is `attached` is a green `●` on the dashboard, in the switcher, and in details — never a different color in a different place.

| State      | Glyph | Color    | Meaning                         |
| ---------- | :---: | -------- | ------------------------------- |
| `attached` |  `●`  | `green`  | Currently attached / foreground |
| `detached` |  `◼`  | `yellow` | Alive and running, not attached |
| `idle`     |  `◌`  | `blue`   | Exists, no recent activity      |
| `done`     |  `✓`  | `green`  | Exited cleanly / task complete  |
| `dead`     |  `✗`  | `red`    | Exited with error               |

Shape disambiguates the two greens (`●` vs `✓`). Every place a status appears, it carries its **label** too — the legend is never a bare glyph row.

Safe ancillary glyphs (also BMP-guaranteed): prompt `❯`, selection caret `>`, arrows `↑ ↓ ← →`, sparkline blocks `▁▂▃▄▅▆▇█`, radio `○ ●`, cursor `▏`.

### 2.4 Spacing & shape

- Panels and modals: `1px` border (`border`), `8–10px` radius.
- Modal title sits flush top-left; actions bottom; a contextual hint bar always anchors the screen footer.
- Selected list rows get a `selection` fill **and** a `2px` left accent bar (color from the row's status). Redundant on purpose.

---

## 3. Components

| Component           | Spec                                                                                                                                                     |
| ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Panel**           | Bordered box, blue uppercase header, content inset. Used for `SESSIONS` and `DETAILS`.                                                                   |
| **List row**        | `<glyph> <name>` left, right-aligned state label in `muted`. Selected row → `selection` fill + left accent bar.                                          |
| **Modal**           | Centered, `border.focus` outline, ~340–420px. Title → body → actions → footer hints. Dims (does not hide) the screen behind it.                          |
| **Button**          | Bordered, sentence case. Default/focused button takes a `selection` fill. **Destructive** focused button uses `red` fill instead of accent (see Delete). |
| **Form field**      | `muted` fixed-width label + value over a `border` underline. Active field shows a block cursor `▏`; selected text gets a `selection` highlight.          |
| **Radio group**     | `○` unselected / `●` selected, selected glyph tinted to the relevant role.                                                                               |
| **Sparkline**       | `▁…█` blocks, preceded by a `muted` label, followed by a value colored by threshold (green < 50%, yellow < 80%, red ≥ 80%).                              |
| **Footer hint bar** | `key` in `accent` + `label` in `muted`, space-separated. Content is **per-screen and context-specific** — it shows only the keys that screen accepts.    |

---

## 4. Screens

Each screen lists its purpose, layout, the **states / edge cases** it must handle (the if-then matrix), and its keys.

### 4.1 Dashboard (home)

The landing screen. Wordmark + at-a-glance counts on the left; live session list and detail panel on the right; resource sparklines and the global hint bar along the bottom.

**Layout:** two-column grid (~0.82 / 1). Left: `HIVE` wordmark (filled blue with a `border.focus` outline echo for depth), version line, labeled status legend, total count. Right: `SESSIONS` panel (top) over `DETAILS` panel (bottom). Footer: `cpu` + `mem` sparklines, then global hints.

**States:**

- **No sessions** → list area shows `No sessions yet — press n to create one`; details panel hidden; counts read `0`.
- **Long session name** → truncate with `…`, never wrap.
- **More sessions than fit** → list scrolls; the legend's total (`N sessions`) reflects the full count.
- **Long path in details** → middle-truncate: `~/…/projects/hive`. Never hard-wrap mid-segment.
- **High load** → sparkline value flips yellow → red past thresholds.

**Keys:** `n` new · `enter` attach · `d` delete · `/` search · `t` theme · `?` help · `q` quit.

### 4.2 Search

A fuzzy filter over session names, invoked with `/`.

**Layout:** modal. `SEARCH` title, single prompt line `❯ <query>`, filtered results below, match counter `m / n` in `muted` at the bottom.

> **Note:** one prompt marker only (`❯`). The original drew `❯ >` — a doubled prompt. Cut the inner `>`.

**States:**

- **Matches** → list narrows live; first result auto-highlighted; counter updates.
- **No matches** → `no matches` line, counter `0 / n`, `enter` is a no-op.
- **Empty query** → show all sessions (counter `n / n`).

**Keys:** `↑↓` select · `enter` attach · `esc` close.

### 4.3 Switch session

Like Search but its `enter` switches the active session rather than attaching a new view. Reachable directly so you can jump without typing.

**Layout:** modal. `SWITCH SESSION` title, `❯ Search…` prompt (single marker), full session list with status glyphs, `n / n` counter.

**States:** mirror Search. The **currently active** session is marked and may be pre-selected; switching to it is a no-op.

**Keys:** `↑↓` select · `enter` switch · `esc` close.

### 4.4 New session

Create a session: pick a tool, a working directory, and a name.

**Layout:** modal. `NEW SESSION` title. `TOOL` radio row (`pi` · `claude` · `nvim` · `bash`, default `bash`). `PATH` field (opens the file picker, §4.5). `NAME` field with cursor. `Create` / `Cancel` buttons, `Create` focused by default.

**States:**

- **Empty name** → `Create` is inert; show inline `name required` rather than failing on confirm.
- **Name collision** → inline `name in use` warning; block create.
- **Invalid / missing path** → inline error; block create.
- **Tool default** → `bash` pre-selected so `enter` works with zero radio interaction.

**Keys:** `tab` next field · `←→` select (radio / buttons) · `enter` confirm · `esc` cancel.

### 4.5 File picker (select directory)

Directory chooser for the new-session path field.

**Layout:** tall modal. `SELECT DIRECTORY` title, scrollable entry list. Current selection prefixed `>` and tinted `green`. **Directories** render in `accent`/`fg`; **files** in `muted` — so you can tell them apart at a glance (e.g. `config`, `quint-lsp.log` read as non-directories).

**States:**

- **Empty directory** → `(empty)` placeholder.
- **Permission denied** → `(cannot read)` line; `enter` is a no-op on it.
- **Long listing** → scrolls; selection stays in view.
- **Hidden entries** → hidden by default (consider a toggle key — see Open questions).

**Keys:** `↑↓` navigate · `enter` select · `esc` cancel.

### 4.6 Delete session (confirm)

A destructive confirmation, invoked with `d`.

**Layout:** modal. `DELETE SESSION` title, body `Delete "<name>"?`, `Yes` / `No` buttons. **`No` is focused by default** — the safe choice never sits under the trigger finger.

**States & rules:**

- **Default focus = No.** Confirming destruction must be deliberate.
- **Destructive emphasis:** when `Yes` is focused, it takes the `red` fill (not the neutral/accent fill the original used). Color reinforces the stakes. `No` keeps the accent fill.
- **Attached session** → append a warning to the body: `Delete "<name>"? It's currently attached and will be killed.` so the stakes are explicit.
- `y` confirms and `n`/`esc` cancel as one-key shortcuts, in addition to `←→` + `enter`.

**Keys:** `←→` select · `enter` confirm · `y` yes · `n`/`esc` cancel.

### 4.7 Help

A keybinding reference overlay, toggled with `?`.

**Layout:** modal. `HELP` title, a two-column key → description list, and a footer line. Bindings are **grouped** for scannability rather than listed flat:

```text
HELP

  navigation
    ↑/k          Move up
    ↓/j          Move down

  actions
    n            Create new session
    enter        Attach to session
    d            Delete session
    t            Switch theme

  global
    ?            Toggle this help
    q            Quit

  hive v0.1 · press ? or esc to close
```

> The footer's `hive dev` in the original reads as a stray fragment — make it an explicit version/branding line (`hive v0.1`).

**States:** opens over any non-modal screen; `?` and `esc` both close it; it does not stack over other modals.

**Keys:** `?` / `esc` close.

---

## 5. Interaction model

- **Vim + arrows.** `j/k` and `↑/↓` both move. Power users never leave the home row; everyone else uses arrows.
- **Universal verbs.** `enter` = confirm/commit, `esc` = cancel/back, single letters = actions. These never change meaning between screens.
- **Modals are a stack of one.** At most one modal is open. Help is the only overlay allowed over a base screen; it won't open atop another modal.
- **Safe defaults.** Destructive flows default to the non-destructive choice; creation flows default to a runnable state (`bash`, `Create` focused).

---

## 6. Accessibility & robustness

- **Never color alone.** Every status pairs a distinct glyph with its color and carries a text label. The UI is fully usable in monochrome.
- **Truncate, never wrap.** Names middle/end-truncate with `…`; the grid stays intact at any width.
- **Always-visible focus.** The focused control has a fill, not just a border tint.
- **Font-safe glyphs only.** No PUA/Nerd-Font dependency (§2.3) — the interface can't render boxes-of-tofu on a fresh machine.
- **Themes are total.** A theme that omits a role token is rejected at load, not rendered with fallbacks.

---

## 7. Changelog of design decisions

Captured so the rationale isn't lost:

- **Unified the status system.** Previously `hive-dev` was blue on the dashboard but green in the switcher; the legend was an unlabeled five-glyph row. Now: one labeled vocabulary, identical everywhere (§2.3).
- **Killed the tofu.** Replaced Nerd-Font/PUA icons with a BMP-only glyph set.
- **Single search prompt.** Removed the doubled `❯ >`.
- **Path truncation.** Details path middle-truncates instead of wrapping mid-word.
- **Sparklines that read.** `cpu`/`mem` sit side by side, data fills the line, value colored by load — replacing a flat line with bars jammed at the edge.
- **Destructive emphasis.** Focused `Yes` on Delete is now red; `No` stays the default focus.

---

## 8. Open questions

- Should the dashboard's lower-left negative space hold something useful (recent-activity log? inline keymap?) or stay empty for calm?
- `t: theme` — cycle inline through themes, or open a picker modal?
- File picker: add a `.` key to toggle hidden entries?
- Help: is the grouped layout worth the extra vertical space versus the flat list?
