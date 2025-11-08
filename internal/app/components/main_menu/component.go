package mainmenu

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor    int
	menuItems []string
}

var menuItems = []string{
	"Start timer",
	"Set duration",
}

func initialModel() model {
	return model{
		menuItems: menuItems,
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Pomodoro title")
}

func (m model) Render() {
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		// todo: log this
	}
}

func NewComponent() model {
	return initialModel()
}
