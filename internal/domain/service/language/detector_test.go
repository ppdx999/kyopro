package language_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/language"
)

func Test_detector_DetectLanguage(t *testing.T) {
	type mock struct {
		allLanguagesFetcher language.AllLanguagesFetcher
		fileExister         language.FileExister
	}
	tests := []struct {
		name    string
		mock    func(c *gomock.Controller) *mock
		want    *model.Language
		wantErr bool
	}{
		{
			name: "success",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					allLanguagesFetcher: func() language.AllLanguagesFetcher {
						m := NewMockAllLanguagesFetcher(c)
						m.EXPECT().FetchAllLanguages().Return([]*model.Language{
							{Name: "Go", MainFile: "main.go"},
							{Name: "Python", MainFile: "main.py"},
						}, nil)
						return m
					}(),
					fileExister: func() language.FileExister {
						m := NewMockFileExister(c)
						m.EXPECT().ExistFile("main.go").Return(true)
						return m
					}(),
				}
			},
			want: &model.Language{
				Name:     "Go",
				MainFile: "main.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := tt.mock(ctrl)

			d := language.NewLanguageDetector(
				mock.allLanguagesFetcher,
				mock.fileExister,
			)

			// Act
			got, err := d.DetectLanguage()

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("detector.DetectLanguage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want.Name != got.Name {
				t.Errorf("detector.DetectLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}
