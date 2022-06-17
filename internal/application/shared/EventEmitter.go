package shared

type EventEmitter interface {
	EmitEvent(event Event)
	AddListener(event string, listener func(event Event))
	RemoveAllListeners(event string)
}

type EventEmitterImpl struct {
	listeners map[string][]func(event Event)
}

func NewEventEmitter() *EventEmitterImpl {
	return &EventEmitterImpl{
		listeners: make(map[string][]func(event Event)),
	}
}

func (e *EventEmitterImpl) EmitEvent(event Event) {
	if listeners, ok := e.listeners[event.Name]; ok {
		for _, listener := range listeners {
			listener(event)
		}
	}
}

func (e *EventEmitterImpl) AddListener(event string, listener func(event Event)) {
	e.listeners[event] = append(e.listeners[event], listener)
}

func (e *EventEmitterImpl) RemoveAllListeners(event string) {
	delete(e.listeners, event)
}
