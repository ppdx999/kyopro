package helper

import (
	"path/filepath"

	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/service/helper/port"
)

/*
MakeProblemDirは問題ディレクトリをWorkspaceに作成します
*/
type MakeProblemDir interface {
	MakeProblemDir(c model.ContestId, p model.ProblemId) error
}

type MakeProblemDirImpl struct {
	MakePublicDir port.MakePublicDir
	GetWorkspace  GetWorkspace
}

func (s *MakeProblemDirImpl) MakeProblemDir(c model.ContestId, p model.ProblemId) error {
	w, err := s.GetWorkspace.GetWorkspace()
	if err != nil {
		return err
	}
	path := filepath.Join(string(w.Path), string(c), string(p))
	if err := s.MakePublicDir.MakePublicDir(path); err != nil {
		return err
	}
	return nil
}
