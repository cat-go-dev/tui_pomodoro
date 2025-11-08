package app

import (
	"context"
	"time"

	mainmenu "pomodoro/internal/app/components/main_menu"
	timerimpl "pomodoro/internal/timer"
)

type component interface {
	Render()
}

type components struct {
	MainMenu component
}

type timer interface {
	Start(ctx context.Context, time time.Duration) <-chan struct{}
}

type app struct {
	components *components
	timer      timer
}

func NewApp() *app {
	return &app{
		components: newComponents(),
		timer:      timerimpl.NewTimer(),
	}
}

func newComponents() *components {
	return &components{
		MainMenu: mainmenu.NewComponent(),
	}
}

func (a *app) Run() error {
	a.components.MainMenu.Render()
	return nil
}
