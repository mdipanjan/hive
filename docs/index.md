---
layout: default
title: Hive
---

<header>
  <h1>Hive</h1>
  <p class="tagline">a lightweight TUI for managing tmux sessions.</p>
  <a href="https://github.com/mdipanjan/hive" class="gh-link">
    <svg viewBox="0 0 16 16" fill="currentColor"><path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"/></svg>
    Star on GitHub
  </a>
</header>

Managing multiple AI coding agents across projects gets messy fast. Hive gives you one interface to see all your tmux sessions at a glance — status, tool, path — without the mental overhead of `tmux ls` and manual switching.

*Built for the age of AI coding agents.*

<div class="demo">
  <img src="demo.gif" alt="Hive demo">
</div>

## Install

<div class="install">
  <div class="install-cmd">curl -fsSL https://raw.githubusercontent.com/mdipanjan/hive/main/install.sh | bash</div>
  <div class="install-alt">Or: go install github.com/mdipanjan/hive@latest</div>
</div>

Then just run `hive`.

## Keys

<div class="keys">
  <span class="key">n new</span>
  <span class="key">enter attach</span>
  <span class="key">d delete</span>
  <span class="key">/ search</span>
  <span class="key">t theme</span>
  <span class="key">? help</span>
  <span class="key">q quit</span>
</div>

## Features

- **Unified dashboard** — see all sessions at a glance
- **Quick search** — find and attach instantly with `/`
- **CLI + JSON output** — agents can manage sessions programmatically
- **12 themes** — Tokyo Night, Dracula, Nord, Gruvbox, and more
- **Lightweight** — ~4MB binary, instant startup, no runtime deps

## CLI

```
hive list --json              # AI-agent friendly
hive create --tool pi --path /projects/myapp
hive attach my-session
hive delete my-session
```

<footer>
  MIT License · Built with <a href="https://github.com/charmbracelet/bubbletea">Bubbletea</a> · Made by <a href="https://github.com/mdipanjan">@mdipanjan</a>
</footer>
