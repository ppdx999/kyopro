package di

import "github.com/ppdx999/kyopro/internal/domain/service/problem"

func ProblemIdsGetter() problem.ProblemIdsGetter {
	return Atcoder()
}

func ProblemDirMaker() problem.ProblemDirMaker {
	var GetWd = func() problem.GetWd {
		return FileSystem()
	}
	var PublicDirMaker = func() problem.PublicDirMaker {
		return FileSystem()
	}

	return problem.NewProblemDirMakerImpl(
		GetWd(),
		PublicDirMaker(),
	)
}

func CurrentProblemLoader() problem.CurrentProblemLoader {
	var GetWd = func() problem.GetWd {
		return FileSystem()
	}

	return problem.NewCurrentProblemLoaderImpl(
		GetWd(),
	)
}
