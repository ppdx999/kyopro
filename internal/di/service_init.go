package di

import init_ "github.com/ppdx999/kyopro/internal/service/init"

func InitializeGetWd() init_.GetWd {
	return InitializeFsImpl()
}

func InitializeMakePublicDir() init_.MakePublicDir {
	return InitializeFsImpl()
}

func InitializeGetProblemIds() init_.GetProblemIds {
	return InitializeAtcoder()
}

func InitializeGetWorkspace() init_.GetWorkspace {
	getWd := InitializeGetWd()
	return init_.NewGetWorkspaceImpl(getWd)
}

func InitializeMakeProblemDir() init_.MakeProblemDir {
	makePublicDir := InitializeMakePublicDir()
	getWorkspace := InitializeGetWorkspace()
	return init_.NewMakeProblemDirImpl(makePublicDir, getWorkspace)
}

func InitializeInitService() init_.InitService {
	getProblemIds := InitializeGetProblemIds()
	makeProblemDir := InitializeMakeProblemDir()
	return init_.NewInitServiceImpl(getProblemIds, makeProblemDir)
}
