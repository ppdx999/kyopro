package init

import (
	"errors"

	"github.com/ppdx999/kyopro/internal/cli"
	"github.com/ppdx999/kyopro/internal/model"
	init_service "github.com/ppdx999/kyopro/internal/service/init"
)

type InitCli struct {
	srvc      init_service.InitService
	msgSender cli.MsgSender
}

func NewInitCli(srvc init_service.InitService, msgSender cli.MsgSender) *InitCli {
	return &InitCli{
		srvc:      srvc,
		msgSender: msgSender,
	}
}

type InitCliOpt struct {
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

func (c *InitCli) ParseArgs(args []string) (*InitCliOpt, error) {
	if len(args) != 1 {
		return nil, errors.New("invalid args")
	}
	if args[0] == "" {
		return nil, errors.New("invalid args")
	}
	if args[0] == "-h" || args[0] == "--help" {
		return nil, errors.New("help flag")
	}

	return &InitCliOpt{
		ContestId: model.ContestId(args[0]),
	}, nil
}

func (c *InitCli) Run(args []string) cli.ExitCode {
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
