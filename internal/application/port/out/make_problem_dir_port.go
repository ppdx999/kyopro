package out

import "github.com/ppdx999/kyopro/internal/application/domain/model"

type MakeProblemDirPort interface {
	MakeProblemDir(c model.ContestId, p model.ProblemId) error
}
