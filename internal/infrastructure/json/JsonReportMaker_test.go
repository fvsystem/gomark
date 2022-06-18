package json

import (
	"testing"

	"github.com/fvsystem/gomark/internal/adapter"
)

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

var fakeGetWd func() (string, error) = func() (string, error) {
	return "", nil
}

func TestJsonReportMaker_MakeReport(t *testing.T) {
	type fields struct {
		storage adapter.Storage
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestJsonReportMaker_MakeReport",
			fields: fields{
				storage: &FakeStorageService{},
			},
			args: args{
				data: []byte{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPath = fakeGetWd
			j := NewJsonReportMaker(tt.fields.storage)
			if err := j.MakeReport(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("JsonReportMaker.MakeReport() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
