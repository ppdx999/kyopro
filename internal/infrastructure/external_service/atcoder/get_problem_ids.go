package atcoder

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

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
