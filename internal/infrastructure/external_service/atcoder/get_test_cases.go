package atcoder

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

func (a *Atcoder) problemPage(cid string, pid string) (*http.Response, error) {
	path := fmt.Sprintf(
		"/contests/%s/tasks/%s_%s",
		cid,
		strings.ReplaceAll(cid, "-", "_"),
		pid,
	)
	return a.get(path)
}

func (a *Atcoder) parseTestCases(html string) ([]*model.TestCase, error) {
	// 入力例・出力例それぞれに続く <pre> の内容を抽出
	re := regexp.MustCompile(`<h3>(入力例|出力例)\s*\d+</h3>\s*<pre>([\s\S]*?)</pre>`)
	matches := re.FindAllStringSubmatch(html, -1)

	if len(matches)%2 != 0 {
		return nil, fmt.Errorf("入力と出力のペアが揃っていません")
	}

	var ts []*model.TestCase
	for i := 0; i < len(matches); i += 2 {
		if matches[i][1] != "入力例" || matches[i+1][1] != "出力例" {
			return nil, fmt.Errorf("入力/出力の順番が想定と異なります")
		}
		id := strconv.Itoa(i/2 + 1)
		input := strings.TrimSpace(matches[i][2]) + "\n"
		want := strings.TrimSpace(matches[i+1][2]) + "\n"
		t := model.NewTestCase(id)
		t.Input = []byte(input)
		t.Want = []byte(want)
		ts = append(ts, t)
	}
	return ts, nil
}

func (a *Atcoder) GetTestCases(cid model.ContestId, pid model.ProblemId) ([]*model.TestCase, error) {
	res, err := a.problemPage(string(cid), string(pid))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ステータスコード: %d", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return a.parseTestCases(string(body))
}
