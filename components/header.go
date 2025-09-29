package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// UnderlinedTitle creates a title with an underline
func UnderlinedTitle(text string) string {
	titleStyled := TitleStyle.Padding(0, 1).Render(text)

	underline := lipgloss.NewStyle().
		Foreground(MutedColor).
		Render(strings.Repeat("─", len(text)+2))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyled,
		underline,
	)
}

// PageHeader creates a styled page header with optional subtitle
func PageHeader(title string, subtitle ...string) string {
	header := TitleStyle.Render(title)

	if len(subtitle) > 0 && subtitle[0] != "" {
		header += " " + subtitle[0]
	}

	return header
}