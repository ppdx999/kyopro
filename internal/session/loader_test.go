package session_test

import (
	"errors"
	"testing"

	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/session"
)

type MockReadSecretFile struct {
	calledWithPath string
	data           []byte
	err            error
}

func (m *MockReadSecretFile) ReadSecretFile(path string) ([]byte, error) {
	m.calledWithPath = path
	return m.data, m.err
}

type MockExistFile struct {
	exist bool
}

func (m *MockExistFile) ExistFile(path string) bool {
	return m.exist
}

func TestSessionLoader(t *testing.T) {
	tests := []struct {
		name              string
		sessionPath       string
		sessionPathErr    error
		existFile         bool
		readSecretFile    []byte
		readSecretFileErr error
		want              model.SessionSecret
		wantErr           bool
	}{
		{
			name:           "正常系",
			sessionPath:    "/home/user/.kyopro/session",
			existFile:      true,
			readSecretFile: []byte("mysecret"),
			want:           model.SessionSecret("mysecret"),
		},
		{
			name:           "sessionPathがエラー",
			existFile:      true,
			sessionPathErr: errors.New("session path error"),
			wantErr:        true,
		},
		{
			name:              "readSecretFileがエラー",
			existFile:         true,
			sessionPath:       "/home/user/.kyopro/session",
			readSecretFileErr: errors.New("read file error"),
			wantErr:           true,
		},
		{
			name:           "sessionが空",
			existFile:      true,
			sessionPath:    "/home/user/.kyopro/session",
			readSecretFile: []byte(""),
			want:           model.SessionSecret(""),
		},
		{
			name:        "sessionファイルが存在しない",
			existFile:   false,
			sessionPath: "/home/user/.kyopro/session",
			want:        model.SessionSecret(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSessionPath := &MockSessionPath{
				path: tt.sessionPath,
				err:  tt.sessionPathErr,
			}
			mockReadSecretFile := &MockReadSecretFile{
				data: tt.readSecretFile,
				err:  tt.readSecretFileErr,
			}
			mockExistFile := &MockExistFile{
				exist: tt.existFile,
			}

			s := session.NewSessionLoaderImpl(
				mockSessionPath,
				mockExistFile,
				mockReadSecretFile,
			)

			got, err := s.LoadSession()

			if (err != nil) != tt.wantErr {
				t.Errorf("SessionLoader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("SessionLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}
