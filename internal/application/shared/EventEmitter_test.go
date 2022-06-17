package shared

import (
	"testing"
)

var fakeEvent Event = Event{
	Name: "test",
	Data: interface{}(nil),
}

var called bool = false

var listener func(event Event) = func(event Event) {
	called = true
}

func TestEventEmitter(t *testing.T) {
	emitter := NewEventEmitter()
	emitter.AddListener("test", listener)
	emitter.EmitEvent(fakeEvent)
	if !called {
		t.Error("Listener was not called")
	}
	called = false
	emitter.RemoveAllListeners("test")
	emitter.EmitEvent(fakeEvent)
	if called {
		t.Error("Listener was called after removal")
	}

}
