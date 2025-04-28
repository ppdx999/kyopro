package cmds

import (
	"bytes"
	"fmt"

	application_service "github.com/ppdx999/kyopro/internal/application/service"
	"github.com/ppdx999/kyopro/internal/domain/service/user"
	"github.com/ppdx999/kyopro/internal/presentation/cli"
)

type SubmitCmd struct {
	srvc      application_service.Submiter
	msgSender user.MsgSender
}

func NewSubmitCmd(srvc application_service.Submiter, msgSender user.MsgSender) *SubmitCmd {
	return &SubmitCmd{
		srvc:      srvc,
		msgSender: msgSender,
	}
}

func (c *SubmitCmd) Name() string {
	return "submit"
}

func (c *SubmitCmd) Description() string {
	return "コードを提出します"
}

func (c *SubmitCmd) Usage() string {
	var usage = `
Usage:
	kyopro submit

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

func (c *SubmitCmd) Run(args []string) cli.ExitCode {
	if len(args) != 0 {
		c.msgSender.SendMsg(c.Usage())
		return cli.ExitErr
	}

	if err := c.srvc.Submit(); err != nil {
		c.msgSender.SendMsg(err.Error())
		return cli.ExitErr
	}

	return cli.ExitOK
}
