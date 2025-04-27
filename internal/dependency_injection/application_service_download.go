package di

import application_service "github.com/ppdx999/kyopro/internal/application/service"

func ApplicationServiceDownload() application_service.Downloader {
	return application_service.NewDownloader(
		CurrentProblemLoader(),
		TestCaseGetter(),
		TestCaseSaver(),
	)
}
