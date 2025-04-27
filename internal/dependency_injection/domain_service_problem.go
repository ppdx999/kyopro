package di

import "github.com/ppdx999/kyopro/internal/domain/service/problem"

func ProblemIdsGetter() problem.ProblemIdsGetter {
	return Atcoder()
}

func ProblemDirMaker() problem.ProblemDirMaker {
	var WdGetter = func() problem.WdGetter {
		return OperationSystem()
	}
	var PublicDirMaker = func() problem.PublicDirMaker {
		return OperationSystem()
	}

	return problem.NewProblemDirMaker(
		WdGetter(),
		PublicDirMaker(),
	)
}

func CurrentProblemLoader() problem.CurrentProblemLoader {
	var WdGetter = func() problem.WdGetter {
		return OperationSystem()
	}

	return problem.NewCurrentProblemLoader(
		WdGetter(),
	)
}
