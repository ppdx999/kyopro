package di

import "github.com/ppdx999/kyopro/internal/problem"

func ProblemDirMakerImpl() *problem.ProblemDirMakerImpl {
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

func CurrentProblemLoaderImpl() *problem.CurrentProblemLoaderImpl {
	var GetWd = func() problem.GetWd {
		return FsImpl()
	}

	return problem.NewCurrentProblemLoaderImpl(
		GetWd(),
	)
}
