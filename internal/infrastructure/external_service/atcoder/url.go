package atcoder

import "net/url"

func (a *Atcoder) url(path string) *url.URL {
	return a.baseUrl.ResolveReference(&url.URL{Path: path})
}
