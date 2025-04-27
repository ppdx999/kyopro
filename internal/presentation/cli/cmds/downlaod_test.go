package cmds_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	application_service "github.com/ppdx999/kyopro/internal/application/service"
	application_service_mock "github.com/ppdx999/kyopro/internal/application/service/mock"
	"github.com/ppdx999/kyopro/internal/domain/service/user"
	user_mock "github.com/ppdx999/kyopro/internal/domain/service/user/mock"
	"github.com/ppdx999/kyopro/internal/presentation/cli"
	"github.com/ppdx999/kyopro/internal/presentation/cli/cmds"
)

func DownloadCmdTestRun(t *testing.T) {
	type mock struct {
		service application_service.Downloader
		msg     user.MsgSender
	}

	tests := []struct {
		name string
		args []string
		mock func(c *gomock.Controller) *mock
		want cli.ExitCode
	}{
		{
			name: "正常系",
			args: []string{},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					service: func() *application_service_mock.MockDownloader {
						m := application_service_mock.NewMockDownloader(c)
						m.EXPECT().Download().Return(nil)
						return m
					}(),
					msg: func() *user_mock.MockMsgSender {
						m := user_mock.NewMockMsgSender(c)
						m.EXPECT().SendMsg(gomock.Any())
						return m
					}(),
				}
			},
			want: cli.ExitOK,
		},
		{
			name: "引数エラー",
			args: []string{"invalidarg"}, // downloadコマンドは引数を取らない
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					msg: func() *user_mock.MockMsgSender {
						m := user_mock.NewMockMsgSender(c)
						m.EXPECT().SendMsg(gomock.Any())
						return m
					}(),
				}
			},
			want: cli.ExitErr,
		},
		{
			name: "Downloadサービスエラー",
			args: []string{},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					service: func() *application_service_mock.MockDownloader {
						m := application_service_mock.NewMockDownloader(c)
						m.EXPECT().Download().Return(errors.New("download error"))
						return m
					}(),
					msg: func() *user_mock.MockMsgSender {
						m := user_mock.NewMockMsgSender(c)
						m.EXPECT().SendMsg(gomock.Any())
						return m
					}(),
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
			cmd := cmds.NewDownloadCmd(
				mock.service,
				mock.msg,
			)

			// Act
			got := cmd.Run(tt.args)

			// Assert
			if tt.want != got {
				t.Errorf("Run() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
