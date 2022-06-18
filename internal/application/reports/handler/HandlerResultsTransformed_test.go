package handler

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/shared"
	"github.com/fvsystem/gomark/internal/domain/reports"
)

type FakeReportMaker struct {
}

func (f *FakeReportMaker) MakeReport(interface{}) error {
	return nil
}

type FakeReportMakerError struct {
}

func (f *FakeReportMakerError) MakeReport(interface{}) error {
	return fmt.Errorf("error")
}

type FakeStorageService struct {
}

func (f *FakeStorageService) Save(bytes []byte, path string) error {
	return nil
}

func (f *FakeStorageService) Load(path string) ([]byte, error) {
	return []byte{}, nil
}

func (f *FakeStorageService) Delete(path string) error {
	return nil
}

func TestNewHandlerResultsTransformed(t *testing.T) {
	type args struct {
		storageService adapter.Storage
		reportMaker    adapter.ReportMaker
		eventEmitter   shared.EventEmitter
		testsFinished  chan bool
	}
	tests := []struct {
		name string
		args args
		want shared.EventEmitter
	}{
		{
			name: "TestNewHandlerResultsTransformed",
			args: args{
				storageService: &FakeStorageService{},
				reportMaker:    &FakeReportMaker{},
				eventEmitter:   &FakeEventEmitter{},
				testsFinished:  make(chan bool),
			},
			want: &FakeEventEmitter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandlerResultsTransformed(tt.args.storageService, tt.args.reportMaker, tt.args.eventEmitter, tt.args.testsFinished).eventEmitter; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandlerResultsTransformed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerResultsTransformed_Run(t *testing.T) {
	type fields struct {
		ReportData     reports.Data
		storageService adapter.Storage
		reportMaker    adapter.ReportMaker
		eventEmitter   shared.EventEmitter
		testsFinished  chan bool
	}
	type args struct {
		event shared.Event
	}
	finished := make(chan bool)
	fakeEventEmitter := &FakeEventEmitter{
		listeners: make(map[string][]func(event shared.Event)),
	}
	errorHandler := func(shared.Event) {
		finished <- true
	}
	fakeEventEmitter.AddListener("error", errorHandler)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestHandlerResultsTransformed_Run",
			fields: fields{
				ReportData: reports.Data{
					Results: []adapter.TestResult{},
				},
				storageService: &FakeStorageService{},
				reportMaker:    &FakeReportMaker{},
				eventEmitter:   &FakeEventEmitter{},
				testsFinished:  finished,
			},
			args: args{
				event: shared.Event{
					Name: "ResultsTransformed",
					Data: []adapter.TestResult{},
				},
			},
		},
		{
			name: "TestHandlerResultsTransformed_Run_Error",
			fields: fields{
				ReportData: reports.Data{
					Results: []adapter.TestResult{},
				},
				storageService: &FakeStorageService{},
				reportMaker:    &FakeReportMakerError{},
				eventEmitter:   fakeEventEmitter,
				testsFinished:  finished,
			},
			args: args{
				event: shared.Event{
					Name: "ResultsTransformed",
					Data: []adapter.TestResult{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerResultsTransformed{
				ReportData:     tt.fields.ReportData,
				storageService: tt.fields.storageService,
				reportMaker:    tt.fields.reportMaker,
				eventEmitter:   tt.fields.eventEmitter,
				testsFinished:  tt.fields.testsFinished,
			}
			go h.Run(tt.args.event)
			<-finished
		})
	}
}
