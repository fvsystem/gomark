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

func TestExecute(t *testing.T) {
	var createTest = &spyCreateTest{called: false}
	var cliExecuter adapter.Executer = &CliExecuter{}
	var eventEmitter = &FakeEventEmitter{}

	cliExecuter.Execute(createTest, &fakeRequester{}, eventEmitter)

	if !createTest.called {
		t.Errorf("Expected %t, got %t", true, createTest.called)
	}
}
