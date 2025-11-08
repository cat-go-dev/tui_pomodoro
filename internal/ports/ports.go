package ports

import "time"

type AppState interface {
	GetDuration() time.Duration
	GetMenuItems() []string
	GetCursorPosition() int
}

type Template interface {
	Render(app AppState) string
}
