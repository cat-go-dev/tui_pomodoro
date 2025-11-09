package ports

type ComponentEvents string

const (
	EventStartTimer ComponentEvents = "START_TIMER"
	EventStopTimer  ComponentEvents = "STOP_TIMER"
)

type Component interface {
	Render()
	ListenToEvents() <-chan ComponentEvents
}
