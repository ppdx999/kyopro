package login

import (
	"bytes"
	"fmt"

	"github.com/ppdx999/kyopro/internal/cli"
	"github.com/ppdx999/kyopro/internal/service/login"
)

type LoginCmd struct {
	srvc      login.LoginService
	msgSender cli.MsgSender
}

func NewLoginCmd(srvc login.LoginService, msgSender cli.MsgSender) *LoginCmd {
	return &LoginCmd{
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

func (c *LoginCmd) Name() string {
	return "login"
}

func (c *LoginCmd) Description() string {
	return "ユーザーからセッション情報を受け取りサービスにログインします"
}

func (c *LoginCmd) Usage() string {
	var buf bytes.Buffer

	buf.WriteString(
		fmt.Sprintf("%s - %s\n", c.Name(), c.Description()),
	)

	buf.WriteString(usage)

	return buf.String()
}

func (c *LoginCmd) Run(args []string) cli.ExitCode {
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
