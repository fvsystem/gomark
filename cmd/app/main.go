package main

import (
	"fmt"

	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/test/service"
	"github.com/fvsystem/gomark/internal/infrastructure/cli"
)

func main() {
	var createTest adapter.TestCreator = &service.CreateTestService{}

	fmt.Println("Executing cli")
	cli.Execute(createTest)
}
