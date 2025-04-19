package out

import "github.com/ppdx999/kyopro/internal/application/domain/model"

type GetProblemIdsPort interface {
	GetProblemIds(c model.ContestId) ([]model.ProblemId, error)
}
