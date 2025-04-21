package infra

import "net/http"

type Requester interface {
	Request(req *http.Request) (*http.Response, error)
}

type RequesterImpl struct {
	client *http.Client
}

func (r *RequesterImpl) Request(req *http.Request) (*http.Response, error) {
	return r.client.Do(req)
}

func NewRequesterImpl() *RequesterImpl {
	return &RequesterImpl{
		client: &http.Client{},
	}
}
