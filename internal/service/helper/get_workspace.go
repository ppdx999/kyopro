package helper

import "github.com/ppdx999/kyopro/internal/model"

/*
GetWorkspaceはWorkspaceを取得します
*/
type GetWorkspace interface {
	GetWorkspace() (*model.Workspace, error)
}
