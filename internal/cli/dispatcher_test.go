package cli_test

import (
	"testing"

	"github.com/ppdx999/kyopro/internal/cli"
	"github.com/ppdx999/kyopro/internal/testutil"
)

type MockCmd struct {
	called bool
}

func (m *MockCmd) Run(args []string) cli.ExitCode {
	m.called = true
	return cli.ExitOK
}

func TestRun(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		sendMsgs  []string
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
			sendMsgs: []string{"no subcommand provided"},
			exitCode: cli.ExitErr,
		},
		{
			name:     "Unknown subcommand",
			args:     []string{"unknown"},
			sendMsgs: []string{"unknown subcommand: unknown"},
			exitCode: cli.ExitErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sub1 := &MockCmd{}
			sub2 := &MockCmd{}
			msgSender := &testutil.MockMsgSender{}
			root := cli.NewDispatcher(msgSender)
			root.Register("sub1", sub1)
			root.Register("sub2", sub2)

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

			if tt.sendMsgs != nil && !msgSender.CheckMsgs(tt.sendMsgs) {
				t.Errorf("SendMsg() = %v, want %v", msgSender.GetSendMsgs(), tt.sendMsgs)
			}
		})
	}
}
