package cmds

import (
	"bytes"
	"fmt"

	application_service "github.com/ppdx999/kyopro/internal/application/service"
	"github.com/ppdx999/kyopro/internal/domain/service/user"
	"github.com/ppdx999/kyopro/internal/presentation/cli"
)

type LoginCmd struct {
	srvc      application_service.Loginer
	msgSender user.MsgSender
}

func NewLoginCmd(srvc application_service.Loginer, msgSender user.MsgSender) *LoginCmd {
	return &LoginCmd{
		srvc:      srvc,
		msgSender: msgSender,
	}
}

func (c *LoginCmd) Name() string {
	return "login"
}

func (c *LoginCmd) Description() string {
	return "ユーザーからセッション情報を受け取りサービスにログインします"
}

func (c *LoginCmd) Usage() string {
	var usage = `
Usage:
	kyopro login

Options:
	-h, --help  ヘルプの表示
`
	var buf bytes.Buffer

	buf.WriteString(
		fmt.Sprintf("%s - %s\n", c.Name(), c.Description()),
	)

	buf.WriteString(usage)

	return buf.String()
}

func (c *LoginCmd) Run(args []string) cli.ExitCode {
	if len(args) != 0 {
		c.msgSender.SendMsg(c.Usage())
		return cli.ExitErr
	}

	if err := c.srvc.Login(); err != nil {
		c.msgSender.SendMsg(err.Error())
		return cli.ExitErr
	}

	return cli.ExitOK
}
