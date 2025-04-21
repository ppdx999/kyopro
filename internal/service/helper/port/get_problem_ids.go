package port

import "github.com/ppdx999/kyopro/internal/model"

type GetProblemIds interface {
	GetProblemIds(c model.ContestId) ([]model.ProblemId, error)
}
