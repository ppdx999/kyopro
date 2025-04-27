package atcoder

import (
	"net/url"
)

type Atcoder struct {
	baseUrl *url.URL
	r       Requester
}

func NewAtcoder(r Requester) *Atcoder {
	baseUrl, _ := url.Parse("https://atcoder.jp")
	return &Atcoder{
		baseUrl: baseUrl,
		r:       r,
	}
}
