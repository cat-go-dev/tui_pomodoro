package app

import "fmt"

func (m model) View() string {
	s := "Pomodoro timer\n\n"

	for i, choice := range m.menuItems {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}
