package adapter

type ResultItem struct {
	Code          int
	Time          int
	Err           bool
	ContentLength int64
}

type TestResult struct {
	Items             []ResultItem
	MaxLatency        int
	MinLatency        int
	StandardDeviation int
	AverageLatency    int
	ContentLengthSent int64
	NumberOfRequests  int
}

type TestInterface interface {
	Start()
	Stop() TestResult
}

type TestCreator interface {
	CreateTest(requester Requester) TestResult
}
