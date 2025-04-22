package di

import init_ "github.com/ppdx999/kyopro/internal/service/init"

func InitializeInitService() init_.InitService {
	var InitializeGetWd = func() init_.GetWd {
		return InitializeFsImpl()
	}

	var InitializeMakePublicDir = func() init_.MakePublicDir {
		return InitializeFsImpl()
	}

	var InitializeGetProblemIds = func() init_.GetProblemIds {
		return InitializeAtcoder()
	}

	var InitializeGetWorkspace = func() init_.GetWorkspace {
		getWd := InitializeGetWd()
		return init_.NewGetWorkspaceImpl(getWd)
	}

	var InitializeMakeProblemDir = func() init_.MakeProblemDir {
		makePublicDir := InitializeMakePublicDir()
		getWorkspace := InitializeGetWorkspace()
		return init_.NewMakeProblemDirImpl(makePublicDir, getWorkspace)
	}

	getProblemIds := InitializeGetProblemIds()
	makeProblemDir := InitializeMakeProblemDir()
	return init_.NewInitServiceImpl(getProblemIds, makeProblemDir)
}
