package components

// HelpText returns formatted help text
func HelpText(text string) string {
	return HelpTextStyle.Render(text)
}

// Navigation help texts
const (
	RootHelpText     = "←/→ Navigate • Enter Select • Q Quit"
	SettingsHelpText = "↑/↓ Navigate • ←/→ Adjust • ESC Back • Q Quit"
)