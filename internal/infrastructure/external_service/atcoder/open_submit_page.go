package atcoder

import (
	"fmt"
	"net/url"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

func (a *Atcoder) OpenSubmitPage(
	c model.ContestId,
	p model.ProblemId,
) error {
	path := fmt.Sprintf("/contests/%s/submit", c)
	q := a.baseUrl.Query()
	q.Add("taskScreenName", fmt.Sprintf("%s_%s", c, p))
	url := a.baseUrl.ResolveReference(&url.URL{Path: path, RawQuery: q.Encode()})
	return a.urlOpener.OpenUrl(url.String())
}
