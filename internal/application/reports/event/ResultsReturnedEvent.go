package event

import (
	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/shared"
)

type ResultsReturnedEventData struct {
	Results []adapter.TestResult
}

var ResultsReturnedEvent shared.Event = shared.Event{
	Name: "ResultsReturn",
	Data: []adapter.TestResult{},
}
