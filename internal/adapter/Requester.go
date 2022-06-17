package adapter

type Response struct {
	StatusCode    int
	ContentLength int64
}
type Requester interface {
	Get(url string) (Response, error)
}
