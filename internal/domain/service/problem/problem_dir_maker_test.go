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
		getWd            string
		getWdErr         error
		makePublicDirErr error
	}
	type want struct {
		createdDirs []string
		err         bool
	}
	tests := []struct {
		name string
		args *args
		mock *mock
		want *want
	}{
		{
			name: "正常系",
			args: &args{
				contest: model.ContestId("abc123"),
				problem: model.ProblemId("1"),
			},
			mock: &mock{
				getWd: "/path/to/workspace",
			},
			want: &want{
				createdDirs: []string{
					"/path/to/workspace/abc123/1",
				},
			},
		},
		{
			name: "MakePublicDirエラー",
			args: &args{
				contest: model.ContestId("abc123"),
				problem: model.ProblemId("1"),
			},
			mock: &mock{
				getWd:            "/path/to/workspace",
				makePublicDirErr: errors.New("make public dir error"),
			},
			want: &want{err: true},
		},
		{
			name: "GetWdエラー",
			args: &args{
				contest: model.ContestId("abc123"),
				problem: model.ProblemId("1"),
			},
			mock: &mock{
				getWdErr: errors.New("get wd error"),
			},
			want: &want{err: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			var gotDirs map[string]bool

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			getWd := NewMockGetWd(mockCtrl)
			getWd.EXPECT().GetWd().Return(tt.mock.getWd, tt.mock.getWdErr)
			dirMaker := NewMockPublicDirMaker(mockCtrl)
			dirMaker.
				EXPECT().
				MakePublicDir(gomock.Any()).
				AnyTimes().
				Do(func(path string) { gotDirs[path] = true }).
				Return(tt.mock.makePublicDirErr)

			m := problem.NewProblemDirMakerImpl(getWd, dirMaker)

			// Act
			err := m.MakeProblemDir(tt.args.contest, tt.args.problem)

			// Asesrt
			if tt.want.err {
				if err == nil {
					t.Error("MakeProblemDir() error is expected but got nil")
					return
				}
				return
			}

			for _, dir := range tt.want.createdDirs {
				if !gotDirs[dir] {
					t.Errorf("MakeProblemDir() did not create directory %s", dir)
				}
			}
		})
	}
}
