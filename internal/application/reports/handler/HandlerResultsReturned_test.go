package handler

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fvsystem/gomark/internal/adapter"
	"github.com/fvsystem/gomark/internal/application/shared"
	"github.com/fvsystem/gomark/internal/domain/reports"
)

type FakeEventEmitter struct {
	listeners map[string][]func(event shared.Event)
	called    bool
}

func (e *FakeEventEmitter) EmitEvent(event shared.Event) {
	e.called = true
	if listeners, ok := e.listeners[event.Name]; ok {
		for _, listener := range listeners {
			go listener(event)
		}
	}
}

func (e *FakeEventEmitter) AddListener(event string, listener func(event shared.Event)) {
	e.listeners[event] = append(e.listeners[event], listener)
}

func (e *FakeEventEmitter) RemoveAllListeners(event string) {

}

type FakeDataTransformerWithError struct {
}

func (f *FakeDataTransformerWithError) Transform(data []adapter.TestResult) (interface{}, error) {
	return nil, fmt.Errorf("error")
}

type FakeDataTransformer struct {
}

func (f *FakeDataTransformer) Transform(data []adapter.TestResult) (interface{}, error) {
	return data, nil
}

func TestNewHandlerResultsReturned(t *testing.T) {
	type args struct {
		eventEmitter    shared.EventEmitter
		dataTransformer adapter.DataTransformer
	}
	tests := []struct {
		name string
		args args
		want shared.EventEmitter
	}{
		{
			name: "TestNewHandlerResultsReturned",
			args: args{
				eventEmitter:    &FakeEventEmitter{},
				dataTransformer: &FakeDataTransformer{},
			},
			want: &FakeEventEmitter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandlerResultsReturned(tt.args.eventEmitter, tt.args.dataTransformer).eventEmitter; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandlerResultsReturned() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerResultsReturned_Run(t *testing.T) {
	type fields struct {
		ReportData      reports.Data
		eventEmitter    shared.EventEmitter
		dataTransformer adapter.DataTransformer
	}
	type args struct {
		event shared.Event
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestHandlerResultsReturned_Run",
			fields: fields{
				ReportData: reports.Data{
					Results: []adapter.TestResult{},
				},
				eventEmitter:    &FakeEventEmitter{},
				dataTransformer: &FakeDataTransformer{},
			},
			args: args{
				event: shared.Event{
					Name: "ResultsReturned",
					Data: []adapter.TestResult{},
				},
			},
		},
		{
			name: "TestHandlerResultsReturned_RunWithError",
			fields: fields{
				ReportData: reports.Data{
					Results: []adapter.TestResult{},
				},
				eventEmitter:    &FakeEventEmitter{},
				dataTransformer: &FakeDataTransformerWithError{},
			},
			args: args{
				event: shared.Event{
					Name: "ResultsReturned",
					Data: []adapter.TestResult{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerResultsReturned{
				ReportData:      tt.fields.ReportData,
				eventEmitter:    tt.fields.eventEmitter,
				dataTransformer: tt.fields.dataTransformer,
			}
			h.Run(tt.args.event)
		})
		fakeEventEmitter, ok := tt.fields.eventEmitter.(*FakeEventEmitter)
		if !ok {
			t.Errorf("HandlerResultsReturned.Run() error = %v", ok)
		}
		if !fakeEventEmitter.called {
			t.Errorf("HandlerResultsReturned.Run() = %v, want %v", fakeEventEmitter.called, true)
		}
	}
}
