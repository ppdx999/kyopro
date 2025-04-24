package di

import "github.com/ppdx999/kyopro/internal/application/service/login"

func LoginService() login.LoginService {
	return login.NewLoginServiceImpl(
		SessionAsker(),
		LoginChecker(),
		SessionSaver(),
		MsgSender(),
	)
}
