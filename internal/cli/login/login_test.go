package login_test

import (
	"errors"
	"testing"

	"github.com/ppdx999/kyopro/internal/cli"
	"github.com/ppdx999/kyopro/internal/cli/login"
)

type MockMsgSender struct {
	sendMsgs []string
}

func (m *MockMsgSender) SendMsg(msg string) {
	m.sendMsgs = append(m.sendMsgs, msg)
}

type MockLoginService struct {
	err error
}

func (m *MockLoginService) Login() error {
	return m.err
}

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		loginErr error
		exitCode cli.ExitCode
	}{
		{
			name:     "success",
			args:     []string{},
			exitCode: cli.ExitOK,
		},
		{
			name:     "invalid args",
			args:     []string{"invalidarg"},
			exitCode: cli.ExitErr,
		},
		{
			name:     "help",
			args:     []string{"--help"},
			exitCode: cli.ExitErr,
		},
		{
			name:     "login error",
			args:     []string{},
			loginErr: errors.New("login error"),
			exitCode: cli.ExitErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockLoginService := &MockLoginService{
				err: tt.loginErr,
			}
			mockMsgSender := &MockMsgSender{}
			cmd := login.NewLoginCli(mockLoginService, mockMsgSender)

			// Act
			exitCode := cmd.Run(tt.args)

			// Assert
			if exitCode != tt.exitCode {
				t.Errorf("Run() = %v, want %v", exitCode, tt.exitCode)
			}
		})
	}
}
