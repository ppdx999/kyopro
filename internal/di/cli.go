package di

import "github.com/ppdx999/kyopro/internal/cli"

func InitializeMsgSender() cli.MsgSender {
	return InitializeConsoleMsgSender()
}
