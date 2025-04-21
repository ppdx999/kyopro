package atcoder

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/ppdx999/kyopro/internal/model"
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

func (a *Atcoder) url(path string) *url.URL {
	return a.baseUrl.ResolveReference(&url.URL{Path: path})
}

func (a *Atcoder) get(path string) (*http.Response, error) {
	url := a.url(path)
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	return a.r.Request(req)
}

func (a *Atcoder) taskPage(contest string) (*http.Response, error) {
	path := fmt.Sprintf("/contests/%s/tasks", contest)
	return a.get(path)
}

func (a *Atcoder) ExtractProblemIds(html string) []model.ProblemId {
	re := regexp.MustCompile(`/contests/[^/]+/tasks/[^/]+_([a-zA-Z0-9]+)`)
	matches := re.FindAllStringSubmatch(html, -1)

	seen := make(map[string]bool)
	var ids []model.ProblemId
	for _, m := range matches {
		id := strings.ToLower(m[1])
		if !seen[id] {
			ids = append(ids, model.ProblemId(id))
			seen[id] = true
		}
	}
	return ids
}

func (a *Atcoder) GetProblemIds(contest model.ContestId) ([]model.ProblemId, error) {
	resp, err := a.taskPage(string(contest))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return a.ExtractProblemIds(string(body)), nil
}

func (a *Atcoder) LoginCheck() (bool, error) {
	path := "/contests/abc001/submit"
	resp, err := a.get(path)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if resp.Request.URL.Path != path {
		return false, nil
	}

	return true, nil
}
