package service_test

import (
	"errors"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/service"
)

func NewProblems(ids ...string) []*model.Problem {
	var problems []*model.Problem
	for _, id := range ids {
		problems = append(problems, model.NewProblem(id))
	}
	return problems
}

type MockGetProblemIds struct {
	problemIds []model.ProblemId
	err        error
}

func NewMockGetProblemIds() *MockGetProblemIds {
	return &MockGetProblemIds{}
}

func (m *MockGetProblemIds) GetProblemIds(c model.ContestId) ([]model.ProblemId, error) {
	return m.problemIds, m.err
}

type MockMakeProblemDir struct {
	createdDirs []string
	err         error
}

func NewMockMakeProblemDir() *MockMakeProblemDir {
	return &MockMakeProblemDir{}
}

func (m *MockMakeProblemDir) MakeProblemDir(c model.ContestId, p model.ProblemId) error {
	if m.err != nil {
		return m.err
	}
	m.createdDirs = append(m.createdDirs, filepath.Join(string(c), string(p)))
	return m.err
}

func (m *MockMakeProblemDir) CreatedDirs() []string {
	return m.createdDirs
}
func TestInitService(t *testing.T) {
	tests := []struct {
		name              string
		contestId         string
		getProblemIds     []model.ProblemId
		getProblemIdsErr  error
		makeProblemDirErr error
		mayError          bool
		wantError         error
		wantCreatedDirs   []string
	}{
		{
			name:      "正常系",
			contestId: "abc100",
			getProblemIds: []model.ProblemId{
				"a",
				"b",
				"c",
			},
			makeProblemDirErr: nil,
			mayError:          false,
			wantError:         nil,
			wantCreatedDirs: []string{
				"abc100/a",
				"abc100/b",
				"abc100/c",
			},
		},
		{
			name:              "問題一覧取得エラー",
			contestId:         "abc100",
			getProblemIds:     nil,
			getProblemIdsErr:  errors.New("get problem ids error"),
			makeProblemDirErr: nil,
			mayError:          true,
			wantError:         errors.New("get problem ids error"),
			wantCreatedDirs:   nil,
		},
		{
			name:      "問題ディレクトリ作成エラー",
			contestId: "abc100",
			getProblemIds: []model.ProblemId{
				"a",
				"b",
				"c",
			},
			makeProblemDirErr: errors.New("make problem dir error"),
			mayError:          true,
			wantError:         errors.New("make problem dir error"),
			wantCreatedDirs:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getProblemIds := NewMockGetProblemIds()
			getProblemIds.problemIds = tt.getProblemIds
			getProblemIds.err = tt.getProblemIdsErr
			makeProblemDir := NewMockMakeProblemDir()
			makeProblemDir.err = tt.makeProblemDirErr

			service := &service.InitServiceImpl{
				GetProblemIds:  getProblemIds,
				MakeProblemDir: makeProblemDir,
			}

			err := service.Init(model.ContestId(tt.contestId))

			if tt.mayError {
				if err == nil {
					t.Errorf("expected error, but got nil")
					return
				}
				if err.Error() != tt.wantError.Error() {
					t.Errorf("expected error %v, but got %v", tt.wantError, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			gotCreatedDirs := makeProblemDir.CreatedDirs()
			if !reflect.DeepEqual(gotCreatedDirs, tt.wantCreatedDirs) {
				t.Errorf("expected created dirs %v, but got %v", tt.wantCreatedDirs, gotCreatedDirs)
			}
		})
	}
}
