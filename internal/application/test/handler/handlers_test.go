package handler

import (
	"testing"

	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/test/event"
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

var called bool = false

var registerResults = func(adapter.TestResult) {
	called = true
}

func TestHandlers(t *testing.T) {
	startHandler := NewHandlerStartTest(&fakeRequester{})
	stopHandler := NewHandlerStopTest(startHandler.Test, registerResults)
	startHandler.Run(event.StartTestEvent)
	stopHandler.Run(event.StopTestEvent)
	if !called {
		t.Error("Handler was not called")
	}
}
