package di

import "github.com/ppdx999/kyopro/internal/cli/download"

func DownloadCmd() *download.DownloadCmd {
	return download.NewDownloadCmd(
		DownloadService(),
		MsgSender(),
	)
}
