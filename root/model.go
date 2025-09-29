package root

import (
	"DisplaySettingsTUI/display"
	"DisplaySettingsTUI/settings"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

/* STYLING */
var (
	columnStyle = lipgloss.NewStyle().
			Padding(2, 4)
	focusedStyle = lipgloss.NewStyle().
			Padding(2, 4).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
)

type Model struct {
	index         int
	displayModels []list.Model
	currentWidth  int
}

func New() *Model {
	return &Model{}
}

func (m *Model) Init() tea.Cmd {
	return tea.SetWindowTitle("DisplaySettingsTUI")
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left":
			m.next()
		case "right":
			m.prev()
		case "enter":
			item := m.displayModels[m.index].Items()[0]
			displayItem, ok := item.(*display.Display)
			if !ok {
				// Handle the error - the item wasn't a *display.Display
				return m, nil // or handle error appropriately
			}
			newModel := settings.NewModel(m, displayItem, m.currentWidth)
			return newModel, newModel.Init()
		}

	case tea.WindowSizeMsg:
		m.currentWidth = msg.Width
		m.loadDisplays(msg.Width, msg.Height)
	}

	if len(m.displayModels) == 0 {
		return m, nil
	}

	var cmd tea.Cmd
	m.displayModels[m.index], cmd = m.displayModels[m.index].Update(msg)
	return m, cmd
}

func (m *Model) View() string {

	if len(m.displayModels) == 0 {
		return "Loading..."
	}

	displayString := make([]string, len(m.displayModels))
	for i, d := range m.displayModels {
		if i == m.index {
			displayString[i] = focusedStyle.Render(d.View())
		} else {
			displayString[i] = columnStyle.Render(d.View())
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Top, underlinedTitle("Available Displays"),
		lipgloss.JoinHorizontal(lipgloss.Left, displayString...),
	)
}

func (m *Model) prev() {
	totalLength := len(m.displayModels)
	m.index = m.index - 1
	if m.index < 0 {
		m.index = totalLength - 1
	}
}

func (m *Model) next() {
	totalLength := len(m.displayModels)
	m.index = m.index + 1
	if m.index >= totalLength {
		m.index = 0
	}
}

func underlinedTitle(text string) string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("222")).
		MarginBottom(0)

	underline := lipgloss.NewStyle().
		Foreground(lipgloss.Color("22")).
		MarginBottom(1).
		Render(strings.Repeat("─", len(text)))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(text),
		underline,
	)
}

func (m *Model) loadDisplays(width, height int) {
	displayList, err := display.DetectDisplays()
	if err != nil {
		log.Error("Failed to load displays: %v", err)
		return
	}

	if len(displayList) == 0 {
		log.Error("No displays found")
		return
	}

	m.displayModels = make([]list.Model, len(displayList))

	for i, display := range displayList {
		// Render a square box
		m.displayModels[i] = list.New([]list.Item{&display}, list.NewDefaultDelegate(), width, height/4)
		m.displayModels[i].Title = fmt.Sprintf("Display %d", display.Index)
		m.displayModels[i].SetShowHelp(false)
		m.displayModels[i].SetShowStatusBar(false)
	}

	m.index = 0
}
