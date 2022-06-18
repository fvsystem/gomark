package localstorage

import (
	"os"
	"path"
	"reflect"
	"testing"
)

type MockFile struct {
}

func (m *MockFile) Close() error {
	return nil
}

func TestLocalStorageService_Save(t *testing.T) {
	type args struct {
		bytes []byte
		path  string
	}
	tests := []struct {
		name    string
		l       *LocalStorageService
		args    args
		wantErr bool
	}{
		{
			name: "TestLocalStorageService_Save",
			l:    &LocalStorageService{},
			args: args{
				bytes: []byte{},
				path:  "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			l := &LocalStorageService{}
			if err := l.Save(tt.args.bytes, path.Join(tempDir, "test.txt")); (err != nil) != tt.wantErr {
				t.Errorf("LocalStorageService.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalStorageService_Load(t *testing.T) {
	tests := []struct {
		name    string
		l       *LocalStorageService
		want    []byte
		wantErr bool
	}{
		{
			name:    "TestLocalStorageService_Load",
			l:       &LocalStorageService{},
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			file, err := os.Create(path.Join(tempDir, "test.txt"))
			if err != nil {
				t.Errorf("LocalStorageService.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			file.Close()
			l := &LocalStorageService{}
			got, err := l.Load(path.Join(tempDir, "test.txt"))
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalStorageService.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocalStorageService.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
