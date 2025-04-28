package language_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	model_mock "github.com/ppdx999/kyopro/internal/domain/model/mock"
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
		want    func(c *gomock.Controller) model.Language
		wantErr bool
	}{
		{
			name: "success",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					allLanguagesFetcher: func() language.AllLanguagesFetcher {
						m := NewMockAllLanguagesFetcher(c)
						m.EXPECT().FetchAllLanguages().Return([]model.Language{
							func() model.Language {
								// Python
								m := model_mock.NewMockLanguage(c)
								m.EXPECT().SourceCode().Return(&model.SourceCode{
									Path: "main.py",
								})
								return m
							}(),
							func() model.Language {
								// Golang
								m := model_mock.NewMockLanguage(c)
								m.EXPECT().Name().Return("Go")
								m.EXPECT().SourceCode().Return(&model.SourceCode{
									Path: "main.go",
								})
								return m
							}(),
						}, nil)
						return m
					}(),
					fileExister: func() language.FileExister {
						m := NewMockFileExister(c)
						m.EXPECT().ExistFile("main.py").Return(false)
						m.EXPECT().ExistFile("main.go").Return(true)
						return m
					}(),
				}
			},
			want: func(c *gomock.Controller) model.Language {
				m := model_mock.NewMockLanguage(c)
				m.EXPECT().Name().Return("Go")
				return m
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

			want := tt.want(ctrl)
			if want.Name() != got.Name() {
				t.Errorf("detector.DetectLanguage() = %v, want %v", got.Name(), want.Name())
			}
		})
	}
}
