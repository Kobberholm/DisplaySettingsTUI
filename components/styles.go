package components

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	PrimaryColor   = lipgloss.Color("99")  // Cyan/Teal for headers
	SecondaryColor = lipgloss.Color("62")  // Blue for focus borders
	AccentColor    = lipgloss.Color("86")  // Green for selections
	MutedColor     = lipgloss.Color("241") // Gray for help text
	TextColor      = lipgloss.Color("252") // Light gray for regular text
	ErrorColor     = lipgloss.Color("196") // Red for errors

	// Layout styles
	ColumnStyle = lipgloss.NewStyle().
			Padding(2, 4)

	FocusedStyle = lipgloss.NewStyle().
			Padding(2, 4).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(SecondaryColor)

	// Text styles
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(PrimaryColor)

	HelpTextStyle = lipgloss.NewStyle().
			Foreground(MutedColor)

	LabelStyle = lipgloss.NewStyle().
			Foreground(TextColor)

	SelectedIndicatorStyle = lipgloss.NewStyle().
				Foreground(AccentColor).
				Bold(true)

	// Common padding value
	Padding = 2
)