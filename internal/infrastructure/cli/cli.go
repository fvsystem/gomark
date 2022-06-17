package cli

import (
	"github.com/fvsystem/gomark/internal/adapter"
)

type CliExecuter struct{}

func (c *CliExecuter) Execute(createTest adapter.TestCreator, requester adapter.Requester) {

	createTest.CreateTest(requester)
}
