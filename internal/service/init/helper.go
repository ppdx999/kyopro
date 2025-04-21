package init

import (
	"path/filepath"

	"github.com/ppdx999/kyopro/internal/model"
)

/*
GetWorkspaceはWorkspaceを取得します
*/
type GetWorkspace interface {
	GetWorkspace() (*model.Workspace, error)
}

type GetWorkspaceImpl struct {
	GetWd GetWd
}

/*
GetWorkspaceは現在のワークディレクトリをもとにWorkspaceを取得します
*/
func (g *GetWorkspaceImpl) GetWorkspace() (*model.Workspace, error) {
	wd, err := g.GetWd.GetWd()
	if err != nil {
		return nil, err
	}
	return model.NewWorkspace(wd), nil
}

func NewGetWorkspaceImpl(GetWd GetWd) *GetWorkspaceImpl {
	return &GetWorkspaceImpl{
		GetWd: GetWd,
	}
}

/*
MakeProblemDirは問題ディレクトリをWorkspaceに作成します
*/
type MakeProblemDir interface {
	MakeProblemDir(c model.ContestId, p model.ProblemId) error
}

type MakeProblemDirImpl struct {
	MakePublicDir MakePublicDir
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

func NewMakeProblemDirImpl(
	MakePublicDir MakePublicDir,
	GetWorkspace GetWorkspace,
) *MakeProblemDirImpl {
	return &MakeProblemDirImpl{
		MakePublicDir: MakePublicDir,
		GetWorkspace:  GetWorkspace,
	}
}
