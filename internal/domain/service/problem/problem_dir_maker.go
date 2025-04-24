package problem

import (
	"path/filepath"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

type ProblemDirMaker interface {
	MakeProblemDir(c model.ContestId, p model.ProblemId) error
}

type ProblemDirMakerImpl struct {
	wd       GetWd
	dirMaker PublicDirMaker
}

func NewProblemDirMakerImpl(
	wd GetWd,
	dirMaker PublicDirMaker,
) *ProblemDirMakerImpl {
	return &ProblemDirMakerImpl{wd: wd, dirMaker: dirMaker}
}

func (s *ProblemDirMakerImpl) MakeProblemDir(c model.ContestId, p model.ProblemId) error {
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
