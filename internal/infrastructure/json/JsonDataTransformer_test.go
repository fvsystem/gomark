package json

import (
	"reflect"
	"testing"

	"github.com/fvsystem/gomark/internal/adapter"
)

var fakeMarshal = func(v any) ([]byte, error) {
	return []byte{}, nil
}

func TestJsonDataTransformer_Transform(t *testing.T) {
	type args struct {
		data []adapter.TestResult
	}
	tests := []struct {
		name    string
		j       *JsonDataTransformer
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "TestJsonDataTransformer_Transform",
			j:    &JsonDataTransformer{},
			args: args{
				data: []adapter.TestResult{},
			},
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JsonTransformer = fakeMarshal
			j := &JsonDataTransformer{}
			got, err := j.Transform(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonDataTransformer.Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDataTransformer.Transform() = %v, want %v", got, tt.want)
			}
		})
	}
}
