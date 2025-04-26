package session_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/session"
)

func TestSessionLoader(t *testing.T) {
	type mock struct {
		sessionPath       string
		sessionPathErr    error
		existFile         bool
		readSecretFile    []byte
		readSecretFileErr error
	}
	tests := []struct {
		name    string
		mock    *mock
		want    model.SessionSecret
		wantErr bool
	}{
		{
			name: "正常系",
			mock: &mock{
				sessionPath:    "/home/user/.kyopro/session",
				existFile:      true,
				readSecretFile: []byte("mysecret"),
			},
			want: model.SessionSecret("mysecret"),
		},
		{
			name:    "sessionPathがエラー",
			mock:    &mock{sessionPathErr: errors.New("session path error")},
			wantErr: true,
		},
		{
			name: "readSecretFileがエラー",
			mock: &mock{
				sessionPath:       "/home/user/.kyopro/session",
				existFile:         true,
				readSecretFileErr: errors.New("read file error"),
			},
			wantErr: true,
		},
		{
			name: "sessionが空",
			mock: &mock{
				sessionPath:    "/home/user/.kyopro/session",
				existFile:      true,
				readSecretFile: []byte(""),
			},
			want: model.SessionSecret(""),
		},
		{
			name: "sessionファイルが存在しない",
			mock: &mock{
				sessionPath: "/home/user/.kyopro/session",
				existFile:   false,
			},
			want: model.SessionSecret(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			sessionPath := NewMockSessionPath(mockCtrl)
			sessionPath.EXPECT().SessionPath().Return(tt.mock.sessionPath, tt.mock.sessionPathErr)
			existFile := NewMockExistFile(mockCtrl)
			existFile.EXPECT().ExistFile(tt.mock.sessionPath).Return(tt.mock.existFile)
			readSecretFile := NewMockReadSecretFile(mockCtrl)
			readSecretFile.EXPECT().ReadSecretFile(tt.mock.sessionPath).Return(tt.mock.readSecretFile, tt.mock)

			s := session.NewSessionLoaderImpl(sessionPath, existFile, readSecretFile)

			// Act
			got, err := s.LoadSession()

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionLoader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}
