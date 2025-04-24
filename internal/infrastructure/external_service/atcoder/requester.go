package atcoder

import (
	"net/http"

	"github.com/ppdx999/kyopro/internal/domain/service/session"
)

type AtcoderRequester struct {
	sessionLoader session.SessionLoader
}

func NewAtcoderRequester(sessionLoader session.SessionLoader) *AtcoderRequester {
	return &AtcoderRequester{
		sessionLoader: sessionLoader,
	}
}

func (r *AtcoderRequester) Request(req *http.Request) (*http.Response, error) {
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
