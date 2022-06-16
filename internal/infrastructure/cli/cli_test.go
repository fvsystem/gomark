package cli

import (
	"testing"
)

type spyCreateTest struct {
	called bool
}

func (n *spyCreateTest) CreateTest() {
	n.called = true
}

func TestExecute(t *testing.T) {
	var createTest = &spyCreateTest{called: false}

	Execute(createTest)

	if !createTest.called {
		t.Errorf("Expected %t, got %t", true, createTest.called)
	}
}
