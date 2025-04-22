package di

import "github.com/ppdx999/kyopro/internal/cli/login"

func LoginCmd() *login.LoginCmd {
	return login.NewLoginCmd(
		LoginService(),
		MsgSender(),
	)
}
