package components

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	loadingStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true).
		Padding(2, 4)

	loadingTextStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Italic(true)
)

// LoadingText returns a formatted loading message
func LoadingText(message string) string {
	icon := loadingStyle.Render("⟳")
	text := loadingTextStyle.Render(message)
	return lipgloss.JoinHorizontal(lipgloss.Center, icon, " ", text)
}

// LoadingBox returns a loading message in a box
func LoadingBox(message string) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 3)

	return boxStyle.Render(LoadingText(message))
}

// CenteredLoading returns a centered loading message
func CenteredLoading(message string, width int) string {
	content := LoadingBox(message)
	return lipgloss.PlaceHorizontal(width, lipgloss.Center, content)
}