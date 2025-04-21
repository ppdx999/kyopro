package init

import "github.com/ppdx999/kyopro/internal/model"

type MakePublicDir interface {
	MakePublicDir(path string) error
}

type GetWd interface {
	GetWd() (string, error)
}

type GetProblemIds interface {
	GetProblemIds(c model.ContestId) ([]model.ProblemId, error)
}
