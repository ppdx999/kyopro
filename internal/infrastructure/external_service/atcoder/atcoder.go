package atcoder

import (
	"net/url"
)

type Atcoder struct {
	baseUrl   *url.URL
	r         Requester
	urlOpener UrlOpner
}

func NewAtcoder(r Requester, u UrlOpner) *Atcoder {
	baseUrl, _ := url.Parse("https://atcoder.jp")
	return &Atcoder{
		baseUrl:   baseUrl,
		r:         r,
		urlOpener: u,
	}
}
