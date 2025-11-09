package app

import (
	"context"
	"os/exec"
	"time"

	mainmenu "pomodoro/internal/app/components/main_menu"
	"pomodoro/internal/ports"
	timerimpl "pomodoro/internal/timer"
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
			case event := <-a.components.MainMenu.ListenToEvents():
				switch event {
				case ports.EventStartTimer:
					testDuration := time.Second * 5
					a.commands.StartTimer(ctx, testDuration)
					a.commands.ListenToTimer(ctx)

					// sound signal
					// todo: add sound component
					exec.Command("afplay", "/System/Library/Sounds/Ping.aiff").Run()
				case ports.EventStopTimer:
					//todo: stop timer
				}
			case <-ctx.Done():
				// todo: log this
				return
			}
		}
	}()
}
