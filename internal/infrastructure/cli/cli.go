package cli

import (
	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/shared"
)

type CliExecuter struct{}

func (c *CliExecuter) Execute(
	createTest adapter.TestCreator,
	requester adapter.Requester,
	eventEmitter shared.EventEmitter,
	handlerResultsReturned adapter.Handler,
	handlerResultsTransformed adapter.Handler,
) {

	createTest.CreateTest(requester, 2, eventEmitter, handlerResultsReturned, handlerResultsTransformed)
}
