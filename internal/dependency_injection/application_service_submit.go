package di

import application_service "github.com/ppdx999/kyopro/internal/application/service"

func ApplicationServiceSubmit() application_service.Submiter {
	var ClipboardWriter = func() application_service.ClipboardWriter {
		return OperationSystem()
	}
	var SubmitPageOpener = func() application_service.SubmitPageOpener {
		return Atcoder()
	}

	return application_service.NewSubmiter(
		LanguageDetector(),
		CurrentProblemLoader(),
		ClipboardWriter(),
		SubmitPageOpener(),
	)
}
