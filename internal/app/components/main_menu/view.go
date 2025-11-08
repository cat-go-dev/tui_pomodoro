package mainmenu

import "fmt"

func (m model) View() string {
	return new(view).String(m)
}

type view string

func (v view) String(m model) string {
	return v.getHeader() +
		v.getBody(m) +
		v.getFooter()
}

func (v view) getHeader() string {
	return "Pomodoro timer \n\n"
}

func (v view) getFooter() string {
	return "\nPress q to quit.\n"
}

func (v view) getBody(m model) string {
	var s string

	for i, item := range m.menuItems {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, item)
	}

	return s
}
