package main

import (
	"fmt"

	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/reports/handler"
	"github.com/fvsystem/gomark/internal/application/shared"
	"github.com/fvsystem/gomark/internal/application/test/service"
	"github.com/fvsystem/gomark/internal/infrastructure/cli"
	"github.com/fvsystem/gomark/internal/infrastructure/http"
	"github.com/fvsystem/gomark/internal/infrastructure/json"
	"github.com/fvsystem/gomark/internal/infrastructure/localstorage"
)

var createTest adapter.TestCreator = service.NewCreateTestService()
var httpRequester adapter.Requester = &http.RequesterHTTPImpl{}
var cliExecuter adapter.Executer = &cli.CliExecuter{}
var eventEmitter shared.EventEmitter = shared.NewEventEmitter()
var storageService adapter.Storage = &localstorage.LocalStorageService{}
var reportMaker adapter.ReportMaker = json.NewJsonReportMaker(storageService)
var dataTransformer adapter.DataTransformer = &json.JsonDataTransformer{}
var handlerResultsReturned adapter.Handler = handler.NewHandlerResultsReturned(eventEmitter, dataTransformer)
var testsFinished = make(chan bool)
var handlerResultsTransformed adapter.Handler = handler.NewHandlerResultsTransformed(storageService, reportMaker, eventEmitter, testsFinished)

func main() {
	cliExecuter.Execute(
		createTest,
		httpRequester,
		eventEmitter,
		handlerResultsReturned,
		handlerResultsTransformed,
	)
	<-testsFinished
	fmt.Print("Finished")
}
