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

func InitializeConsole() infra.Console {
	wire.Build(
		infra.NewConsoleImpl,
		wire.Bind(new(infra.Console), new(*infra.ConsoleImpl)),
	)
	return nil
}

func InitializeGetWd() service_helper_port.GetWd {
	wire.Build(
		infra.NewFsImple,
		wire.Bind(new(service_helper_port.GetWd), new(*infra.FsImpl)),
	)
	return nil
}

func InitializeMakePublicDir() service_helper_port.MakePublicDir {
	wire.Build(
		infra.NewFsImple,
		wire.Bind(new(service_helper_port.MakePublicDir), new(*infra.FsImpl)),
	)
	return nil
}

func InitializeGetWorkspace() service_helper.GetWorkspace {
	wire.Build(
		InitializeGetWd,
		service_helper.NewGetWorkspaceImpl,
		wire.Bind(new(service_helper.GetWorkspace), new(*service_helper.GetWorkspaceImpl)),
	)
	return nil
}

func InitializeMakeProblemDir() service_helper.MakeProblemDir {
	wire.Build(
		InitializeMakePublicDir,
		InitializeGetWorkspace,
		service_helper.NewMakeProblemDirImpl,
		wire.Bind(new(service_helper.MakeProblemDir), new(*service_helper.MakeProblemDirImpl)),
	)
	return nil
}

func InitializeGetProblemIds() service_helper.GetProblemIds {
	wire.Build(
		service_helper.NewGetProblemIdsImpl,
		wire.Bind(new(service_helper.GetProblemIds), new(*service_helper.GetProblemIdsImpl)),
	)
	return nil
}

func InitializeInitService() service.InitService {
	wire.Build(
		InitializeGetProblemIds,
		InitializeMakeProblemDir,
		service.NewInitServiceImpl,
		wire.Bind(new(service.InitService), new(*service.InitServiceImpl)),
	)
	return nil
}

func InitializeInitCli() *cli.InitCli {
	wire.Build(
		InitializeConsole,
		InitializeInitService,
		cli.NewInitCli,
	)
	return &cli.InitCli{}
}
