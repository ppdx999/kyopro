package atcoder

import (
	"net/http"

	"github.com/ppdx999/kyopro/internal/domain/service/session"
)

type Requester interface {
	Request(req *http.Request) (*http.Response, error)
}

type requester struct {
	sessionLoader session.SessionLoader
}

func NewRequester(sessionLoader session.SessionLoader) Requester {
	return &requester{
		sessionLoader: sessionLoader,
	}
}

func (r *requester) Request(req *http.Request) (*http.Response, error) {
	session, err := r.sessionLoader.LoadSession()
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "REVEL_SESSION",
		Value: string(session),
	})

	return http.DefaultClient.Do(req)
}
