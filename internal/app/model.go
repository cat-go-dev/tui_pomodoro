package app

import (
	"pomodoro/internal/ports"
	"time"
)

type appImpl struct {
	menuItems []string
	cursor    int
	views     views
	state     state
}

const defaultDuration = time.Minute * 25

type state struct {
	duration    time.Duration
	currentView ports.Template
}

type views map[view]ports.Template

var startMenuItems = []string{
	"Start timer",
	"Set duration",
}

func (a appImpl) GetDuration() time.Duration {
	return a.state.duration
}

func (a appImpl) GetMenuItems() []string {
	return a.menuItems
}

func (a appImpl) GetCursorPosition() int {
	return a.cursor
}
