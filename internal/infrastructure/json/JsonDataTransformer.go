package json

import (
	"encoding/json"

	"github.com/fvsystem/gomark/internal/adapter"
)

type JsonDataTransformer struct {
}

func (j *JsonDataTransformer) Transform(data []adapter.TestResult) (interface{}, error) {
	dataJson, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return dataJson, nil
}
