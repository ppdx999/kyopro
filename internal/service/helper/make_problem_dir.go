package helper

import "github.com/ppdx999/kyopro/internal/model"

type MakeProblemDir interface {
	MakeProblemDir(c model.ContestId, p model.ProblemId) error
}
