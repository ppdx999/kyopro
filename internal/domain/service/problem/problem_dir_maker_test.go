package problem_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/problem"
)

func TestMakeProblemDir(t *testing.T) {
	type args struct {
		contest model.ContestId
		problem model.ProblemId
	}
	type mock struct {
		getWd         *MockGetWd
		makePublicDir *MockPublicDirMaker
	}
	tests := []struct {
		name    string
		args    *args
		mock    func(c *gomock.Controller) *mock
		wantErr bool
	}{
		{
			name: "正常系",
			args: &args{
				contest: model.ContestId("abc123"),
				problem: model.ProblemId("1"),
			},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockGetWd {
						m := NewMockGetWd(c)
						m.EXPECT().GetWd().Return("/path/to/workspace", nil)
						return m
					}(),
					makePublicDir: func() *MockPublicDirMaker {
						m := NewMockPublicDirMaker(c)
						m.EXPECT().MakePublicDir("/path/to/workspace/abc123/1").Return(nil)
						return m
					}(),
				}
			},
		},
		{
			name: "MakePublicDirエラー",
			args: &args{
				contest: model.ContestId("abc123"),
				problem: model.ProblemId("1"),
			},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockGetWd {
						m := NewMockGetWd(c)
						m.EXPECT().GetWd().Return("/path/to/workspace", nil)
						return m
					}(),
					makePublicDir: func() *MockPublicDirMaker {
						m := NewMockPublicDirMaker(c)
						m.EXPECT().
							MakePublicDir("/path/to/workspace/abc123/1").
							Return(errors.New("make public dir error"))
						return m
					}(),
				}
			},
			wantErr: true,
		},
		{
			name: "GetWdエラー",
			args: &args{
				contest: model.ContestId("abc123"),
				problem: model.ProblemId("1"),
			},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockGetWd {
						m := NewMockGetWd(c)
						m.EXPECT().GetWd().Return("", errors.New("get wd error"))
						return m
					}(),
					makePublicDir: func() *MockPublicDirMaker {
						return NewMockPublicDirMaker(c)
					}(),
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mock := tt.mock(mockCtrl)

			m := problem.NewProblemDirMaker(
				mock.getWd,
				mock.makePublicDir,
			)

			// Act
			err := m.MakeProblemDir(tt.args.contest, tt.args.problem)

			// Asesrt
			if (err != nil) != tt.wantErr {
				t.Errorf("ProblemDirMakerImpl.MakeProblemDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
