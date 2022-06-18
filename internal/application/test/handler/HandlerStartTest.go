package handler

import (
	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/shared"
	"github.com/fvsystem/gomark/internal/domain/test"
)

type HandlerStartTest struct {
	Test *test.TestEntity
}

func NewHandlerStartTest(requester adapter.Requester) *HandlerStartTest {
	testEntity := new(test.TestEntity)
	testEntity.Init(requester)
	return &HandlerStartTest{
		Test: testEntity,
	}
}

func (h *HandlerStartTest) Run(event shared.Event) {
	h.Test.Start()
}
