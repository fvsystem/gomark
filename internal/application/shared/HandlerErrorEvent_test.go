package shared

import "testing"

func TestHandlerError_Run(t *testing.T) {
	type args struct {
		event Event
	}
	tests := []struct {
		name string
		e    *HandlerError
		args args
	}{
		{
			name: "TestHandlerError_Run",
			e:    &HandlerError{},
			args: args{
				event: NewEventError("test"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()
			e := &HandlerError{}
			e.Run(tt.args.event)
		})
	}
}
