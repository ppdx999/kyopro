package testcase

import (
	"path/filepath"
	"strings"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

type testCaseCurrentLoader struct {
	wdGetter             WdGetter
	childFileNamesGetter ChildFileNamesGetter
	publicFileReader     PublicFileReader
}

func (t *testCaseCurrentLoader) LoadCurrentTestCases() ([]*model.TestCase, error) {
	cwd, err := t.wdGetter.GetWd()
	if err != nil {
		return nil, err
	}

	testdir := filepath.Join(cwd, "test")

	files, err := t.childFileNamesGetter.ChildFileNames(testdir)
	if err != nil {
		return nil, err
	}

	var ts []*model.TestCase = make([]*model.TestCase, 0)
	for _, f := range files {
		if filepath.Ext(f) != ".in" {
			continue
		}
		id := strings.TrimSuffix(f, ".in")
		inputPath := filepath.Join(testdir, id+".in")
		input, err := t.publicFileReader.ReadPublicFile(inputPath)
		if err != nil {
			return nil, err
		}

		wantPath := filepath.Join(testdir, id+".out")
		want, err := t.publicFileReader.ReadPublicFile(wantPath)
		if err != nil {
			return nil, err
		}

		ts = append(ts, &model.TestCase{
			ID:    model.TestCaseID(id),
			Input: input,
			Want:  want,
		})
	}

	return ts, nil
}
