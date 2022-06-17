package main

import (
	"fmt"

	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/test/service"
	"github.com/fvsystem/gomark/internal/infrastructure/cli"
	"github.com/fvsystem/gomark/internal/infrastructure/http"
)

var createTest adapter.TestCreator = &service.CreateTestService{}
var httpRequester adapter.Requester = &http.RequesterHTTPImpl{}
var cliExecuter adapter.Executer = &cli.CliExecuter{}

func main() {
	cliExecuter.Execute(createTest, httpRequester)
	fmt.Print("Finished")
}
