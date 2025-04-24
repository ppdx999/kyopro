package di

import init_ "github.com/ppdx999/kyopro/internal/presentation/cli/init"

func InitCmd() *init_.InitCmd {
	return init_.NewInitCmd(
		InitService(),
		MsgSender(),
	)
}
