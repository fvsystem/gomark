package service

import (
	"testing"

	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/shared"
)

type fakeRequester struct {
	called bool
}

func (n *fakeRequester) Get(url string) (adapter.Response, error) {
	n.called = true
	return adapter.Response{
		StatusCode:    200,
		ContentLength: 0,
	}, nil
}

type FakeEventEmitter struct {
	listeners map[string][]func(event shared.Event)
}

func (e *FakeEventEmitter) EmitEvent(event shared.Event) {
	if listeners, ok := e.listeners[event.Name]; ok {
		for _, listener := range listeners {
			listener(event)
		}
	}
}

func (e *FakeEventEmitter) AddListener(event string, listener func(event shared.Event)) {
	e.listeners[event] = append(e.listeners[event], listener)
}

func (e *FakeEventEmitter) RemoveAllListeners(event string) {

}

func TestCreateTestService_CreateTest(t *testing.T) {
	type args struct {
		requester           adapter.Requester
		numberOfConnections int
		eventEmitter        shared.EventEmitter
	}
	var requester *fakeRequester = &fakeRequester{}
	var eventEmitter *FakeEventEmitter = &FakeEventEmitter{
		listeners: make(map[string][]func(event shared.Event)),
	}
	var createTestService adapter.TestCreator = NewCreateTestService()
	tests := []struct {
		name string
		c    adapter.TestCreator
		args args
	}{
		{
			name: "should return response",
			c:    createTestService,
			args: args{
				requester:           requester,
				numberOfConnections: 2,
				eventEmitter:        eventEmitter,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requester.called = false
			c := tt.c
			c.CreateTest(tt.args.requester, tt.args.numberOfConnections, tt.args.eventEmitter)
			if !requester.called {
				t.Errorf("Expected %t, got %t", true, requester.called)
			}
		})
	}
}
