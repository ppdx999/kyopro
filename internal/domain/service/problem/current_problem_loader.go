package problem

import (
	"errors"
	"path/filepath"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

type currentProblemLoader struct {
	wd WdGetter
}

func (l *currentProblemLoader) LoadCurrentProblem() (*model.Problem, error) {
	cwd, err := l.wd.GetWd()
	if err != nil {
		return nil, err
	}

	pid := filepath.Base(cwd)
	if pid == "/" {
		return nil, errors.New("contest or problem not found")
	}

	p := model.NewProblem(pid)

	cid := filepath.Base(filepath.Dir(cwd))
	if cid == "/" {
		return nil, errors.New("contest or problem not found")
	}

	c := model.NewContest(cid)
	p.Contest = c
	return p, nil
}
