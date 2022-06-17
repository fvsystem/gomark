package service

import (
	"testing"

	"github.com/fvsystem/gomark/internal/adapter"
)

type fakeRequester struct {
	called bool
}

func (n *fakeRequester) Get(url string) (adapter.Response, error) {
	n.called = true
	return adapter.Response{
		StatusCode:    200,
		ContentLength: 0,
	}, nil
}

func TestCreateTestService_CreateTest(t *testing.T) {
	type args struct {
		requester adapter.Requester
	}
	var requester *fakeRequester = &fakeRequester{}
	tests := []struct {
		name string
		c    *CreateTestService
		args args
	}{
		{
			name: "should return response",
			c:    &CreateTestService{},
			args: args{
				requester: requester,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requester.called = false
			c := &CreateTestService{}
			c.CreateTest(tt.args.requester)
			if !requester.called {
				t.Errorf("Expected %t, got %t", true, requester.called)
			}
		})
	}
}
