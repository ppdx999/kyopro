package testcase_test

import (
	"fmt"
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
		getWd              string
		getWdErr           error
		makePublicDirErr   error
		writePublicFileErr error
	}
	type want struct {
		writtenFiles []string
		err          bool
	}

	tests := []struct {
		name string
		args *args
		mock *mock
		want *want
	}{
		{
			name: "success",
			args: &args{tc: &model.TestCase{
				ID:    "1",
				Input: []byte("input 1.in"),
				Want:  []byte("output 1.out"),
			}},
			mock: &mock{
				getWd: "/tmp",
			},
		},
		{
			name: "getWd error",
			args: &args{tc: &model.TestCase{ID: "1"}},
			mock: &mock{getWdErr: fmt.Errorf("getWd error")},
			want: &want{err: true},
		},
		{
			name: "makePublicDir error",
			args: &args{tc: &model.TestCase{ID: "1"}},
			mock: &mock{
				getWd:            "/tmp",
				makePublicDirErr: fmt.Errorf("makePublicDir error"),
			},
			want: &want{err: true},
		},
		{
			name: "writePublicFile error",
			args: &args{tc: &model.TestCase{ID: "1"}},
			mock: &mock{
				getWd:              "/tmp",
				writePublicFileErr: fmt.Errorf("writePublicFile error"),
			},
			want: &want{err: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			var writtenFiles map[string]bool
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockGetWd := NewMockGetWd(mockCtrl)
			mockGetWd.EXPECT().GetWd().Return(tt.mock.getWd, tt.mock.getWdErr)
			mockDirMaker := NewMockPublicDirMaker(mockCtrl)
			mockDirMaker.EXPECT().MakePublicDir(gomock.Any()).Return(tt.mock.makePublicDirErr)
			mockPublicFileWriter := NewMockPublicFileWriter(mockCtrl)
			mockPublicFileWriter.EXPECT().
				WritePublicFile(gomock.Any(), gomock.Any()).AnyTimes().
				Do(func(path string, data []byte) {
					writtenFiles[path] = true
				}).Return(tt.mock.writePublicFileErr)

			saver := testcase.NewTestCaseFsSaver(mockGetWd, mockDirMaker, mockPublicFileWriter)

			// Act
			err := saver.SaveTestCase(tt.args.tc)

			// Assert
			if (err != nil) != tt.want.err {
				t.Errorf("TestCaseFsSaver.Save() error = %v, wantErr %v", err, tt.want.err)
			}

			for _, path := range tt.want.writtenFiles {
				if !writtenFiles[path] {
					t.Errorf("TestCaseFsSaver.Save() did not write file %s", path)
				}
			}
		})
	}
}
