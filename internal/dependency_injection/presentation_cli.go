package di

import (
	"github.com/ppdx999/kyopro/internal/presentation/cli"
	"github.com/ppdx999/kyopro/internal/presentation/cli/cmds"
)

func Dispatcher() cli.Dispatcher {
	return *cli.NewDispatcher(MsgSender())
}

func Cmd() cli.Cmd {
	var DownloadCmd = func() *cmds.DownloadCmd {
		return cmds.NewDownloadCmd(
			ApplicationServiceDownload(),
			MsgSender(),
		)
	}
	var InitCmd = func() *cmds.InitCmd {
		return cmds.NewInitCmd(
			ApplicationServiceInit(),
			MsgSender(),
		)
	}
	var LoginCmd = func() *cmds.LoginCmd {
		return cmds.NewLoginCmd(
			ApplicationServiceLogin(),
			MsgSender(),
		)
	}

	cmd := Dispatcher()
	cmd.Register(InitCmd())
	cmd.Register(LoginCmd())
	cmd.Register(DownloadCmd())
	return &cmd
}
