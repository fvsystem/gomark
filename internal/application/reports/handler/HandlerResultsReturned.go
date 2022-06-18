package handler

import (
	"github.com/fvsystem/gomark/internal/adapter"
	reportsEvent "github.com/fvsystem/gomark/internal/application/reports/event"
	"github.com/fvsystem/gomark/internal/application/shared"
	"github.com/fvsystem/gomark/internal/domain/reports"
)

type HandlerResultsReturned struct {
	ReportData      reports.Data
	eventEmitter    shared.EventEmitter
	dataTransformer adapter.DataTransformer
}

func NewHandlerResultsReturned(eventEmitter shared.EventEmitter, dataTransformer adapter.DataTransformer) *HandlerResultsReturned {

	return &HandlerResultsReturned{
		ReportData: reports.Data{
			Results: []adapter.TestResult{},
		},
		eventEmitter:    eventEmitter,
		dataTransformer: dataTransformer,
	}
}

func (h *HandlerResultsReturned) Run(event shared.Event) {
	results, ok := event.Data.([]adapter.TestResult)
	if !ok {
		return
	}
	h.ReportData.Results = results
	data, error := h.dataTransformer.Transform(h.ReportData.Results)
	if error != nil {
		h.eventEmitter.EmitEvent(shared.NewEventError(error.Error()))
		h.eventEmitter.RemoveAllListeners(event.Name)
		return
	}
	h.ReportData.ResulstTransformed = data

	reportsEvent.ResultsTransformedEvent.Data = h.ReportData.ResulstTransformed

	h.eventEmitter.EmitEvent(reportsEvent.ResultsTransformedEvent)
}
