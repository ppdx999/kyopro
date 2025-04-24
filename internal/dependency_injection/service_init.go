package di

import init_ "github.com/ppdx999/kyopro/internal/application/service/init"

func InitService() init_.InitService {
	return init_.NewInitServiceImpl(
		ProblemIdsGetter(),
		ProblemDirMaker(),
	)
}
