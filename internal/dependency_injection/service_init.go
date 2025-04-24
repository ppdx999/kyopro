package di

import init_ "github.com/ppdx999/kyopro/internal/application/service/init"

func InitService() init_.InitService {
	var GetProblemIds = func() init_.GetProblemIds {
		return Atcoder()
	}

	var ProblemDirMaker = func() init_.ProblemDirMaker {
		return ProblemDirMakerImpl()
	}

	return init_.NewInitServiceImpl(
		GetProblemIds(),
		ProblemDirMaker(),
	)

}
