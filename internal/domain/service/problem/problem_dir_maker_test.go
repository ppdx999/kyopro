package problem_test

import (
	"errors"
	"testing"

	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/problem"
	"github.com/ppdx999/kyopro/internal/testutil"
)

func TestMakeProblemDir(t *testing.T) {
	tests := []struct {
		name             string
		contest          model.ContestId
		problem          model.ProblemId
		wd               string
		getWdErr         error
		makePublicDirErr error
		wantErr          bool
		createdDirs      []string
	}{
		{
			name:    "正常系",
			contest: model.ContestId("abc123"),
			problem: model.ProblemId("1"),
			wd:      "/path/to/workspace",
			createdDirs: []string{
				"/path/to/workspace/abc123/1",
			},
		},
		{
			name:             "MakePublicDirエラー",
			contest:          model.ContestId("abc123"),
			problem:          model.ProblemId("1"),
			wd:               "/path/to/workspace",
			makePublicDirErr: errors.New("make public dir error"),
			createdDirs:      []string{},
			wantErr:          true,
		},
		{
			name:        "GetWdエラー",
			contest:     model.ContestId("abc123"),
			problem:     model.ProblemId("1"),
			getWdErr:    errors.New("get wd error"),
			createdDirs: []string{},
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dirMaker := &testutil.MockMakePublicDir{
				Errs: []error{tt.makePublicDirErr},
			}
			getWd := &testutil.MockGetWd{
				Wd:  tt.wd,
				Err: tt.getWdErr,
			}

			m := problem.NewProblemDirMakerImpl(getWd, dirMaker)

			err := m.MakeProblemDir(tt.contest, tt.problem)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeProblemDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, dir := range tt.createdDirs {
				if !dirMaker.CreatedDirs[dir] {
					t.Errorf("MakeProblemDir() did not create directory %s", dir)
				}
			}
		})
	}
}
