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

/*********************************************************************
Infra
**********************************************************************/

func InitializeConsole() infra.Console {
	wire.Build(
		infra.NewConsoleImpl,
		wire.Bind(new(infra.Console), new(*infra.ConsoleImpl)),
	)
	return nil
}

func InitializeRequester() infra.Requester {
	wire.Build(
		infra.NewRequesterImpl,
		wire.Bind(new(infra.Requester), new(*infra.RequesterImpl)),
	)
	return nil
}

func InitializeAtcoder() *infra.Atcoder {
	wire.Build(
		InitializeRequester,
		infra.NewAtcoder,
	)
	return &infra.Atcoder{}
}

/*********************************************************************
Service Helper Port
**********************************************************************/

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

func InitializeGetProblemIds() service_helper_port.GetProblemIds {
	wire.Build(
		InitializeAtcoder,
		wire.Bind(new(service_helper_port.GetProblemIds), new(*infra.Atcoder)),
	)
	return nil
}

/*
********************************************************************
Service Helper
*********************************************************************
*/
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

/*
********************************************************************
Service
*********************************************************************
*/
func InitializeInitService() service.InitService {
	wire.Build(
		InitializeGetProblemIds,
		InitializeMakeProblemDir,
		service.NewInitServiceImpl,
		wire.Bind(new(service.InitService), new(*service.InitServiceImpl)),
	)
	return nil
}

/*
********************************************************************
Cli
*********************************************************************
*/
func InitializeInitCli() *cli.InitCli {
	wire.Build(
		InitializeConsole,
		InitializeInitService,
		cli.NewInitCli,
	)
	return &cli.InitCli{}
}
