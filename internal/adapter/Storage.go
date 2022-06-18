package adapter

type Storage interface {
	Save(bytes []byte, path string) error
	Load(path string) ([]byte, error)
	Delete(path string) error
}
