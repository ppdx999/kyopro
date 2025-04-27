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
		getWd *MockGetWd
	}
	tests := []struct {
		name string
		mock func(c *gomock.Controller) *mock
		want *model.Problem
	}{
		{
			name: "正常系",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockGetWd {
						m := NewMockGetWd(c)
						m.EXPECT().GetWd().Return("/home/atcoder/contest_A/problem_B", nil)
						return m
					}(),
				}
			},
			want: &model.Problem{ID: "problem_B", Contest: &model.Contest{ID: "contest_A"}},
		},
		{
			name: "ルートディレクトリで実行",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockGetWd {
						m := NewMockGetWd(c)
						m.EXPECT().GetWd().Return("/", nil)
						return m
					}(),
				}
			},
			want: nil,
		},
		{
			name: "問題IDが空",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockGetWd {
						m := NewMockGetWd(c)
						m.EXPECT().GetWd().Return("/contest_A", nil)
						return m
					}(),
				}
			},
		},
		{
			name: "GetWdエラー",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockGetWd {
						m := NewMockGetWd(c)
						m.EXPECT().GetWd().Return("", errors.New("get wd error"))
						return m
					}(),
				}
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mock := tt.mock(mockCtrl)

			l := problem.NewCurrentProblemLoader(mock.getWd)

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
