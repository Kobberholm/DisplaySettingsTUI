package settings

import (
	"DisplaySettingsTUI/display"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time
type loadInitialValues string

const (
	padding  = 2
	maxWidth = 600
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type Model struct {
	mainModel      tea.Model
	display        *display.Display
	currentSetting setting
	models         []progress.Model
	initialized    bool
}

func NewModel(mainModel tea.Model, display *display.Display) *Model {
	return &Model{
		mainModel: mainModel,
		display:   display,
	}
}

func (m *Model) Init() tea.Cmd {
	m.currentSetting = brightness
	m.models = make([]progress.Model, maxSetting)
	for i, _ := range m.models {
		m.models[i] = progress.New(progress.WithDefaultGradient())
	}

	return tea.Batch(tickCmd(), loadInitialValuesCmd())
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.models[m.currentSetting].Width = msg.Width - padding*2 - 4
		if m.models[m.currentSetting].Width > maxWidth {
			m.models[m.currentSetting].Width = maxWidth
		}
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m.mainModel, tea.Quit
		case "esc":
			return m.mainModel, m.mainModel.Init()
		case "up", "k":
			m.previousSetting()
		case "down", "j":
			m.previousSetting()
		case "left":
			return m.decrement(0.05)
		case "right":
			return m.increment(0.05)
		}
	case tickMsg:
		return m, tickCmd()
	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		return m.updateAll(msg)
	case loadInitialValues:
		return m, m.loadInitialValues()
	}

	return m, nil
}

func (m *Model) View() string {
	if !m.initialized {
		return "Loading..."
	}

	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.models[brightness].View() + "\n\n" +
		pad + m.models[contrast].View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}

func (m *Model) nextSetting() {
	m.currentSetting++
	if m.currentSetting == maxSetting {
		m.currentSetting = 0
	}
}

func (m *Model) previousSetting() {
	if m.currentSetting == 0 {
		m.currentSetting = maxSetting - 1
		return
	}

	m.currentSetting--
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func loadInitialValuesCmd() tea.Cmd {
	return func() tea.Msg {
		return loadInitialValues("")
	}
}
