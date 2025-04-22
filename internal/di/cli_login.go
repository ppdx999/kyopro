package di

import "github.com/ppdx999/kyopro/internal/cli/login"

func InitializeLoginCli() *login.LoginCli {
	return login.NewLoginCli(
		InitializeLoginService(),
		InitializeMsgSender(),
	)
}
