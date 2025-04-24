package download

import (
	"bytes"
	"fmt"

	"github.com/ppdx999/kyopro/internal/application/service/downlaod"
	"github.com/ppdx999/kyopro/internal/presentation/cli"
)

type DownloadCmd struct {
	s downlaod.DownloadService
	m cli.MsgSender
}

func NewDownloadCmd(s downlaod.DownloadService, m cli.MsgSender) *DownloadCmd {
	return &DownloadCmd{
		s: s,
		m: m,
	}
}

func (c *DownloadCmd) Name() string {
	return "download"
}

func (c *DownloadCmd) Description() string {
	return "問題のテストケースをダウンロードする"
}

func (c *DownloadCmd) Usage() string {
	var buf bytes.Buffer

	buf.WriteString(
		fmt.Sprintf("%s - %s\n", c.Name(), c.Description()),
	)

	return buf.String()
}

func (c *DownloadCmd) Run(args []string) cli.ExitCode {
	if len(args) != 0 {
		c.m.SendMsg(c.Usage())
		return cli.ExitErr
	}

	if err := c.s.Download(); err != nil {
		c.m.SendMsg(err.Error())
		return cli.ExitErr
	}

	return cli.ExitOK
}
