package templates

import (
	"fmt"

	"pomodoro/internal/ports"
)

func NewCommonTempate() commonTemplate {
	return commonTemplate{}
}

type commonTemplate struct{}

const (
	cursorSymbol = ">"
	cursorEmpty  = " "
)

func (c commonTemplate) Render(app ports.AppState) string {
	s := fmt.Sprintf("\nCurrent duration: %v\n\n", app.GetDuration())

	for i, choice := range app.GetMenuItems() {
		cursor := cursorEmpty
		if app.GetCursorPosition() == i {
			cursor = cursorSymbol
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	return s
}
