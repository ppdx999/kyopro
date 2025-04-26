package problem_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/problem"
)

func TestLoadCurrentProblem(t *testing.T) {
	type mock struct {
		getWd    string
		getWdErr error
	}
	tests := []struct {
		name string
		mock *mock
		want *model.Problem
	}{
		{
			name: "正常系",
			mock: &mock{getWd: "/home/atcoder/contest_A/problem_B"},
			want: &model.Problem{ID: "problem_B", Contest: &model.Contest{ID: "contest_A"}},
		},
		{
			name: "ルートディレクトリで実行",
			mock: &mock{getWd: "/"},
		},
		{
			name: "問題IDが空",
			mock: &mock{getWd: "/contest_A"},
		},
		{
			name: "GetWdエラー",
			mock: &mock{getWdErr: errors.New("get wd error")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			getWd := NewMockGetWd(mockCtrl)
			getWd.EXPECT().GetWd().Return(tt.mock.getWd, tt.mock.getWdErr)
			l := problem.NewCurrentProblemLoaderImpl(getWd)

			// Act
			got, err := l.LoadCurrentProblem()

			// Assert
			if tt.want == nil {
				if err == nil {
					t.Error("CurrentProblemLoaderImpl.LoadCurrentProblem() error is expected but got nil")
					return
				}
				return
			}

			if got.ID != tt.want.ID {
				t.Errorf("CurrentProblemLoaderImpl.LoadCurrentProblem() problemId= %v, want %v", got.ID, tt.want.ID)
				return
			}
			if got.Contest.ID != tt.want.Contest.ID {
				t.Errorf("LoadCurrentProblem() contestId= %v, want %v", got.Contest.ID, tt.want.Contest.ID)
				return
			}
		})
	}
}
