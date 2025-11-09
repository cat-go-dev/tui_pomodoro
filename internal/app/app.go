package app

import (
	"context"
	mainmenu "pomodoro/internal/app/components/main_menu"
	"pomodoro/internal/ports"
	timerimpl "pomodoro/internal/timer"
	"time"
)

type components struct {
	MainMenu ports.Component
}

type app struct {
	components *components
	commands   *commands
}

func NewApp() *app {
	return &app{
		components: newComponents(),
		commands:   newCommnads(timerimpl.NewTimer()),
	}
}

func newComponents() *components {
	return &components{
		MainMenu: mainmenu.NewComponent(),
	}
}

func (a *app) Run(ctx context.Context) error {
	a.startListenToComponentEvents(ctx)
	a.renderComponents()

	return nil
}

func (a *app) renderComponents() {
	a.components.MainMenu.Render()
}

func (a *app) startListenToComponentEvents(ctx context.Context) {
	go func() {
		for {
			select {
			case <-a.components.MainMenu.ListenToEvents():
				testDuration := time.Second * 5
				a.commands.StartTimer(ctx, testDuration)
				a.commands.ListenToTimer(ctx)
			case <-ctx.Done():
				// todo: log this
				return
			}
		}
	}()
}
