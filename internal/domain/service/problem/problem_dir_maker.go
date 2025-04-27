package problem

import (
	"path/filepath"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

type problemDirMaker struct {
	wd       GetWd
	dirMaker PublicDirMaker
}

func (s *problemDirMaker) MakeProblemDir(c model.ContestId, p model.ProblemId) error {
	cwd, err := s.wd.GetWd()
	if err != nil {
		return err
	}
	path := filepath.Join(cwd, string(c), string(p))
	if err := s.dirMaker.MakePublicDir(path); err != nil {
		return err
	}
	return nil
}
