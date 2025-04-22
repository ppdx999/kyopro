package init_test

import (
	"errors"
	"testing"

	"github.com/ppdx999/kyopro/internal/cli"
	cli_init "github.com/ppdx999/kyopro/internal/cli/init"
	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/testutil"
)

type mockInitService struct {
	calledWith model.ContestId
	err        error
}

func (m *mockInitService) Init(c model.ContestId) error {
	m.calledWith = c
	return m.err
}

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    model.ContestId
		wantErr bool
	}{
		{
			name:    "valid args",
			args:    []string{"abc123"},
			want:    "abc123",
			wantErr: false,
		},
		{
			name:    "blank args",
			args:    []string{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "too many args",
			args:    []string{"abc123", "def456"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "help flag",
			args:    []string{"--help"},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &mockInitService{}
			msgSender := &testutil.MockMsgSender{}

			initCli := cli_init.NewInitCli(service, msgSender)
			opt, err := initCli.ParseArgs(tt.args)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if opt != nil && opt.ContestId != tt.want {
				t.Errorf("ParseArgs() = %v, want %v", opt.ContestId, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		ucErr    error
		exitCode cli.ExitCode
	}{
		{
			name:     "valid case",
			args:     []string{"abc123"},
			ucErr:    nil,
			exitCode: cli.ExitOK,
		},
		{
			name:     "parseArg error",
			args:     []string{"abc123", "def456"},
			ucErr:    nil,
			exitCode: cli.ExitErr,
		},
		{
			name:     "usecase error",
			args:     []string{"abc123"},
			ucErr:    errors.New("usecase error"),
			exitCode: cli.ExitErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &mockInitService{err: tt.ucErr}
			msgSender := &testutil.MockMsgSender{}

			initCli := cli_init.NewInitCli(service, msgSender)
			got := initCli.Run(tt.args)
			if got != tt.exitCode {
				t.Errorf("Run() = %v, want %v", got, tt.exitCode)
			}
		})
	}
}
