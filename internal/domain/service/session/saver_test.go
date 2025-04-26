package session_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/session"
)

func TestSessionSaver(t *testing.T) {
	type args struct {
		session model.SessionSecret
	}
	type mock struct {
		sessionPath         string
		sessionPathErr      error
		makePublicDirArg    string
		makePublicDirErr    error
		writeSecretFilePath string
		writeSecretFileData []byte
		writeSecretFileErr  error
	}
	tests := []struct {
		name    string
		args    *args
		mock    *mock
		wantErr bool
	}{
		{
			name: "正常系",
			args: &args{session: model.SessionSecret("mysecret")},
			mock: &mock{
				sessionPath:         "/home/user/.kyopro/session",
				makePublicDirArg:    "/home/user/.kyopro",
				writeSecretFilePath: "/home/user/.kyopro/session",
				writeSecretFileData: []byte("mysecret"),
			},
		},
		{
			name:    "SessionPathでエラー",
			args:    &args{session: model.SessionSecret("mysecret")},
			mock:    &mock{sessionPathErr: errors.New("session path error")},
			wantErr: true,
		},
		{
			name: "MakePublicDirでエラー",
			args: &args{session: model.SessionSecret("mysecret")},
			mock: &mock{
				sessionPath:      "/home/user/.kyopro/session",
				makePublicDirArg: "/home/user/.kyopro",
				makePublicDirErr: errors.New("make public dir error"),
			},
			wantErr: true,
		},
		{
			name: "WriteSecretFileでエラー",
			args: &args{session: model.SessionSecret("mysecret")},
			mock: &mock{
				sessionPath:         "/home/user/.kyopro/session",
				makePublicDirArg:    "/home/user/.kyopro",
				writeSecretFilePath: "/home/user/.kyopro/session",
				writeSecretFileData: []byte("mysecret"),
				writeSecretFileErr:  errors.New("write secret file error"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockSessionPath := NewMockSessionPath(mockCtrl)
			mockSessionPath.
				EXPECT().
				SessionPath().
				Return(tt.mock.sessionPath, tt.mock.sessionPathErr)

			mockMakePublicDir := NewMockMakePublicDir(mockCtrl)
			mockMakePublicDir.
				EXPECT().
				MakePublicDir(tt.mock.makePublicDirArg).
				Return(tt.mock.makePublicDirErr)

			mockWriteSecretFile := NewMockWriteSecretFile(mockCtrl)
			mockWriteSecretFile.
				EXPECT().
				WriteSecretFile(tt.mock.writeSecretFilePath, tt.mock.writeSecretFileData).
				Return(tt.mock.writeSecretFileErr)

			s := session.NewSessionSaverImpl(
				mockSessionPath,
				mockMakePublicDir,
				mockWriteSecretFile,
			)

			// Act
			err := s.SaveSession(tt.args.session)

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionSaver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
