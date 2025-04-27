package application_service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	application_service "github.com/ppdx999/kyopro/internal/application/service"
	"github.com/ppdx999/kyopro/internal/domain/model"
	problem_mock "github.com/ppdx999/kyopro/internal/domain/service/problem/mock"
	testcase_mock "github.com/ppdx999/kyopro/internal/domain/service/testcase/mock"
)

func Test_downloader_Download(t *testing.T) {
	type mocks struct {
		loader *problem_mock.MockCurrentProblemLoader
		getter *testcase_mock.MockTestCasesGetter
		saver  *testcase_mock.MockTestCaseSaver
	}
	tests := []struct {
		name    string
		mocks   func(c *gomock.Controller) *mocks
		wantErr bool
	}{
		{
			name: "正常系",
			mocks: func(c *gomock.Controller) *mocks {
				return &mocks{
					loader: func() *problem_mock.MockCurrentProblemLoader {
						m := problem_mock.NewMockCurrentProblemLoader(c)
						m.EXPECT().LoadCurrentProblem().Return(&model.Problem{
							ID: "a",
							Contest: &model.Contest{
								ID: "abc100",
							},
						}, nil)
						return m
					}(),
					getter: func() *testcase_mock.MockTestCasesGetter {
						m := testcase_mock.NewMockTestCasesGetter(c)
						m.EXPECT().GetTestCases("abc100", "a").Return([]*model.TestCase{
							{
								ID:    "0",
								Input: []byte("input"),
								Want:  []byte("want"),
							},
						}, nil)
						return m
					}(),
					saver: func() *testcase_mock.MockTestCaseSaver {
						m := testcase_mock.NewMockTestCaseSaver(c)
						m.EXPECT().SaveTestCase(&model.TestCase{
							ID:    "0",
							Input: []byte("input"),
							Want:  []byte("want"),
						}).Return(nil)
						return m
					}(),
				}
			},
		},
		{
			name: "currentProblemLoaderでエラー",
			mocks: func(c *gomock.Controller) *mocks {
				return &mocks{
					loader: func() *problem_mock.MockCurrentProblemLoader {
						m := problem_mock.NewMockCurrentProblemLoader(c)
						m.EXPECT().LoadCurrentProblem().Return(nil, errors.New("currentProblemLoader error"))
						return m
					}(),
				}
			},
			wantErr: true,
		},
		{
			name: "testCasesGetterでエラー",
			mocks: func(c *gomock.Controller) *mocks {
				return &mocks{
					loader: func() *problem_mock.MockCurrentProblemLoader {
						m := problem_mock.NewMockCurrentProblemLoader(c)
						m.EXPECT().LoadCurrentProblem().Return(&model.Problem{
							ID: "a",
							Contest: &model.Contest{
								ID: "abc100",
							},
						}, nil)
						return m
					}(),
					getter: func() *testcase_mock.MockTestCasesGetter {
						m := testcase_mock.NewMockTestCasesGetter(c)
						m.EXPECT().GetTestCases("abc100", "a").Return(nil, errors.New("testCasesGetter error"))
						return m
					}(),
				}
			},
			wantErr: true,
		},
		{
			name: "testCaseSaverでエラー",
			mocks: func(c *gomock.Controller) *mocks {
				return &mocks{
					loader: func() *problem_mock.MockCurrentProblemLoader {
						m := problem_mock.NewMockCurrentProblemLoader(c)
						m.EXPECT().LoadCurrentProblem().Return(&model.Problem{
							ID: "a",
							Contest: &model.Contest{
								ID: "abc100",
							},
						}, nil)
						return m
					}(),
					getter: func() *testcase_mock.MockTestCasesGetter {
						m := testcase_mock.NewMockTestCasesGetter(c)
						m.EXPECT().GetTestCases("abc100", "a").Return([]*model.TestCase{
							{
								ID:    "0",
								Input: []byte("input"),
								Want:  []byte("want"),
							},
						}, nil)
						return m
					}(),
					saver: func() *testcase_mock.MockTestCaseSaver {
						m := testcase_mock.NewMockTestCaseSaver(c)
						m.EXPECT().SaveTestCase(&model.TestCase{
							ID:    "0",
							Input: []byte("input"),
							Want:  []byte("want"),
						}).Return(errors.New("testCaseSaver error"))
						return m
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
			mocks := tt.mocks(mockCtrl)

			d := application_service.NewDownloader(
				mocks.loader,
				mocks.getter,
				mocks.saver,
			)

			// Act
			err := d.Download()

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("downloader.Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
