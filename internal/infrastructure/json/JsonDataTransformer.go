package json

import (
	"encoding/json"

	"github.com/fvsystem/gomark/internal/adapter"
)

type JsonDataTransformer struct {
}

var JsonTransformer = json.Marshal

func (j *JsonDataTransformer) Transform(data []adapter.TestResult) (interface{}, error) {
	dataJson, err := JsonTransformer(data)

	if err != nil {
		return nil, err
	}

	return dataJson, nil
}
