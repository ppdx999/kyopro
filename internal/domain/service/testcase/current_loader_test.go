package testcase_test

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/testcase"
)

func Test_testCaseCurrentLoader_LoadCurrentTestCases(t *testing.T) {
	type mock struct {
		wdGetter             *MockWdGetter
		childFileNamesGetter *MockChildFileNamesGetter
		publicFileReader     *MockPublicFileReader
	}
	tests := []struct {
		name    string
		mock    func(c *gomock.Controller) *mock
		want    []*model.TestCase
		wantErr bool
	}{
		{
			name: "success",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					wdGetter: func() *MockWdGetter {
						m := NewMockWdGetter(c)
						m.EXPECT().GetWd().Return("/abc100/b", nil)
						return m
					}(),
					childFileNamesGetter: func() *MockChildFileNamesGetter {
						m := NewMockChildFileNamesGetter(c)
						m.EXPECT().ChildFileNames("/abc100/b/test").Return([]string{"1.in", "1.out"}, nil)
						return m
					}(),
					publicFileReader: func() *MockPublicFileReader {
						m := NewMockPublicFileReader(c)
						m.EXPECT().ReadPublicFile("/abc100/b/test/1.in").Return([]byte("input 1.in"), nil)
						m.EXPECT().ReadPublicFile("/abc100/b/test/1.out").Return([]byte("output 1.out"), nil)
						return m
					}(),
				}
			},
			want: []*model.TestCase{
				{
					ID:    "1",
					Input: []byte("input 1.in"),
					Want:  []byte("output 1.out"),
				},
			},
		},
		// TODO: Error Cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := tt.mock(ctrl)

			tr := testcase.NewTestCaseCurrentLoader(
				mock.wdGetter,
				mock.childFileNamesGetter,
				mock.publicFileReader,
			)

			// Act
			got, err := tr.LoadCurrentTestCases()

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("testCaseCurrentLoader.LoadCurrentTestCases() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testCaseCurrentLoader.LoadCurrentTestCases() = %v, want %v", got, tt.want)
			}
		})
	}
}
