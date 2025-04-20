package cli

import (
	"errors"

	"github.com/ppdx999/kyopro/internal/infra"
	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/service"
)

type InitCli struct {
	cnsl infra.Console
	srvc service.InitService
}

func NewInitCli(cnsl infra.Console, srvc service.InitService) *InitCli {
	return &InitCli{
		cnsl: cnsl,
		srvc: srvc,
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

func (c *InitCli) ShowUsage() {
	_, err := c.cnsl.WriteErr([]byte(usage))
	if err != nil {
		panic(err)
	}
}

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

func (c *InitCli) Run(args []string) ExitCode {
	opt, err := c.ParseArgs(args)
	if err != nil {
		c.ShowUsage()
		return ExitErr
	}

	if err := c.srvc.Init(opt.ContestId); err != nil {
		c.cnsl.WriteErr([]byte(err.Error()))
		return ExitErr
	}

	return ExitOK
}
