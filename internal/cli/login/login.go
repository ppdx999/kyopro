package login

import (
	"github.com/ppdx999/kyopro/internal/cli"
	"github.com/ppdx999/kyopro/internal/service/login"
)

type LoginCli struct {
	srvc      login.LoginService
	msgSender cli.MsgSender
}

func NewLoginCli(srvc login.LoginService, msgSender cli.MsgSender) *LoginCli {
	return &LoginCli{
		srvc:      srvc,
		msgSender: msgSender,
	}
}

var usage = `
Usage:
	kyopro login

Options:
	-h, --help  ヘルプの表示
`

func (c *LoginCli) Run(args []string) cli.ExitCode {
	if len(args) != 0 {
		c.msgSender.SendMsg(usage)
		return cli.ExitErr
	}

	if err := c.srvc.Login(); err != nil {
		c.msgSender.SendMsg(err.Error())
		return cli.ExitErr
	}

	return cli.ExitOK
}
