package cli_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	user_mock "github.com/ppdx999/kyopro/internal/domain/service/user/mock"
	"github.com/ppdx999/kyopro/internal/presentation/cli"
)

func TestRun(t *testing.T) {
	type mock struct {
		sub1 *MockCmd
		sub2 *MockCmd
	}
	tests := []struct {
		name string
		args []string
		mock func(c *gomock.Controller) *mock
		want cli.ExitCode
	}{
		{
			name: "Correct Case Call sub1",
			args: []string{"sub1"},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sub1: func() *MockCmd {
						m := NewMockCmd(c)
						m.EXPECT().Name().Return("sub1")
						m.EXPECT().Run(gomock.Any()).Return(cli.ExitOK)
						return m
					}(),
					sub2: func() *MockCmd {
						m := NewMockCmd(c)
						m.EXPECT().Name().Return("sub2")
						return m
					}(),
				}
			},
			want: cli.ExitOK,
		},
		{
			name: "Correct Case Call sub2",
			args: []string{"sub2"},
			want: cli.ExitOK,
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sub1: func() *MockCmd {
						m := NewMockCmd(c)
						m.EXPECT().Name().Return("sub1")
						return m
					}(),
					sub2: func() *MockCmd {
						m := NewMockCmd(c)
						m.EXPECT().Name().Return("sub2")
						m.EXPECT().Run(gomock.Any()).Return(cli.ExitOK)
						return m
					}(),
				}
			},
		},
		{
			name: "No subcommand args provided",
			args: nil,
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sub1: func() *MockCmd {
						m := NewMockCmd(c)
						m.EXPECT().Name().Return("sub1")
						m.EXPECT().Description().Return("sub1 description")
						return m
					}(),
					sub2: func() *MockCmd {
						m := NewMockCmd(c)
						m.EXPECT().Name().Return("sub2")
						m.EXPECT().Description().Return("sub2 description")
						return m
					}(),
				}
			},
			want: cli.ExitErr,
		},
		{
			name: "Unknown subcommand",
			args: []string{"unknown"},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sub1: func() *MockCmd {
						m := NewMockCmd(c)
						m.EXPECT().Name().Return("sub1")
						m.EXPECT().Description().Return("sub1 description")
						return m
					}(),
					sub2: func() *MockCmd {
						m := NewMockCmd(c)
						m.EXPECT().Name().Return("sub2")
						m.EXPECT().Description().Return("sub2 description")
						return m
					}(),
				}
			},
			want: cli.ExitErr,
		},
		{
			name: "Help flag",
			args: []string{"-h"},
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					sub1: func() *MockCmd {
						m := NewMockCmd(c)
						m.EXPECT().Name().Return("sub1")
						m.EXPECT().Description().Return("sub1 description")
						return m
					}(),
					sub2: func() *MockCmd {
						m := NewMockCmd(c)
						m.EXPECT().Name().Return("sub2")
						m.EXPECT().Description().Return("sub2 description")
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

			msgSender := user_mock.NewMockMsgSender(mockCtrl)
			msgSender.EXPECT().SendMsg(gomock.Any()).AnyTimes()

			root := cli.NewDispatcher(msgSender)
			root.Register(mock.sub1)
			root.Register(mock.sub2)

			// Act
			got := root.Run(tt.args)

			// Assert
			if got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
