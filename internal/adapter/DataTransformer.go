package adapter

type DataTransformer interface {
	Transform(data []TestResult) (interface{}, error)
}
