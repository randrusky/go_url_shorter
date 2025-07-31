package event

const (
	EventLinkVisited = "link.visited"
)

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus: make(chan Event), // Buffered channel to handle events
	}
}

func (e *EventBus) Publish(event Event) {
	e.bus <- event // Send the event to the channel
}

func (e *EventBus) Subscribe() <-chan Event {
	return e.bus // Return the channel to subscribe to events
}
