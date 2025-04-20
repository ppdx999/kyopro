package helper

import "github.com/ppdx999/kyopro/internal/model"

type GetProblemIds interface {
	GetProblemIds(c model.ContestId) ([]model.ProblemId, error)
}

type GetProblemIdsImpl struct{}

func (g *GetProblemIdsImpl) GetProblemIds(c model.ContestId) ([]model.ProblemId, error) {
	// TOOD: Implement
	return []model.ProblemId{"A", "B", "C"}, nil
}

func NewGetProblemIdsImpl() *GetProblemIdsImpl {
	return &GetProblemIdsImpl{}
}
