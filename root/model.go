package root

import (
	"DisplaySettingsTUI/components"
	"DisplaySettingsTUI/display"
	"DisplaySettingsTUI/settings"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
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
		return "\n" + components.CenteredLoading("Detecting displays...", m.currentWidth)
	}

	displayString := make([]string, len(m.displayModels))
	for i, d := range m.displayModels {
		if i == m.index {
			displayString[i] = components.FocusedStyle.Render(d.View())
		} else {
			displayString[i] = components.ColumnStyle.Render(d.View())
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		components.UnderlinedTitle("Available Displays"),
		lipgloss.JoinHorizontal(lipgloss.Left, displayString...),
		"\n"+components.HelpText(components.RootHelpText),
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
