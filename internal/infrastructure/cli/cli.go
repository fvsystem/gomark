package cli

import (
	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/test/service"

	"fmt"
)

func Execute() {
	fmt.Println("Executing cli")

	var serviceCreateTest adapter.TestCreator = &service.CreateTestService{}
	serviceCreateTest.CreateTest()
}
