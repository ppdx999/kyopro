package di

import "github.com/ppdx999/kyopro/internal/cli/login"

func InitializeLoginCmd() *login.LoginCmd {
	return login.NewLoginCmd(
		InitializeLoginService(),
		InitializeMsgSender(),
	)
}
