package download_test

import (
	"errors"
	"testing"

	"github.com/ppdx999/kyopro/internal/cli"
	"github.com/ppdx999/kyopro/internal/cli/download"

	// login_test.go で使われている testutil を利用します
	// パスが異なる場合は適宜修正してください
	"github.com/ppdx999/kyopro/internal/testutil"
)

// Mock for download.DownloadService
type MockDownloadService struct {
	err error
}

func (m *MockDownloadService) Download() error {
	return m.err
}

func TestRun(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		downloadErr  error
		wantExitCode cli.ExitCode
	}{
		{
			name:         "正常系",
			args:         []string{},
			downloadErr:  nil,
			wantExitCode: cli.ExitOK,
		},
		{
			name:         "引数エラー",
			args:         []string{"invalidarg"}, // downloadコマンドは引数を取らない
			downloadErr:  nil,
			wantExitCode: cli.ExitErr,
		},
		{
			name:         "Downloadサービスエラー",
			args:         []string{},
			downloadErr:  errors.New("download service error"),
			wantExitCode: cli.ExitErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockDownloadService := &MockDownloadService{
				err: tt.downloadErr,
			}
			mockMsgSender := &testutil.MockMsgSender{}
			cmd := download.NewDownloadCmd(mockDownloadService, mockMsgSender)

			// Act
			exitCode := cmd.Run(tt.args)

			// Assert
			if exitCode != tt.wantExitCode {
				t.Errorf("Run() exitCode = %v, want %v", exitCode, tt.wantExitCode)
			}
		})
	}
}
