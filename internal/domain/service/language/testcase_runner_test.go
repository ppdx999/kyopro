package language_test

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	model_mock "github.com/ppdx999/kyopro/internal/domain/model/mock"
	"github.com/ppdx999/kyopro/internal/domain/service/language"
)

func Test_testcaseRunner_Run(t *testing.T) {
	type args struct {
		l *model.Language
		t *model.TestCase
	}
	tests := []struct {
		name    string
		args    func(c *gomock.Controller) *args
		want    []byte
		want2   []byte
		wantErr bool
	}{
		{
			name: "success",
			args: func(c *gomock.Controller) *args {
				return &args{
					l: &model.Language{
						Name:     "Go",
						MainFile: "main.go",
						Runner: func() model.LanguageRunner {
							m := model_mock.NewMockLanguageRunner(c)
							m.EXPECT().Run("main.go", gomock.Any()).DoAndReturn(
								func(entryFile string, p *model.Pipeline) error {
									p.Outflow.Write([]byte("Output of main.go"))
									p.ErrFlow.Write([]byte("Error of main.go"))
									return nil
								},
							)
							return m
						}(),
					},
					t: &model.TestCase{
						ID:    "1",
						Input: []byte("input 1.in"),
						Want:  []byte("output 1.out"),
					},
				}
			},
			want:  []byte("Output of main.go"),
			want2: []byte("Error of main.go"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			args := tt.args(ctrl)

			r := language.NewTestCaseRunner()

			// Act
			got, got2, err := r.Run(args.l, args.t)

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("testcaseRunner.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testcaseRunner.Run() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("testcaseRunner.Run() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
