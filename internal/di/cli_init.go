package di

import init_ "github.com/ppdx999/kyopro/internal/cli/init"

func InitializeInitCmd() *init_.InitCmd {
	return init_.NewInitCmd(
		InitializeInitService(),
		InitializeMsgSender(),
	)
}
