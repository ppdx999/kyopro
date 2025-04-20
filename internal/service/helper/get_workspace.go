package helper

import (
	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/service/helper/port"
)

/*
GetWorkspaceはWorkspaceを取得します
*/
type GetWorkspace interface {
	GetWorkspace() (*model.Workspace, error)
}

type GetWorkspaceImpl struct {
	GetWd port.GetWd
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
