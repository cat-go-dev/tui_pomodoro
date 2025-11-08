package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

func NewApp() model {
	return model{
		menuItems: startMenuItems,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
