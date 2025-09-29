# DisplaySettingsTUI Demo
# VHS tape file for generating demo GIF

# Terminal setup
Output demo.gif
Set Shell zsh
Set FontSize 46
Set Width 1920
Set Height 1080
Set Theme "Dracula"

# Start recording
Sleep 1s
Type "display-settings-tui"
Sleep 500ms
Enter

# Wait for app to load and show displays
Sleep 2s

# Navigate between displays
Right
Sleep 800ms
Left
Sleep 800ms
Right
Sleep 800ms

# Select first display
Left
Sleep 500ms
Enter

# Wait for settings to load
Sleep 2s

# Adjust brightness - lots of swinging!
Left 8
Sleep 300ms
Right 12
Sleep 300ms
Left 5
Sleep 300ms
Right 10
Sleep 300ms
Left 3
Sleep 500ms

# Switch to contrast
Down
Sleep 800ms

# Adjust contrast - more swinging
Right 8
Sleep 300ms
Left 10
Sleep 300ms
Right 6
Sleep 300ms
Left 4
Sleep 300ms
Right 7
Sleep 500ms

# Go back to brightness
Up
Sleep 500ms

# Final adjustments with swinging
Left 6
Sleep 300ms
Right 9
Sleep 300ms
Left 2
Sleep 1s

# Return to display selection
Escape
Sleep 1.5s

# Switch to second display
Right
Sleep 800ms
Enter

# Quick adjustments on second display with swinging
Sleep 2s
Right 10
Sleep 300ms
Left 15
Sleep 300ms
Right 8
Sleep 300ms
Left 4
Sleep 300ms
Right 11
Sleep 300ms
Left 5
Sleep 1s

# Return to main menu
Escape
Sleep 1.5s

# Exit application
Type "q"
Sleep 500ms