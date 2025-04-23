package atcoder

import (
	"net/url"

	"github.com/ppdx999/kyopro/internal/requester"
)

type Atcoder struct {
	baseUrl *url.URL
	r       requester.Requester
}

func NewAtcoder(r requester.Requester) *Atcoder {
	baseUrl, _ := url.Parse("https://atcoder.jp")
	return &Atcoder{
		baseUrl: baseUrl,
		r:       r,
	}
}
