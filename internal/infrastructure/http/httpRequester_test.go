package http

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/fvsystem/gomark/internal/adapter"
)

func TestRequesterHTTPImpl_Get(t *testing.T) {
	type args struct {
		url string
	}
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
	}))
	tests := []struct {
		name    string
		r       *RequesterHTTPImpl
		args    args
		want    adapter.Response
		wantErr bool
	}{
		{
			name: "should return response",
			r:    &RequesterHTTPImpl{},
			args: args{
				url: testServer.URL,
			},
			want: adapter.Response{
				StatusCode:    200,
				ContentLength: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RequesterHTTPImpl{}
			got, err := r.Get(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequesterHTTPImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequesterHTTPImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
