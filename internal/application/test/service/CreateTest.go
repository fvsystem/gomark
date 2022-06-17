package service

import (
	"fmt"
	"time"

	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/shared"
	"github.com/fvsystem/gomark/internal/application/test/event"
	"github.com/fvsystem/gomark/internal/application/test/handler"
)

type HandlersTest struct {
	startHandler *handler.HandlerStartTest
	stopHandler  *handler.HandlerStopTest
}

type CreateTestService struct {
	testResults      chan []adapter.TestResult
	handlers         []HandlersTest
	handlersReturned int
}

func NewCreateTestService() adapter.TestCreator {
	return &CreateTestService{
		testResults:      make(chan ([]adapter.TestResult)),
		handlers:         []HandlersTest{},
		handlersReturned: 0,
	}
}

func (c *CreateTestService) CreateTest(
	requester adapter.Requester,
	numberOfConnections int,
	eventEmitter shared.EventEmitter,
) {
	for i := 0; i < numberOfConnections; i++ {
		startHandler := handler.NewHandlerStartTest(requester)
		stopHandler := handler.NewHandlerStopTest(startHandler.Test, c.RegisterResults)
		c.handlers = append(c.handlers, HandlersTest{
			startHandler: startHandler,
			stopHandler:  stopHandler,
		})
		eventEmitter.AddListener(event.StartTestEvent.Name, startHandler.Run)
		eventEmitter.AddListener(event.StopTestEvent.Name, stopHandler.Run)
	}
	eventEmitter.EmitEvent(event.StartTestEvent)
	time.Sleep(2 * time.Second)
	go eventEmitter.EmitEvent(event.StopTestEvent)
	c.testResults <- []adapter.TestResult{}
}

func (c *CreateTestService) RegisterResults(testResults adapter.TestResult) {
	var currentTestResults []adapter.TestResult = <-c.testResults
	currentTestResults = append(currentTestResults, testResults)
	c.handlersReturned++

	if c.handlersReturned == len(c.handlers) {
		//TODO will trigger an event to execute reports
		fmt.Print("Tests finished")
	} else {
		c.testResults <- currentTestResults
	}
}
