package di

import "github.com/ppdx999/kyopro/internal/application/service/downlaod"

func DownloadService() downlaod.DownloadService {
	var ContestAndProblemLoader = func() downlaod.CurrentProblemLoader {
		return CurrentProblemLoaderImpl()
	}
	var TestCaseGetter = func() downlaod.TestCasesGetter {
		return Atcoder()
	}
	var TestCaseSaver = func() downlaod.TestCaseSaver {
		return TestCaseFsSaver()
	}
	return downlaod.NewDownloadServiceImpl(
		ContestAndProblemLoader(),
		TestCaseGetter(),
		TestCaseSaver(),
	)
}
