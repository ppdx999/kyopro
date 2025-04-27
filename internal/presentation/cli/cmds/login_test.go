package cmds_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	application_service "github.com/ppdx999/kyopro/internal/application/service"
	application_service_mock "github.com/ppdx999/kyopro/internal/application/service/mock"
	"github.com/ppdx999/kyopro/internal/domain/service/user"
	user_mock "github.com/ppdx999/kyopro/internal/domain/service/user/mock"
	"github.com/ppdx999/kyopro/internal/presentation/cli"
	"github.com/ppdx999/kyopro/internal/presentation/cli/cmds"
)

func LoginCmdTestRun(t *testing.T) {
	var defaultMsgSender = func(c *gomock.Controller) *user_mock.MockMsgSender {
		m := user_mock.NewMockMsgSender(c)
		m.EXPECT().SendMsg(gomock.Any())
		return m
	}
	type mock struct {
		service application_service.Loginer
		msg     user.MsgSender
	}

	tests := []struct {
		name string
		args []string
		mock func(c *gomock.Controller) *mock
		want cli.ExitCode
	}{
		{
			name: "success",
			args: []string{},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					service: func() *application_service_mock.MockLoginer {
						m := application_service_mock.NewMockLoginer(c)
						m.EXPECT().Login().Return(nil)
						return m
					}(),
					msg: defaultMsgSender(c),
				}
			},
			want: cli.ExitOK,
		},
		{
			name: "invalid args",
			args: []string{"invalidarg"},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					msg: defaultMsgSender(c),
				}
			},
			want: cli.ExitErr,
		},
		{
			name: "help",
			args: []string{"--help"},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					msg: defaultMsgSender(c),
				}
			},
			want: cli.ExitErr,
		},
		{
			name: "login error",
			args: []string{},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					msg: defaultMsgSender(c),
				}
			},
			want: cli.ExitErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mock := tt.mock(mockCtrl)
			cmd := cmds.NewLoginCmd(
				mock.service,
				mock.msg,
			)

			// Act
			want := cmd.Run(tt.args)

			// Assert
			if want != tt.want {
				t.Errorf("Run() = %v, want %v", want, tt.want)
			}
		})
	}
}
