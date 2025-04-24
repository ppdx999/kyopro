package di

import "github.com/ppdx999/kyopro/internal/domain/service/problem"

func ProblemIdsGetter() problem.ProblemIdsGetter {
	return Atcoder()
}

func ProblemDirMaker() problem.ProblemDirMaker {
	var GetWd = func() problem.GetWd {
		return FsImpl()
	}
	var PublicDirMaker = func() problem.PublicDirMaker {
		return FsImpl()
	}

	return problem.NewProblemDirMakerImpl(
		GetWd(),
		PublicDirMaker(),
	)
}

func CurrentProblemLoader() problem.CurrentProblemLoader {
	var GetWd = func() problem.GetWd {
		return FsImpl()
	}

	return problem.NewCurrentProblemLoaderImpl(
		GetWd(),
	)
}
