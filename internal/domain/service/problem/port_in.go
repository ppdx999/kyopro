package problem

import "github.com/ppdx999/kyopro/internal/domain/model"

type CurrentProblemLoader interface {
	LoadCurrentProblem() (*model.Problem, error)
}

func NewCurrentProblemLoader(wd GetWd) CurrentProblemLoader {
	return &currentProblemLoader{wd: wd}
}

type ProblemDirMaker interface {
	MakeProblemDir(c model.ContestId, p model.ProblemId) error
}

func NewProblemDirMaker(
	wd GetWd,
	dirMaker PublicDirMaker,
) ProblemDirMaker {
	return &problemDirMaker{wd: wd, dirMaker: dirMaker}
}

type ProblemIdsGetter interface {
	GetProblemIds(c model.ContestId) ([]model.ProblemId, error)
}
