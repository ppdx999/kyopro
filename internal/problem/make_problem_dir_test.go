package problem_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/problem"
)

type MockGetWd struct {
	wd  string
	err error
}

func (m *MockGetWd) GetWd() (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.wd, nil
}

type MockDirMaker struct {
	makedDirs []string
	err       error
}

func (m *MockDirMaker) MakePublicDir(dir string) error {
	if m.err != nil {
		return m.err
	}
	m.makedDirs = append(m.makedDirs, dir)
	return nil
}

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
			dirMaker := &MockDirMaker{
				makedDirs: []string{},
				err:       tt.makePublicDirErr,
			}
			getWd := &MockGetWd{
				wd:  tt.wd,
				err: tt.getWdErr,
			}

			m := problem.NewProblemDirMakerImpl(getWd, dirMaker)

			err := m.MakeProblemDir(tt.contest, tt.problem)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeProblemDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(dirMaker.makedDirs, tt.createdDirs) {
				t.Errorf("MakeProblemDir() createdDirs = %v, want %v", dirMaker.makedDirs, tt.createdDirs)
				return
			}
		})
	}
}
