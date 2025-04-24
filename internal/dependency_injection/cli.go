package di

import "github.com/ppdx999/kyopro/internal/presentation/cli"

func MsgSender() cli.MsgSender {
	return ConsoleMsgSender()
}

func Dispatcher() cli.Dispatcher {
	return *cli.NewDispatcher(MsgSender())
}

func Cmd() cli.Cmd {
	cmd := Dispatcher()
	cmd.Register(InitCmd())
	cmd.Register(LoginCmd())
	cmd.Register(DownloadCmd())
	return &cmd
}
