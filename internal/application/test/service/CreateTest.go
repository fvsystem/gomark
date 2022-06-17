package service

import (
	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/domain/test"

	"time"
)

type CreateTestService struct {
}

func (c *CreateTestService) CreateTest(requester adapter.Requester) adapter.TestResult {
	testEntity := new(test.TestEntity)
	testEntity.Init(requester)
	go testEntity.Start()
	time.Sleep(1 * time.Second)
	testResults := testEntity.Stop()
	return testResults
}
