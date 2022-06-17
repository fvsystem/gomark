package test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/fvsystem/gomark/internal/adapter"
)

type mockHTTPRequester struct {
}

func (m *mockHTTPRequester) Get(url string) (adapter.Response, error) {
	return adapter.Response{
		StatusCode:    200,
		ContentLength: 100,
	}, nil

}

type mockHTTPRequesterWithHTTPError struct {
}

func (m *mockHTTPRequesterWithHTTPError) Get(url string) (adapter.Response, error) {
	return adapter.Response{
		StatusCode:    400,
		ContentLength: 100,
	}, nil

}

type mockHTTPRequesterWithError struct {
	called bool
}

func (m *mockHTTPRequesterWithError) Get(url string) (adapter.Response, error) {
	m.called = true
	return adapter.Response{}, errors.New("error")

}

func fakeStart(channel chan bool) {
	<-channel
}

func TestTestEntity_Stop(t *testing.T) {
	type fields struct {
		port       int
		host       string
		start      chan bool
		testResult adapter.TestResult
	}
	start := make(chan bool)
	tests := []struct {
		name   string
		fields fields
		want   adapter.TestResult
	}{
		{
			name: "TestEntity_Stop",
			fields: fields{
				port:  8080,
				host:  "localhost",
				start: start,
				testResult: adapter.TestResult{
					Items: []adapter.ResultItem{
						{Code: 200, Time: 100, Err: false, ContentLength: 100},
						{Code: 200, Time: 200, Err: false, ContentLength: 200},
						{Code: 200, Time: 300, Err: true, ContentLength: 300},
					},
				},
			}, want: adapter.TestResult{
				Items: []adapter.ResultItem{
					{Code: 200, Time: 100, Err: false, ContentLength: 100},
					{Code: 200, Time: 200, Err: false, ContentLength: 200},
					{Code: 200, Time: 300, Err: true, ContentLength: 300},
				},
				MaxLatency:        200,
				MinLatency:        100,
				StandardDeviation: 50,
				AverageLatency:    150,
				ContentLengthSent: 300,
				NumberOfRequests:  3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TestEntity{
				port:       tt.fields.port,
				host:       tt.fields.host,
				start:      tt.fields.start,
				testResult: tt.fields.testResult,
			}
			go fakeStart(tt.fields.start)
			if got := tr.Stop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestEntity.Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestEntity_Start(t *testing.T) {
	type fields struct {
		port       int
		host       string
		start      chan bool
		testResult adapter.TestResult
		requester  adapter.Requester
	}
	start := make(chan bool)
	requestWithError := &mockHTTPRequesterWithError{}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "TestEntity_Start",
			fields: fields{
				port:  8080,
				host:  "localhost",
				start: start,
				testResult: adapter.TestResult{
					Items: []adapter.ResultItem{},
				},
				requester: &mockHTTPRequester{},
			},
		},
		{
			name: "TestEntity_Start_WithHTTPError",
			fields: fields{
				port:  8080,
				host:  "localhost",
				start: start,
				testResult: adapter.TestResult{
					Items: []adapter.ResultItem{},
				},
				requester: &mockHTTPRequesterWithHTTPError{},
			},
		},
		{
			name: "TestEntity_Start_WithError",
			fields: fields{
				port:  8080,
				host:  "localhost",
				start: start,
				testResult: adapter.TestResult{
					Items: []adapter.ResultItem{},
				},
				requester: requestWithError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TestEntity{
				port:       tt.fields.port,
				host:       tt.fields.host,
				start:      tt.fields.start,
				testResult: tt.fields.testResult,
				requester:  tt.fields.requester,
			}
			go tr.Start()

			time.Sleep(200 * time.Millisecond)

			tr.start <- false

			fmt.Print("called", requestWithError.called)

			if len(tr.testResult.Items) == 0 && !requestWithError.called {
				t.Errorf("TestEntity.Start() = %v, want %v", len(tr.testResult.Items), ">0")
			}
		})
	}
}

func TestTestEntity_Init(t *testing.T) {
	type fields struct {
		port       int
		host       string
		start      chan bool
		testResult adapter.TestResult
		requester  adapter.Requester
	}
	type args struct {
		requester adapter.Requester
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "should init test",
			fields: fields{
				port:       8080,
				host:       "localhost",
				start:      make(chan bool),
				testResult: adapter.TestResult{},
			},
			args: args{
				requester: &mockHTTPRequester{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TestEntity{
				port:       tt.fields.port,
				host:       tt.fields.host,
				start:      tt.fields.start,
				testResult: tt.fields.testResult,
				requester:  tt.fields.requester,
			}
			tr.Init(tt.args.requester)
			if tr.port != 8080 {
				t.Errorf("TestEntity.Init() = %v, want %v", tr.port, 8080)
			}
		})
	}
}
