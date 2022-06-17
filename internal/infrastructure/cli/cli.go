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
) {

	createTest.CreateTest(requester, 2, eventEmitter)
}
