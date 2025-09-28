package settings

import (
	"DisplaySettingsTUI/display"
	"context"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type tickMsg time.Time
type loadInitialValues string
type valueModified float64

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
	cancelFunc     context.CancelFunc
	ctx            context.Context
	chanData       chan data
}

type data struct {
	s setting
}

func NewModel(mainModel tea.Model, display *display.Display) *Model {

	ctx, cancel := context.WithCancel(context.Background())

	return &Model{
		mainModel:  mainModel,
		display:    display,
		ctx:        ctx,
		cancelFunc: cancel,
		chanData:   make(chan data, 100),
	}
}

func (m *Model) Init() tea.Cmd {
	m.currentSetting = brightness
	m.models = make([]progress.Model, maxSetting)
	for i, _ := range m.models {
		m.models[i] = progress.New(progress.WithDefaultGradient())
	}

	m.eventLoop()

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
			m.cancelFunc()
			return m.mainModel, tea.Quit
		case "esc":
			m.cancelFunc()
			return m.mainModel, m.mainModel.Init()
		case "up":
			m.previousSetting()
		case "down":
			m.nextSetting()
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
	case valueModified:
		select {
		case m.chanData <- data{s: m.getCurrentSetting()}:
		default:
			log.Error("Channel is full")
		}
	}

	return m, nil
}

func (m *Model) View() string {
	if !m.initialized {
		return "Loading..."
	}

	// Build rows
	rows := []string{
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("99")).
			Bold(true).
			Render("Display Settings ", m.display.Title()),
		"",
		m.buildRow(brightness, "Brightness", m.currentSetting == brightness),
		"",
		m.buildRow(contrast, "Contrast", m.currentSetting == contrast),
		"",
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Render("↑/↓ Navigate • ←/→ Adjust • ESC Back • Q Quit"),
	}

	// Add padding and join
	pad := strings.Repeat(" ", padding)
	for i, row := range rows {
		if row != "" {
			rows[i] = pad + row
		}
	}

	return "\n" + strings.Join(rows, "\n")
}

func (m *Model) buildRow(s setting, label string, selected bool) string {
	indicator := "  "
	if selected {
		indicator = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86")).
			Bold(true).
			Render("▶ ")
	}

	labelStyled := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Width(12).
		Align(lipgloss.Left).
		Render(label + ":")

	return indicator + labelStyled + m.models[s].View()
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

func valueModifiedCmd() tea.Cmd {
	return func() tea.Msg {
		return valueModified(0)
	}
}

func (m *Model) eventLoop() {
	go func() {
		defer close(m.chanData)

		var lastData *data
		timer := time.NewTimer(time.Hour) // Create inactive timer
		timer.Stop()

		for {
			select {
			case <-m.ctx.Done():
				return

			case d := <-m.chanData:
				lastData = &d
				timer.Reset(100 * time.Millisecond) // Reset debounce timer

			case <-timer.C:
				if lastData != nil {
					var code VCPCode
					switch lastData.s {
					case brightness:
						code = VCPBrightness
					case contrast:
						code = VCPContrast
					default:
						panic("unhandled default case")
					}
					value := m.getPercent(lastData.s) * 100
					err := setVCP(m.display.Index, code, int(value))
					if err != nil {
						log.Error("Error setting VCP value", err)
					}
					lastData = nil
				}
			}
		}
	}()
}
