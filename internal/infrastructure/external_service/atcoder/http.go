package atcoder

import "net/http"

func (a *Atcoder) get(path string) (*http.Response, error) {
	url := a.url(path)
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	return a.r.Request(req)
}
