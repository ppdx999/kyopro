package atcoder

import "net/http"

type AtcoderRequester struct {
	sessionLoader SessionLoader
}

func NewAtcoderRequester(sessionLoader SessionLoader) *AtcoderRequester {
	return &AtcoderRequester{}
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
