package app

const (
	header = "Pomodoro timer\n\n"
	footer = "\nPress q to quit.\n"
)

type view string

var (
	commonView   view = "COMMON"
	durationForm view = "DURATION_FORN"
)

func (app appImpl) View() string {
	s := header

	s += app.state.currentView.Render(app)

	s += footer

	return s
}
