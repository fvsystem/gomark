package adapter

import "github.com/fvsystem/gomark/internal/application/shared"

type Executer interface {
	Execute(
		testCreator TestCreator,
		requester Requester,
		eventEmitter shared.EventEmitter,
		handlerResultsReturned Handler,
		handlerResultsTransformed Handler,
	)
}
