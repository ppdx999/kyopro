package problem

import "github.com/ppdx999/kyopro/internal/domain/model"

type CurrentProblemLoader interface {
	LoadCurrentProblem() (*model.Problem, error)
}

type ProblemDirMaker interface {
	MakeProblemDir(c model.ContestId, p model.ProblemId) error
}

type ProblemIdsGetter interface {
	GetProblemIds(c model.ContestId) ([]model.ProblemId, error)
}
