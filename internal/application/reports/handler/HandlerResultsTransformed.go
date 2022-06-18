package handler

import (
	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/shared"
	"github.com/fvsystem/gomark/internal/domain/reports"
)

type HandlerResultsTransformed struct {
	ReportData     reports.Data
	storageService adapter.Storage
	reportMaker    adapter.ReportMaker
	eventEmitter   shared.EventEmitter
	testsFinished  chan bool
}

func NewHandlerResultsTransformed(
	storageService adapter.Storage,
	reportMaker adapter.ReportMaker,
	eventEmitter shared.EventEmitter,
	testsFinished chan bool,
) *HandlerResultsTransformed {

	return &HandlerResultsTransformed{
		ReportData: reports.Data{
			Results: []adapter.TestResult{},
		},
		storageService: storageService,
		reportMaker:    reportMaker,
		eventEmitter:   eventEmitter,
		testsFinished:  testsFinished,
	}
}

func (h *HandlerResultsTransformed) Run(event shared.Event) {
	error := h.reportMaker.MakeReport(event.Data)
	if error != nil {
		h.eventEmitter.EmitEvent(shared.NewEventError(error.Error()))
	}
	h.testsFinished <- true
}
