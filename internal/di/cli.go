package di

import "github.com/ppdx999/kyopro/internal/cli"

func MsgSender() cli.MsgSender {
	return ConsoleMsgSender()
}

func Dispatcher() cli.Dispatcher {
	return *cli.NewDispatcher(MsgSender())
}

func Cmd() cli.Cmd {
	cmd := Dispatcher()
	cmd.Register("init", InitCmd())
	cmd.Register("login", LoginCmd())
	return &cmd
}
