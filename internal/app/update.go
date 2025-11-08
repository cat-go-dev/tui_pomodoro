package app

import tea "github.com/charmbracelet/bubbletea"

const (
	// quit button
	quitButton       = "ctrl+c"
	quitButtonVim = "q"

	// up button
	upButton       = "up"
	upButtonVim = "k"

	// down button
	downButton       = "down"
	downButtonVim = "j"

	// enter button
	enterButton       = "enter"
	enterButtonAlt = " "
	enterButtonVim = "l"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case quitButton, quitButtonVim:
			return m, tea.Quit
		case upButton, upButtonVim:
			if m.cursor > 0 {
				m.cursor--
			}
		case downButton, downButtonVim:
			if m.cursor < len(m.menuItems)-1 {
				m.cursor++
			}
		case enterButton, enterButtonAlt, enterButtonVim:
			// todo: реагировать
		}
	}

	return m, nil
}
