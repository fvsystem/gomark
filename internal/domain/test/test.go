package test

import (
	"github.com/fvsystem/gomark/internal/adapter"

	"math"
	"strconv"
	"time"
)

type TestEntity struct {
	port       int
	host       string
	start      chan bool
	path       string
	testResult adapter.TestResult
	requester  adapter.Requester
}

func (t *TestEntity) Init(requester adapter.Requester) {
	t.start = make(chan bool)
	t.testResult = adapter.TestResult{
		Items: []adapter.ResultItem{},
	}
	t.host = "http://localhost"
	t.port = 8080
	t.requester = requester
	t.path = "/test"
}

func (t *TestEntity) Start() {
	shouldContinue := true
	for shouldContinue {
		start := time.Now()
		var errHappened bool = false
		response, err := t.requester.Get(t.host + ":" + strconv.Itoa(t.port) + t.path)
		if err != nil {
			select {
			case shouldContinue = <-t.start:
			default:
				shouldContinue = true
			}
			continue
		}
		elapsed := time.Since(start)
		if response.StatusCode < 200 || response.StatusCode > 299 {
			errHappened = true
		}
		t.testResult.Items = append(t.testResult.Items, adapter.ResultItem{
			Code:          response.StatusCode,
			Time:          int(elapsed.Milliseconds()),
			Err:           errHappened,
			ContentLength: response.ContentLength,
		})

		select {
		case shouldContinue = <-t.start:
		default:
			shouldContinue = true
		}
	}
}

func (t *TestEntity) Stop(registerResults func(testResults adapter.TestResult)) {
	t.start <- false
	if len(t.testResult.Items) == 0 {
		t.testResult.MaxLatency = 0
		t.testResult.MinLatency = 0
		t.testResult.AverageLatency = 0
		t.testResult.ContentLengthSent = 0
		t.testResult.NumberOfRequests = 0
		t.testResult.StandardDeviation = 0
		registerResults(t.testResult)
		return
	}
	var max int = t.testResult.Items[0].Time
	var min int = t.testResult.Items[0].Time
	var sum int = 0
	var sumContentLength int64 = 0
	var lengthWithoutErrors int = 0

	for _, item := range t.testResult.Items {
		if !item.Err {
			if item.Time > max {
				max = item.Time
			}
			if item.Time < min {
				min = item.Time
			}
			sum += item.Time
			sumContentLength += item.ContentLength
			lengthWithoutErrors++
		}
	}

	t.testResult.MaxLatency = max
	t.testResult.MinLatency = min
	if sum > 0 {
		t.testResult.AverageLatency = sum / lengthWithoutErrors
	} else {
		t.testResult.AverageLatency = 0
	}
	t.testResult.ContentLengthSent = sumContentLength
	t.testResult.NumberOfRequests = len(t.testResult.Items)

	var sumDeviation int = 0

	for _, item := range t.testResult.Items {
		if !item.Err {
			sumDeviation += (item.Time - t.testResult.AverageLatency) * (item.Time - t.testResult.AverageLatency)
		}
	}
	if sumDeviation > 0 {
		t.testResult.StandardDeviation = int(math.Sqrt(float64(sumDeviation / lengthWithoutErrors)))
	} else {
		t.testResult.StandardDeviation = 0
	}
	registerResults(t.testResult)
}
