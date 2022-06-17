package handler

import (
	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/shared"
	"github.com/fvsystem/gomark/internal/domain/test"
)

type HandlerStopTest struct {
	test            *test.TestEntity
	registerResults func(adapter.TestResult)
}

func NewHandlerStopTest(test *test.TestEntity, registerResults func(adapter.TestResult)) *HandlerStopTest {
	return &HandlerStopTest{
		test:            test,
		registerResults: registerResults,
	}
}

func (h *HandlerStopTest) Run(event shared.Event) {
	h.test.Stop(h.registerResults)
}
