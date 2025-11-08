package app

import (
	"pomodoro/internal/app/templates"
	"pomodoro/internal/ports"

	tea "github.com/charmbracelet/bubbletea"
)

func NewApp() appImpl {
	app := appImpl{
		menuItems: startMenuItems,
		views: getViews(),
		state: state{
			duration: defaultDuration,
		},
	}
	app.state.currentView = app.views[commonView]

	return app
}

func getViews() views {
	return map[view]ports.Template{
		commonView: templates.NewCommonTempate(),
	}
}

func (app appImpl) Init() tea.Cmd {
	return nil
}
