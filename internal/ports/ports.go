package ports

type ComponentEvents string

const (
	EventStartTimer ComponentEvents = "START_TIMER"
)

type Component interface {
	Render()
	ListenToEvents() <-chan ComponentEvents
}
