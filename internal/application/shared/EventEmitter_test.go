package shared

import (
	"fmt"
	"testing"
)

var fakeEvent Event = Event{
	Name: "test",
	Data: interface{}(nil),
}

var called chan bool = make(chan bool)

var listener func(event Event) = func(event Event) {
	called <- true
}

func TestEventEmitter(t *testing.T) {
	emitter := NewEventEmitter()
	emitter.AddListener("test", listener)
	emitter.EmitEvent(fakeEvent)
	eventCalled := <-called
	if !eventCalled {
		t.Error("Listener was not called")
	}
	called = make(chan bool)
	emitter.RemoveAllListeners("test")
	emitter.EmitEvent(fakeEvent)
	select {
	case <-called:
		err := fmt.Errorf("event called after removal")
		t.Errorf("error %+v", err)
	default:
		t.Log("OK")
	}

}
