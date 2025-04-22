package session_test

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/session"
)

type MockMakePublicDir struct {
	calledWithPath string
	err            error
}

func (m *MockMakePublicDir) MakePublicDir(path string) error {
	m.calledWithPath = path
	return m.err
}

type MockWriteSecretFile struct {
	calledWithPath string
	calledWithData []byte
	err            error
}

func (m *MockWriteSecretFile) WriteSecretFile(path string, data []byte) error {
	m.calledWithPath = path
	m.calledWithData = data
	return m.err
}

func TestSessionSaver(t *testing.T) {
	tests := []struct {
		name                string
		sessionPath         string
		sessionPathErr      error
		makePublicDirErr    error
		writeSecretFileErr  error
		inputSession        model.SessionSecret
		wantMakePublicDir   string
		wantWriteSecretPath string
		wantWriteSecretData []byte
		wantErr             bool
	}{
		{
			name:                "正常系",
			sessionPath:         "/home/user/.kyopro/session",
			inputSession:        model.SessionSecret("mysecret"),
			wantMakePublicDir:   "/home/user/.kyopro",
			wantWriteSecretPath: "/home/user/.kyopro/session",
			wantWriteSecretData: []byte("mysecret"),
			wantErr:             false,
		},
		{
			name:           "SessionPathでエラー",
			sessionPathErr: errors.New("session path error"),
			inputSession:   model.SessionSecret("mysecret"),
			wantErr:        true,
		},
		{
			name:             "MakePublicDirでエラー",
			sessionPath:      "/home/user/.kyopro/session",
			makePublicDirErr: errors.New("make dir error"),
			inputSession:     model.SessionSecret("mysecret"),
			wantErr:          true,
		},
		{
			name:               "WriteSecretFileでエラー",
			sessionPath:        "/home/user/.kyopro/session",
			writeSecretFileErr: errors.New("write file error"),
			inputSession:       model.SessionSecret("mysecret"),
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSessionPath := &MockSessionPath{
				path: tt.sessionPath,
				err:  tt.sessionPathErr,
			}
			mockMakePublicDir := &MockMakePublicDir{
				err: tt.makePublicDirErr,
			}
			mockWriteSecretFile := &MockWriteSecretFile{
				err: tt.writeSecretFileErr,
			}

			s := session.NewSessionSaverImpl(
				mockSessionPath,
				mockMakePublicDir,
				mockWriteSecretFile,
			)

			err := s.SaveSession(tt.inputSession)

			if (err != nil) != tt.wantErr {
				t.Errorf("SessionSaver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If no error was expected, check if mocks were called correctly
			if !tt.wantErr {
				// Check MakePublicDir call
				expectedDir := filepath.Dir(tt.sessionPath)
				if mockMakePublicDir.calledWithPath != expectedDir {
					t.Errorf("MakePublicDir called with path = %q, want %q", mockMakePublicDir.calledWithPath, expectedDir)
				}

				// Check WriteSecretFile call
				if mockWriteSecretFile.calledWithPath != tt.sessionPath {
					t.Errorf("WriteSecretFile called with path = %q, want %q", mockWriteSecretFile.calledWithPath, tt.sessionPath)
				}
				if string(mockWriteSecretFile.calledWithData) != string(tt.inputSession) {
					t.Errorf("WriteSecretFile called with data = %q, want %q", string(mockWriteSecretFile.calledWithData), string(tt.inputSession))
				}
			}
		})
	}
}
