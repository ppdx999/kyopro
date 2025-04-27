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
		sessionPath     *MockSessionPath
		makePublicDir   *MockMakePublicDir
		writeSecretFile *MockWriteSecretFile
	}
	tests := []struct {
		name    string
		args    *args
		mock    func(c *gomock.Controller) *mock
		wantErr bool
	}{
		{
			name: "正常系",
			args: &args{session: model.SessionSecret("mysecret")},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sessionPath: func() *MockSessionPath {
						m := NewMockSessionPath(c)
						m.EXPECT().SessionPath().Return("/home/user/.kyopro/session", nil)
						return m
					}(),
					makePublicDir: func() *MockMakePublicDir {
						m := NewMockMakePublicDir(c)
						m.EXPECT().MakePublicDir("/home/user/.kyopro").Return(nil)
						return m
					}(),
					writeSecretFile: func() *MockWriteSecretFile {
						m := NewMockWriteSecretFile(c)
						m.EXPECT().WriteSecretFile("/home/user/.kyopro/session", []byte("mysecret")).Return(nil)
						return m
					}(),
				}
			},
		},
		{
			name: "SessionPathでエラー",
			args: &args{session: model.SessionSecret("mysecret")},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sessionPath: func() *MockSessionPath {
						m := NewMockSessionPath(c)
						m.EXPECT().SessionPath().Return("", errors.New("session path error"))
						return m
					}(),
					makePublicDir: func() *MockMakePublicDir {
						m := NewMockMakePublicDir(c)
						return m
					}(),
					writeSecretFile: func() *MockWriteSecretFile {
						m := NewMockWriteSecretFile(c)
						return m
					}(),
				}
			},
			wantErr: true,
		},
		{
			name: "MakePublicDirでエラー",
			args: &args{session: model.SessionSecret("mysecret")},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sessionPath: func() *MockSessionPath {
						m := NewMockSessionPath(c)
						m.EXPECT().SessionPath().Return("/home/user/.kyopro/session", nil)
						return m
					}(),
					makePublicDir: func() *MockMakePublicDir {
						m := NewMockMakePublicDir(c)
						m.EXPECT().MakePublicDir("/home/user/.kyopro").Return(errors.New("make public dir error"))
						return m
					}(),
					writeSecretFile: func() *MockWriteSecretFile {
						m := NewMockWriteSecretFile(c)
						return m
					}(),
				}
			},
			wantErr: true,
		},
		{
			name: "WriteSecretFileでエラー",
			args: &args{session: model.SessionSecret("mysecret")},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sessionPath: func() *MockSessionPath {
						m := NewMockSessionPath(c)
						m.EXPECT().SessionPath().Return("/home/user/.kyopro/session", nil)
						return m
					}(),
					makePublicDir: func() *MockMakePublicDir {
						m := NewMockMakePublicDir(c)
						m.EXPECT().MakePublicDir("/home/user/.kyopro").Return(nil)
						return m
					}(),
					writeSecretFile: func() *MockWriteSecretFile {
						m := NewMockWriteSecretFile(c)
						m.EXPECT().WriteSecretFile("/home/user/.kyopro/session", []byte("mysecret")).Return(errors.New("write secret file error"))
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

			s := session.NewSessionSaver(
				mock.sessionPath,
				mock.makePublicDir,
				mock.writeSecretFile,
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
