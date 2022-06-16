package cli

import (
	"github.com/fvsystem/gomark/internal/adapter"

	"fmt"
)

func Execute(createTest adapter.TestCreator) {
	fmt.Println("Executing cli")

	createTest.CreateTest()
}
