package event

import (
	"github.com/fvsystem/gomark/internal/application/shared"
)

type ResultsTransformedEventData struct {
	ResultsTransformed any
}

var ResultsTransformedEvent shared.Event = shared.Event{
	Name: "ResultsTransformed",
	Data: ResultsTransformedEventData{},
}
