package cmds_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	application_service "github.com/ppdx999/kyopro/internal/application/service"
	application_service_mock "github.com/ppdx999/kyopro/internal/application/service/mock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/user"
	user_mock "github.com/ppdx999/kyopro/internal/domain/service/user/mock"
	"github.com/ppdx999/kyopro/internal/presentation/cli"
	"github.com/ppdx999/kyopro/internal/presentation/cli/cmds"
)

func InitCmdTestRun(t *testing.T) {
	var defaultMsgSender = func(c *gomock.Controller) *user_mock.MockMsgSender {
		m := user_mock.NewMockMsgSender(c)
		m.EXPECT().SendMsg(gomock.Any())
		return m
	}
	type mock struct {
		service application_service.Initer
		msg     user.MsgSender
	}
	tests := []struct {
		name string
		args []string
		mock func(c *gomock.Controller) *mock
		want cli.ExitCode
	}{
		{
			name: "valid case",
			args: []string{"abc123"},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					service: func() *application_service_mock.MockIniter {
						m := application_service_mock.NewMockIniter(c)
						m.EXPECT().Init(model.ContestId("abc123")).Return(nil)
						return m
					}(),
					msg: defaultMsgSender(c),
				}
			},
			want: cli.ExitOK,
		},
		{
			name: "too many args",
			args: []string{"abc123", "def456"},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					msg: defaultMsgSender(c),
				}
			},
			want: cli.ExitErr,
		},
		{
			name: "blank args",
			args: []string{},
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
			name: "usecase error",
			args: []string{"abc123"},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					service: func() *application_service_mock.MockIniter {
						m := application_service_mock.NewMockIniter(c)
						m.EXPECT().
							Init(model.ContestId("abc123")).
							Return(errors.New("usecase error"))
						return m
					}(),
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

			initCli := cmds.NewInitCmd(
				mock.service,
				mock.msg,
			)

			// Act
			got := initCli.Run(tt.args)

			// Assert
			if got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
