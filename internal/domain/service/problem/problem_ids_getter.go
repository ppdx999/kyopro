package problem

import "github.com/ppdx999/kyopro/internal/domain/model"

type ProblemIdsGetter interface {
	GetProblemIds(c model.ContestId) ([]model.ProblemId, error)
}
