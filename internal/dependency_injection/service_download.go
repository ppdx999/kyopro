package di

import "github.com/ppdx999/kyopro/internal/application/service/downlaod"

func DownloadService() downlaod.DownloadService {
	return downlaod.NewDownloadServiceImpl(
		CurrentProblemLoader(),
		TestCaseGetter(),
		TestCaseSaver(),
	)
}
