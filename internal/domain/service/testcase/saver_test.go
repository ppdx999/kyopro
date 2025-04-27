package testcase_test

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/testcase"
)

func TestSaveTestCase(t *testing.T) {
	type args struct {
		tc *model.TestCase
	}
	type mock struct {
		getWd            *MockWdGetter
		publicDirMaker   *MockPublicDirMaker
		publicFileWriter *MockPublicFileWriter
	}

	tests := []struct {
		name    string
		args    *args
		mock    func(c *gomock.Controller) *mock
		wantErr bool
	}{
		{
			name: "success",
			args: &args{tc: &model.TestCase{
				ID:    "1",
				Input: []byte("input 1.in"),
				Want:  []byte("output 1.out"),
			}},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockWdGetter {
						m := NewMockWdGetter(c)
						m.EXPECT().GetWd().Return("/tmp", nil)
						return m
					}(),
					publicDirMaker: func() *MockPublicDirMaker {
						m := NewMockPublicDirMaker(c)
						m.EXPECT().MakePublicDir("/tmp/test").Return(nil)
						return m
					}(),
					publicFileWriter: func() *MockPublicFileWriter {
						m := NewMockPublicFileWriter(c)
						m.EXPECT().WritePublicFile("/tmp/test/1.in", []byte("input 1.in")).Return(nil)
						m.EXPECT().WritePublicFile("/tmp/test/1.out", []byte("output 1.out")).Return(nil)
						return m
					}(),
				}
			},
		},
		{
			name: "getWd error",
			args: &args{tc: &model.TestCase{
				ID:    "1",
				Input: []byte("input 1.in"),
				Want:  []byte("output 1.out"),
			}},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockWdGetter {
						m := NewMockWdGetter(c)
						m.EXPECT().GetWd().Return("", errors.New("getWd error"))
						return m
					}(),
					publicDirMaker: func() *MockPublicDirMaker {
						m := NewMockPublicDirMaker(c)
						return m
					}(),
					publicFileWriter: func() *MockPublicFileWriter {
						m := NewMockPublicFileWriter(c)
						return m
					}(),
				}
			},
			wantErr: true,
		},
		{
			name: "makePublicDir error",
			args: &args{tc: &model.TestCase{
				ID:    "1",
				Input: []byte("input 1.in"),
				Want:  []byte("output 1.out"),
			}},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockWdGetter {
						m := NewMockWdGetter(c)
						m.EXPECT().GetWd().Return("/tmp", nil)
						return m
					}(),
					publicDirMaker: func() *MockPublicDirMaker {
						m := NewMockPublicDirMaker(c)
						m.EXPECT().MakePublicDir("/tmp/test").Return(errors.New("makePublicDir error"))
						return m
					}(),
					publicFileWriter: func() *MockPublicFileWriter {
						m := NewMockPublicFileWriter(c)
						return m
					}(),
				}
			},
			wantErr: true,
		},
		{
			name: "writePublicFile error",
			args: &args{tc: &model.TestCase{
				ID:    "1",
				Input: []byte("input 1.in"),
				Want:  []byte("output 1.out"),
			}},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					getWd: func() *MockWdGetter {
						m := NewMockWdGetter(c)
						m.EXPECT().GetWd().Return("/tmp", nil)
						return m
					}(),
					publicDirMaker: func() *MockPublicDirMaker {
						m := NewMockPublicDirMaker(c)
						m.EXPECT().MakePublicDir("/tmp/test").Return(nil)
						return m
					}(),
					publicFileWriter: func() *MockPublicFileWriter {
						m := NewMockPublicFileWriter(c)
						m.EXPECT().WritePublicFile("/tmp/test/1.in", []byte("input 1.in")).Return(errors.New("writePublicFile error"))
						m.EXPECT().WritePublicFile("/tmp/test/1.out", []byte("output 1.out")).Return(nil)
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
			mock := tt.mock(mockCtrl)

			saver := testcase.NewTestCaseSaver(
				mock.getWd,
				mock.publicDirMaker,
				mock.publicFileWriter,
			)

			// Act
			err := saver.SaveTestCase(tt.args.tc)

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("TestCaseFsSaver.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
