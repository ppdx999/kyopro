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
		sessionPath    *MockSessionPath
		existFile      *MockExistFile
		readSecretFile *MockReadSecretFile
	}
	tests := []struct {
		name    string
		mock    func(c *gomock.Controller) *mock
		want    model.SessionSecret
		wantErr bool
	}{
		{
			name: "正常系",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sessionPath: func() *MockSessionPath {
						m := NewMockSessionPath(c)
						m.EXPECT().SessionPath().Return("/home/user/.kyopro/session", nil)
						return m
					}(),
					existFile: func() *MockExistFile {
						m := NewMockExistFile(c)
						m.EXPECT().ExistFile("/home/user/.kyopro/session").Return(true)
						return m
					}(),
					readSecretFile: func() *MockReadSecretFile {
						m := NewMockReadSecretFile(c)
						m.EXPECT().ReadSecretFile("/home/user/.kyopro/session").Return([]byte("mysecret"), nil)
						return m
					}(),
				}
			},
			want: model.SessionSecret("mysecret"),
		},
		{
			name: "sessionPathがエラー",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sessionPath: func() *MockSessionPath {
						m := NewMockSessionPath(c)
						m.EXPECT().SessionPath().Return("", errors.New("session path error"))
						return m
					}(),
				}
			},
			wantErr: true,
		},
		{
			name: "readSecretFileがエラー",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sessionPath: func() *MockSessionPath {
						m := NewMockSessionPath(c)
						m.EXPECT().SessionPath().Return("/home/user/.kyopro/session", nil)
						return m
					}(),
					existFile: func() *MockExistFile {
						m := NewMockExistFile(c)
						m.EXPECT().ExistFile("/home/user/.kyopro/session").Return(true)
						return m
					}(),
					readSecretFile: func() *MockReadSecretFile {
						m := NewMockReadSecretFile(c)
						m.EXPECT().ReadSecretFile("/home/user/.kyopro/session").Return(nil, errors.New("read secret file error"))
						return m
					}(),
				}
			},
			wantErr: true,
		},
		{
			name: "sessionが空",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sessionPath: func() *MockSessionPath {
						m := NewMockSessionPath(c)
						m.EXPECT().SessionPath().Return("/home/user/.kyopro/session", nil)
						return m
					}(),
					existFile: func() *MockExistFile {
						m := NewMockExistFile(c)
						m.EXPECT().ExistFile("/home/user/.kyopro/session").Return(true)
						return m
					}(),
					readSecretFile: func() *MockReadSecretFile {
						m := NewMockReadSecretFile(c)
						m.EXPECT().ReadSecretFile("/home/user/.kyopro/session").Return([]byte(""), nil)
						return m
					}(),
				}
			},
			want: model.SessionSecret(""),
		},
		{
			name: "sessionファイルが存在しない",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sessionPath: func() *MockSessionPath {
						m := NewMockSessionPath(c)
						m.EXPECT().SessionPath().Return("/home/user/.kyopro/session", nil)
						return m
					}(),
					existFile: func() *MockExistFile {
						m := NewMockExistFile(c)
						m.EXPECT().ExistFile("/home/user/.kyopro/session").Return(false)
						return m
					}(),
				}
			},
			want: model.SessionSecret(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mock := tt.mock(mockCtrl)

			s := session.NewSessionLoader(
				mock.sessionPath,
				mock.existFile,
				mock.readSecretFile,
			)

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
