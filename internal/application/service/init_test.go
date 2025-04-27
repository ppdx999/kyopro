package application_service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	application_service "github.com/ppdx999/kyopro/internal/application/service"
	"github.com/ppdx999/kyopro/internal/domain/model"
	problem_mock "github.com/ppdx999/kyopro/internal/domain/service/problem/mock"
)

func Test_initer_Init(t *testing.T) {
	type mocks struct {
		problemIdsGetter *problem_mock.MockProblemIdsGetter
		problemDirMaker  *problem_mock.MockProblemDirMaker
	}
	type args struct {
		c model.ContestId
	}
	tests := []struct {
		name    string
		args    *args
		mocks   func(c *gomock.Controller) *mocks
		wantErr bool
	}{
		{
			name: "正常系",
			args: &args{c: model.ContestId("abc100")},
			mocks: func(c *gomock.Controller) *mocks {
				return &mocks{
					problemIdsGetter: func() *problem_mock.MockProblemIdsGetter {
						m := problem_mock.NewMockProblemIdsGetter(c)
						m.EXPECT().GetProblemIds(model.ContestId("abc100")).Return([]model.ProblemId{"a", "b", "c"}, nil)
						return m
					}(),
					problemDirMaker: func() *problem_mock.MockProblemDirMaker {
						m := problem_mock.NewMockProblemDirMaker(c)
						m.EXPECT().MakeProblemDir(model.ContestId("abc100"), model.ProblemId("a")).Return(nil)
						m.EXPECT().MakeProblemDir(model.ContestId("abc100"), model.ProblemId("b")).Return(nil)
						m.EXPECT().MakeProblemDir(model.ContestId("abc100"), model.ProblemId("c")).Return(nil)
						return m
					}(),
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mocks := tt.mocks(mockCtrl)

			s := application_service.NewIniter(
				mocks.problemIdsGetter,
				mocks.problemDirMaker,
			)

			// Act & Assert
			if err := s.Init(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("initer.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
