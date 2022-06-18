package service

import (
	"time"

	"github.com/fvsystem/gomark/internal/adapter"
	reportEvent "github.com/fvsystem/gomark/internal/application/reports/event"
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
	eventEmitter     shared.EventEmitter
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
	handlerResultsReturned adapter.Handler,
	handlerResultsTransformed adapter.Handler,
) {
	c.eventEmitter = eventEmitter
	for i := 0; i < numberOfConnections; i++ {
		startHandler := handler.NewHandlerStartTest(requester)
		stopHandler := handler.NewHandlerStopTest(startHandler.Test, c.RegisterResults)
		c.handlers = append(c.handlers, HandlersTest{
			startHandler: startHandler,
			stopHandler:  stopHandler,
		})
		c.eventEmitter.AddListener(event.StartTestEvent.Name, startHandler.Run)
		c.eventEmitter.AddListener(event.StopTestEvent.Name, stopHandler.Run)
	}
	c.eventEmitter.AddListener(reportEvent.ResultsReturnedEvent.Name, handlerResultsReturned.Run)
	c.eventEmitter.AddListener(reportEvent.ResultsTransformedEvent.Name, handlerResultsTransformed.Run)
	c.eventEmitter.EmitEvent(event.StartTestEvent)
	time.Sleep(2 * time.Second)
	eventEmitter.EmitEvent(event.StopTestEvent)
	c.testResults <- []adapter.TestResult{}
}

func (c *CreateTestService) RegisterResults(testResults adapter.TestResult) {
	var currentTestResults []adapter.TestResult = <-c.testResults
	currentTestResults = append(currentTestResults, testResults)
	c.handlersReturned++

	if c.handlersReturned == len(c.handlers) {
		reportEvent.ResultsReturnedEvent.Data = currentTestResults
		c.eventEmitter.RemoveAllListeners(event.StartTestEvent.Name)
		c.eventEmitter.RemoveAllListeners(event.StopTestEvent.Name)
		c.eventEmitter.EmitEvent(reportEvent.ResultsReturnedEvent)
	} else {
		c.testResults <- currentTestResults
	}
}
