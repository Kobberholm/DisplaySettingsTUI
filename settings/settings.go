package settings

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type setting int

const (
	brightness setting = iota
	contrast
	maxSetting // Not a setting. Used for enumeration
)

func (m *Model) getCurrentSetting() setting {
	return m.currentSetting
}

func (m *Model) setBrightnessPercent(vcp *VCPValue) tea.Cmd {
	return m.models[brightness].SetPercent(float64(vcp.Current) / float64(vcp.Max))
}

func (m *Model) setContrastPercent(vcp *VCPValue) tea.Cmd {
	return m.models[contrast].SetPercent(float64(vcp.Current) / float64(vcp.Max))
}

func (m *Model) getPercent() float64 {
	s := m.getCurrentSetting()
	return m.models[s].Percent()
}

func (m *Model) setPercent(percent float64) {
	s := m.getCurrentSetting()
	m.models[s].SetPercent(percent)
}

func (m *Model) increment(v float64) (tea.Model, tea.Cmd) {
	s := m.getCurrentSetting()
	cmd := m.models[s].IncrPercent(v)
	return m, tea.Batch(tickCmd(), cmd)
}

func (m *Model) decrement(v float64) (tea.Model, tea.Cmd) {
	s := m.getCurrentSetting()
	cmd := m.models[s].DecrPercent(v)
	return m, tea.Batch(tickCmd(), cmd)
}

func (m *Model) update(msg tea.Msg) (tea.Model, tea.Cmd) {
	s := m.getCurrentSetting()
	progressModel, cmd := m.models[s].Update(msg)
	m.models[s] = progressModel.(progress.Model)
	return m, cmd
}

func (m *Model) updateAll(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, maxSetting)

	var progressModel tea.Model
	for i := range cmds {
		progressModel, cmds[i] = m.models[i].Update(msg)
		m.models[i] = progressModel.(progress.Model)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) loadInitialValues() tea.Cmd {
	cmds := make([]tea.Cmd, maxSetting)

	vcp, err := getVCP(m.display.Index, VCPBrightness)
	if err != nil {
		log.Error("could not get VCP settings: ", err)
		return nil
	}

	cmds[brightness] = m.setBrightnessPercent(vcp)

	vcp, err = getVCP(m.display.Index, VCPContrast)
	if err != nil {
		log.Error("could not get VCP settings: ", err)
		return nil
	}

	cmds[contrast] = m.setContrastPercent(vcp)
	cmds = append([]tea.Cmd{tickCmd()}, cmds...)

	m.initialized = true

	return tea.Batch(cmds...)
}
