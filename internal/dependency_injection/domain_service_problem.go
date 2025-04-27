package di

import "github.com/ppdx999/kyopro/internal/domain/service/problem"

func ProblemIdsGetter() problem.ProblemIdsGetter {
	return Atcoder()
}

func ProblemDirMaker() problem.ProblemDirMaker {
	var GetWd = func() problem.GetWd {
		return OperationSystem()
	}
	var PublicDirMaker = func() problem.PublicDirMaker {
		return OperationSystem()
	}

	return problem.NewProblemDirMaker(
		GetWd(),
		PublicDirMaker(),
	)
}

func CurrentProblemLoader() problem.CurrentProblemLoader {
	var GetWd = func() problem.GetWd {
		return OperationSystem()
	}

	return problem.NewCurrentProblemLoader(
		GetWd(),
	)
}
