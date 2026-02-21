# Hive Demo - VHS Tape
# Run: vhs docs/demo.tape

Output docs/demo.gif

Set FontSize 16
Set Width 1200
Set Height 900
Set Theme TokyoNight

# Start hive
Type "./hive"
Sleep 500ms
Enter
Sleep 3s

# Navigate sessions
Down
Sleep 800ms
Down
Sleep 800ms
Up
Sleep 800ms

# Open new session dialog
Type "n"
Sleep 2s

# Select tool (move to claude)
Right
Sleep 600ms

# Move to path field
Tab
Sleep 600ms

# Move to name field
Tab
Sleep 600ms

# Move to buttons and cancel
Tab
Sleep 600ms
Right
Sleep 500ms
Enter
Sleep 2s

# Search
Type "/"
Sleep 1s
Type "hive"
Sleep 1.5s
Down
Sleep 500ms
Down
Sleep 500ms
Escape
Sleep 1s

# Delete confirmation
Type "d"
Sleep 2s
Left
Sleep 600ms
Right
Sleep 600ms
Escape
Sleep 1s

# Show help
Type "?"
Sleep 3s
Escape
Sleep 1s

# Cycle all 12 themes
Type "t"
Sleep 800ms
Type "t"
Sleep 800ms
Type "t"
Sleep 800ms
Type "t"
Sleep 800ms
Type "t"
Sleep 800ms
Type "t"
Sleep 800ms
Type "t"
Sleep 800ms
Type "t"
Sleep 800ms
Type "t"
Sleep 800ms
Type "t"
Sleep 800ms
Type "t"
Sleep 800ms
Type "t"
Sleep 1.5s

# Quit
Type "q"
Sleep 1s
