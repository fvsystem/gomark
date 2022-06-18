package cli

import (
	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/shared"

	"testing"
)

type spyCreateTest struct {
	called bool
}

func (n *spyCreateTest) CreateTest(requester adapter.Requester,
	numberOfConnections int,
	eventEmitter shared.EventEmitter,
	handlerResultsReturned adapter.Handler,
	handlerResultsTransformed adapter.Handler,
) {
	n.called = true
}

type fakeRequester struct{}

func (n *fakeRequester) Get(url string) (adapter.Response, error) {
	return adapter.Response{}, nil
}

type FakeEventEmitter struct{}

func (e *FakeEventEmitter) EmitEvent(event shared.Event) {

}

func (e *FakeEventEmitter) AddListener(event string, listener func(event shared.Event)) {

}

func (e *FakeEventEmitter) RemoveAllListeners(event string) {

}

type FakeHandler struct {
}

func (e *FakeHandler) Run(event shared.Event) {
}

var fakeHandler adapter.Handler = &FakeHandler{}

func TestExecute(t *testing.T) {
	var createTest = &spyCreateTest{called: false}
	var cliExecuter adapter.Executer = &CliExecuter{}
	var eventEmitter = &FakeEventEmitter{}

	cliExecuter.Execute(createTest, &fakeRequester{}, eventEmitter, fakeHandler, fakeHandler)

	if !createTest.called {
		t.Errorf("Expected %t, got %t", true, createTest.called)
	}
}
