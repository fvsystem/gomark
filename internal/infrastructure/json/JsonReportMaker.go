package json

import (
	"fmt"
	"os"

	"github.com/fvsystem/gomark/internal/adapter"
)

type JsonReportMaker struct {
	storage adapter.Storage
}

var GetPath = os.Getwd

func (j *JsonReportMaker) MakeReport(data interface{}) error {
	path, err := GetPath()
	if err != nil {
		return err
	}
	jsonData, ok := data.([]byte)
	if !ok {
		err = fmt.Errorf("data is not []byte")
		return err
	}
	err = j.storage.Save(jsonData, path+"/report.json")

	return err
}

func NewJsonReportMaker(storage adapter.Storage) adapter.ReportMaker {
	return &JsonReportMaker{
		storage: storage,
	}
}
