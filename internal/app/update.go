package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	// quit button
	quitButton    = "ctrl+c"
	quitButtonVim = "q"

	// up button
	upButton    = "up"
	upButtonVim = "k"

	// down button
	downButton    = "down"
	downButtonVim = "j"

	// enter button
	enterButton    = "enter"
	enterButtonAlt = " "
	enterButtonVim = "l"
)

func (app appImpl) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case quitButton, quitButtonVim:
			return app, tea.Quit
		case upButton, upButtonVim:
			if app.cursor > 0 {
				app.cursor--
			}
		case downButton, downButtonVim:
			if app.cursor < len(app.menuItems)-1 {
				app.cursor++
			}
		case enterButton, enterButtonAlt, enterButtonVim:
			// todo: реагировать
		}
	}

	return app, nil
}
