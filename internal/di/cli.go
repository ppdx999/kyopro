package di

import "github.com/ppdx999/kyopro/internal/cli"

func InitializeMsgSender() cli.MsgSender {
	return InitializeConsoleMsgSender()
}

func InitializeDispatcher() cli.Dispatcher {
	return *cli.NewDispatcher(InitializeMsgSender())
}

func InitializeCmd() cli.Cmd {
	cmd := InitializeDispatcher()
	cmd.Register("init", InitializeInitCmd())
	cmd.Register("login", InitializeLoginCmd())
	return &cmd
}
