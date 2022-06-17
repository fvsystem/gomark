package http

import (
	"net/http"

	"github.com/fvsystem/gomark/internal/adapter"
)

type RequesterHTTPImpl struct {
}

func (r *RequesterHTTPImpl) Get(url string) (adapter.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		return adapter.Response{}, err
	}

	var response adapter.Response = adapter.Response{
		StatusCode:    resp.StatusCode,
		ContentLength: resp.Request.ContentLength,
	}

	return response, nil
}
