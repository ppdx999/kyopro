package init

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ppdx999/kyopro/internal/cli"
	"github.com/ppdx999/kyopro/internal/model"
	init_service "github.com/ppdx999/kyopro/internal/service/init"
)

type InitCmd struct {
	srvc      init_service.InitService
	msgSender cli.MsgSender
}

func NewInitCmd(srvc init_service.InitService, msgSender cli.MsgSender) *InitCmd {
	return &InitCmd{
		srvc:      srvc,
		msgSender: msgSender,
	}
}

type InitCmdOpt struct {
	ContestId model.ContestId
}

var usage = `
Usage:
	kyopro init <contest_id>

Args:
	contest_id  ContestのID

Options:
	-h, --help  Show this screen.
`

func (c *InitCmd) Name() string {
	return "init"
}

func (c *InitCmd) Description() string {
	return "コンテストの初期設定を行います"
}

func (c *InitCmd) Usage() string {
	var buf bytes.Buffer

	buf.WriteString(
		fmt.Sprintf("%s - %s\n", c.Name(), c.Description()),
	)

	buf.WriteString(usage)

	return buf.String()
}

func (c *InitCmd) ParseArgs(args []string) (*InitCmdOpt, error) {
	if len(args) != 1 {
		return nil, errors.New("invalid args")
	}
	if args[0] == "" {
		return nil, errors.New("invalid args")
	}
	if args[0] == "-h" || args[0] == "--help" {
		return nil, errors.New("help flag")
	}

	return &InitCmdOpt{
		ContestId: model.ContestId(args[0]),
	}, nil
}

func (c *InitCmd) Run(args []string) cli.ExitCode {
	opt, err := c.ParseArgs(args)
	if err != nil {
		c.msgSender.SendMsg(usage)
		return cli.ExitErr
	}

	if err := c.srvc.Init(opt.ContestId); err != nil {
		c.msgSender.SendMsg(err.Error())
		return cli.ExitErr
	}

	return cli.ExitOK
}
