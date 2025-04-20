//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ppdx999/kyopro/internal/cli"
	"github.com/ppdx999/kyopro/internal/infra"
	"github.com/ppdx999/kyopro/internal/service"
	service_helper "github.com/ppdx999/kyopro/internal/service/helper"
	service_helper_port "github.com/ppdx999/kyopro/internal/service/helper/port"
)

func InitializeGetWorkspaceImpl() *service_helper.GetWorkspaceImpl {
	wire.Build(
		infra.NewFsImple,
		wire.Bind(new(service_helper_port.GetWd), new(*infra.FsImpl)),
		service_helper.NewGetWorkspaceImpl,
	)
	return &service_helper.GetWorkspaceImpl{}
}

func InitializeMakeProblemDirImpl() *service_helper.MakeProblemDirImpl {
	wire.Build(
		wire.Bind(new(service_helper_port.MakePublicDir), new(*infra.FsImpl)),
		infra.NewFsImple,
		InitializeGetWorkspaceImpl,
		wire.Bind(new(service_helper.GetWorkspace), new(*service_helper.GetWorkspaceImpl)),
		service_helper.NewMakeProblemDirImpl,
	)
	return &service_helper.MakeProblemDirImpl{}
}

func InitializeInitServiceImpl() *service.InitServiceImpl {
	wire.Build(
		InitializeMakeProblemDirImpl,
		wire.Bind(new(service_helper.MakeProblemDir), new(*service_helper.MakeProblemDirImpl)),
		service_helper.NewGetProblemIdsImpl,
		wire.Bind(new(service_helper.GetProblemIds), new(*service_helper.GetProblemIdsImpl)),
		service.NewInitServiceImpl,
	)
	return &service.InitServiceImpl{}
}

func InitializeInitCli() *cli.InitCli {
	wire.Build(
		infra.NewConsoleImpl,
		wire.Bind(new(infra.Console), new(*infra.ConsoleImpl)),
		InitializeInitServiceImpl,
		wire.Bind(new(service.InitService), new(*service.InitServiceImpl)),
		cli.NewInitCli,
	)
	return &cli.InitCli{}
}
