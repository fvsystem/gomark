package localstorage

import (
	"io/ioutil"
	"os"
)

type LocalStorageService struct {
}

var CreateFile = os.Create

var OpenFile = os.Open

var RemoveFile = os.Remove

func (l *LocalStorageService) Save(bytes []byte, path string) error {
	file, err := CreateFile(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)

	return err
}

func (l *LocalStorageService) Load(path string) ([]byte, error) {
	file, err := OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	return bytes, err
}

func (l *LocalStorageService) Delete(path string) error {
	return RemoveFile(path)
}
