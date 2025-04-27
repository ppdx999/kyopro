package di

import application_service "github.com/ppdx999/kyopro/internal/application/service"

func ApplicationServiceLogin() application_service.Loginer {
	return application_service.NewLoginer(
		SessionAsker(),
		LoginChecker(),
		SessionSaver(),
		MsgSender(),
	)
}
