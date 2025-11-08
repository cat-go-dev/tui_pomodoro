package mainmenu

import tea "github.com/charmbracelet/bubbletea"

const (
	// quit buttons
	quitButton    = "ctrl+c"
	quitButtonVim = "q"

	// up buttons
	upButton    = "up"
	upButtonVim = "k"

	// down buttons
	downButton    = "down"
	downButtonVim = "j"

	// enter buttons
	enterButton    = "enter"
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
			// todo: emit signal
		}
	}

	return m, nil
}
