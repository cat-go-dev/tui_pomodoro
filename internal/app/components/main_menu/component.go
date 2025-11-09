package mainmenu

import (
	"pomodoro/internal/ports"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor    int
	menuItems []string
	events    chan ports.ComponentEvents
}

var menuItems = []string{
	"Start timer",
	"Set duration",
}

func initialModel() *model {
	return &model{
		menuItems: menuItems,
		events:    make(chan ports.ComponentEvents, 10),
	}
}

func (m *model) Init() tea.Cmd {
	return tea.SetWindowTitle("Pomodoro title")
}

func (m *model) Render() {
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		// todo: log this
	}
}

func (m *model) ListenToEvents() <-chan ports.ComponentEvents {
	return m.events
}

func NewComponent() *model {
	return initialModel()
}
