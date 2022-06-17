package cli

import (
	"github.com/fvsystem/gomark/internal/adapter"

	"testing"
)

type spyCreateTest struct {
	called bool
}

func (n *spyCreateTest) CreateTest(requester adapter.Requester) adapter.TestResult {
	n.called = true
	return adapter.TestResult{}
}

type fakeRequester struct{}

func (n *fakeRequester) Get(url string) (adapter.Response, error) {
	return adapter.Response{}, nil
}

func TestExecute(t *testing.T) {
	var createTest = &spyCreateTest{called: false}
	var cliExecuter adapter.Executer = &CliExecuter{}

	cliExecuter.Execute(createTest, &fakeRequester{})

	if !createTest.called {
		t.Errorf("Expected %t, got %t", true, createTest.called)
	}
}
