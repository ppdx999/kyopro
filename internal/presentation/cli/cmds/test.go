package cmds

import (
	"bytes"
	"fmt"

	application_service "github.com/ppdx999/kyopro/internal/application/service"
	"github.com/ppdx999/kyopro/internal/domain/service/user"
	"github.com/ppdx999/kyopro/internal/presentation/cli"
)

type TestCmd struct {
	srvc      application_service.Tester
	msgSender user.MsgSender
}

func NewTestCmd(srvc application_service.Tester, msgSender user.MsgSender) *TestCmd {
	return &TestCmd{
		srvc:      srvc,
		msgSender: msgSender,
	}
}

func (c *TestCmd) Name() string {
	return "test"
}

func (c *TestCmd) Description() string {
	return "テストを実行します"
}

func (c *TestCmd) Usage() string {
	var usage = `
Usage:
	kyopro test

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

func (c *TestCmd) Run(args []string) cli.ExitCode {
	if len(args) != 0 {
		c.msgSender.SendMsg(c.Usage())
		return cli.ExitErr
	}

	if err := c.srvc.Test(); err != nil {
		c.msgSender.SendMsg(err.Error())
		return cli.ExitErr
	}

	return cli.ExitOK
}
