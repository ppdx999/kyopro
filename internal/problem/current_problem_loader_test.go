package problem_test

import (
	"errors"
	"testing"

	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/problem"
	"github.com/ppdx999/kyopro/internal/testutil"
)

func TestLoadCurrentProblem(t *testing.T) {
	tests := []struct {
		name     string
		getWd    string
		getWdErr error
		want     *model.Problem
		wantErr  bool
	}{
		{
			name:     "正常系",
			getWd:    "/home/atcoder/contest_A/problem_B",
			getWdErr: nil,
			want: &model.Problem{
				ID:      "problem_B",
				Contest: &model.Contest{ID: "contest_A"},
			},
		},
		{
			name:     "ルートディレクトリで実行",
			getWd:    "/",
			getWdErr: errors.New("contest or problem not found"),
			wantErr:  true,
		},
		{
			name:     "問題IDが空",
			getWd:    "/contest_A",
			getWdErr: errors.New("contest or problem not found"),
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getWd := &testutil.MockGetWd{
				Wd:  tt.getWd,
				Err: tt.getWdErr,
			}
			l := problem.NewCurrentProblemLoaderImpl(getWd)

			got, err := l.LoadCurrentProblem()

			if (err != nil) != tt.wantErr {
				t.Errorf("LoadCurrentProblem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got.ID != tt.want.ID {
					t.Errorf("LoadCurrentProblem() problemId= %v, want %v", got.ID, tt.want.ID)
					return
				}
				if got.Contest.ID != tt.want.Contest.ID {
					t.Errorf("LoadCurrentProblem() contestId= %v, want %v", got.Contest.ID, tt.want.Contest.ID)
					return
				}
			}
		})
	}
}
