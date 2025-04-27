package cli_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	user_mock "github.com/ppdx999/kyopro/internal/domain/service/user/mock"
	"github.com/ppdx999/kyopro/internal/presentation/cli"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		want      cli.ExitCode
		calledCmd string
	}{
		{
			name:      "Correct Case Call sub1",
			args:      []string{"sub1"},
			want:      cli.ExitOK,
			calledCmd: "sub1",
		},
		{
			name:      "Correct Case Call sub2",
			args:      []string{"sub2"},
			want:      cli.ExitOK,
			calledCmd: "sub2",
		},
		{
			name: "No subcommand provided",
			args: nil,
			want: cli.ExitErr,
		},
		{
			name: "Unknown subcommand",
			args: []string{"unknown"},
			want: cli.ExitErr,
		},
		{
			name: "Help flag",
			args: []string{"-h"},
			want: cli.ExitErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			var calledCmd string
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			sub1 := NewMockCmd(mockCtrl)
			sub1.EXPECT().Name().Return("sub1")
			sub1.EXPECT().Description().Return("this is sub1 description")
			sub1.EXPECT().Run(gomock.Any()).DoAndReturn(func(args []string) cli.ExitCode {
				calledCmd = "sub1"
				return cli.ExitOK
			})

			sub2 := NewMockCmd(mockCtrl)
			sub2.EXPECT().Name().Return("sub2")
			sub2.EXPECT().Description().Return("this is sub2 description")
			sub2.EXPECT().Run(gomock.Any()).DoAndReturn(func(args []string) cli.ExitCode {
				calledCmd = "sub2"
				return cli.ExitOK
			})

			msgSender := user_mock.NewMockMsgSender(mockCtrl)
			msgSender.EXPECT().SendMsg(gomock.Any())

			root := cli.NewDispatcher(msgSender)
			root.Register(sub1)
			root.Register(sub2)

			// Act
			got := root.Run(tt.args)

			if got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
				return
			}

			if calledCmd != tt.calledCmd {
				t.Errorf("%s was not called", tt.calledCmd)
			}
		})
	}
}
