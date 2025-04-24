package init

import "github.com/ppdx999/kyopro/internal/domain/model"

type GetProblemIds interface {
	GetProblemIds(c model.ContestId) ([]model.ProblemId, error)
}

type ProblemDirMaker interface {
	MakeProblemDir(c model.ContestId, p model.ProblemId) error
}
