package service

import (
	"errors"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ppdx999/kyopro/internal/application/domain/model"
)

func NewProblems(ids ...string) []*model.Problem {
	var problems []*model.Problem
	for _, id := range ids {
		problems = append(problems, model.NewProblem(id))
	}
	return problems
}

type MockGetProblemIdsPort struct {
	problemIds []model.ProblemId
	err        error
}

func NewMockGetProblemIdsPort() *MockGetProblemIdsPort {
	return &MockGetProblemIdsPort{}
}

func (m *MockGetProblemIdsPort) GetProblemIds(c model.ContestId) ([]model.ProblemId, error) {
	return m.problemIds, m.err
}

type MockMakeProblemDirPort struct {
	createdDirs []string
	err         error
}

func NewMockMakeProblemDirPort() *MockMakeProblemDirPort {
	return &MockMakeProblemDirPort{}
}

func (m *MockMakeProblemDirPort) MakeProblemDir(c model.ContestId, p model.ProblemId) error {
	if m.err != nil {
		return m.err
	}
	m.createdDirs = append(m.createdDirs, filepath.Join(string(c), string(p)))
	return m.err
}

func (m *MockMakeProblemDirPort) CreatedDirs() []string {
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
			getProblemIdsPort := NewMockGetProblemIdsPort()
			getProblemIdsPort.problemIds = tt.getProblemIds
			getProblemIdsPort.err = tt.getProblemIdsErr
			makeProblemDirPort := NewMockMakeProblemDirPort()
			makeProblemDirPort.err = tt.makeProblemDirErr

			service := &InitService{
				getProblemIds:  getProblemIdsPort,
				makeProblemDir: makeProblemDirPort,
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

			gotCreatedDirs := makeProblemDirPort.CreatedDirs()
			if !reflect.DeepEqual(gotCreatedDirs, tt.wantCreatedDirs) {
				t.Errorf("expected created dirs %v, but got %v", tt.wantCreatedDirs, gotCreatedDirs)
			}
		})
	}
}
