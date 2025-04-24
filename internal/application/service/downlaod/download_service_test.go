package downlaod_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ppdx999/kyopro/internal/application/service/downlaod"
	"github.com/ppdx999/kyopro/internal/domain/model"
)

type MockCurrentProblemLoader struct {
	p   *model.Problem
	err error
}

func (m *MockCurrentProblemLoader) LoadCurrentProblem() (*model.Problem, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.p, nil
}

type MockTestCasesGetter struct {
	tcs []*model.TestCase
	err error
}

func (m *MockTestCasesGetter) GetTestCases(
	contestId model.ContestId,
	problemId model.ProblemId,
) (
	[]*model.TestCase,
	error,
) {
	if m.err != nil {
		return nil, m.err
	}
	return m.tcs, nil
}

type MockTestCaseSaver struct {
	savedTestCase []*model.TestCase
	err           error
}

func (m *MockTestCaseSaver) SaveTestCase(t *model.TestCase) error {
	if m.err != nil {
		return m.err
	}
	m.savedTestCase = append(m.savedTestCase, t)
	return nil
}

func TestDownload(t *testing.T) {
	tests := []struct {
		name                  string
		loadCurrentProblem    *model.Problem
		loadCurrentProblemErr error
		getTestCases          []*model.TestCase
		getTestCasesErr       error
		saveTestCaseErr       error
		wantError             bool
	}{
		{
			name: "Correct Case",
			loadCurrentProblem: &model.Problem{
				ID: "a",
				Contest: &model.Contest{
					ID: "abc100",
				},
			},
			getTestCases: []*model.TestCase{
				{
					ID:    "0",
					Input: []byte("input"),
					Want:  []byte("want"),
				},
			},
		},
		{
			name:                  "Load Current Problem Error",
			loadCurrentProblemErr: errors.New("load current problem error"),
			wantError:             true,
		},
		{
			name: "Get Test Cases Error",
			loadCurrentProblem: &model.Problem{
				ID: "a",
				Contest: &model.Contest{
					ID: "abc100",
				},
			},
			getTestCasesErr: errors.New("get test cases error"),
			wantError:       true,
		},
		{
			name: "Save Test Case Error",
			loadCurrentProblem: &model.Problem{
				ID: "a",
				Contest: &model.Contest{
					ID: "abc100",
				},
			},
			getTestCases: []*model.TestCase{
				{
					ID:    "0",
					Input: []byte("input"),
					Want:  []byte("want"),
				},
			},
			saveTestCaseErr: errors.New("save test case error"),
			wantError:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loader := &MockCurrentProblemLoader{
				p:   tt.loadCurrentProblem,
				err: tt.loadCurrentProblemErr,
			}
			getter := &MockTestCasesGetter{
				tcs: tt.getTestCases,
				err: tt.getTestCasesErr,
			}
			saver := &MockTestCaseSaver{
				err: tt.saveTestCaseErr,
			}

			download := downlaod.NewDownloadServiceImpl(
				loader,
				getter,
				saver,
			)

			err := download.Download()

			if (err != nil) != tt.wantError {
				t.Errorf("expected error %v, but got %v", tt.wantError, err)
				return
			}

			if tt.saveTestCaseErr == nil && !reflect.DeepEqual(saver.savedTestCase, tt.getTestCases) {
				t.Errorf("expected saved test cases %v, but got %v", tt.getTestCases, saver.savedTestCase)
			}
		})
	}

}
