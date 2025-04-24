package cli_test

import (
	"testing"

	"github.com/ppdx999/kyopro/internal/presentation/cli"
	"github.com/ppdx999/kyopro/internal/testutil"
)

type MockCmd struct {
	name        string
	description string
	called      bool
}

func (m *MockCmd) Name() string {
	return m.name
}

func (m *MockCmd) Description() string {
	return m.description
}

func (m *MockCmd) Run(args []string) cli.ExitCode {
	m.called = true
	return cli.ExitOK
}

func TestRun(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		exitCode  cli.ExitCode
		calledCmd string
	}{
		{
			name:      "Correct Case Call sub1",
			args:      []string{"sub1"},
			exitCode:  cli.ExitOK,
			calledCmd: "sub1",
		},
		{
			name:      "Correct Case Call sub2",
			args:      []string{"sub2"},
			exitCode:  cli.ExitOK,
			calledCmd: "sub2",
		},
		{
			name:     "No subcommand provided",
			args:     nil,
			exitCode: cli.ExitErr,
		},
		{
			name:     "Unknown subcommand",
			args:     []string{"unknown"},
			exitCode: cli.ExitErr,
		},
		{
			name:     "Help flag",
			args:     []string{"-h"},
			exitCode: cli.ExitErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sub1 := &MockCmd{
				name:        "sub1",
				description: "this is sub1 description",
			}
			sub2 := &MockCmd{
				name:        "sub2",
				description: "this is sub2 description",
			}
			msgSender := &testutil.MockMsgSender{}
			root := cli.NewDispatcher(msgSender)
			root.Register(sub1)
			root.Register(sub2)

			gotExitCode := root.Run(tt.args)

			var CheckCalledCmd = func(cmd string) bool {
				switch cmd {
				case "sub1":
					return sub1.called
				case "sub2":
					return sub2.called
				default:
					t.Fatalf("Unknown subcommand: %s", cmd)
					return false
				}
			}

			if gotExitCode != tt.exitCode {
				t.Errorf("Run() = %v, want %v", gotExitCode, tt.exitCode)
				return
			}
			if tt.calledCmd != "" && !CheckCalledCmd(tt.calledCmd) {
				t.Errorf("%s was not called", tt.calledCmd)
			}
		})
	}
}
