package localstorage

import (
	"io/ioutil"
	"os"
)

type LocalStorageService struct {
}

func (l *LocalStorageService) Save(bytes []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)

	return err
}

func (l *LocalStorageService) Load(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	return bytes, err
}

func (l *LocalStorageService) Delete(path string) error {
	return os.Remove(path)
}
