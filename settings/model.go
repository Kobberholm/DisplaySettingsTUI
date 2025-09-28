package settings

import (
	"DisplaySettingsTUI/display"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	mainModel tea.Model
	display   *display.Display
}

func NewModel(mainModel tea.Model, display *display.Display) *Model {
	return &Model{
		mainModel: mainModel,
		display:   display,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m.mainModel, tea.Quit
		case "esc":
			return m.mainModel, m.mainModel.Init()
		}
	}

	return m, nil
}

func (m *Model) View() string {
	return "Settings are not ready yet"
}
