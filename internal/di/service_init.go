package di

import init_ "github.com/ppdx999/kyopro/internal/service/init"

func InitService() init_.InitService {
	var GetWd = func() init_.GetWd {
		return FsImpl()
	}

	var MakePublicDir = func() init_.MakePublicDir {
		return FsImpl()
	}

	var GetProblemIds = func() init_.GetProblemIds {
		return Atcoder()
	}

	var GetWorkspace = func() init_.GetWorkspace {
		getWd := GetWd()
		return init_.NewGetWorkspaceImpl(getWd)
	}

	var MakeProblemDir = func() init_.MakeProblemDir {
		makePublicDir := MakePublicDir()
		getWorkspace := GetWorkspace()
		return init_.NewMakeProblemDirImpl(makePublicDir, getWorkspace)
	}

	getProblemIds := GetProblemIds()
	makeProblemDir := MakeProblemDir()
	return init_.NewInitServiceImpl(getProblemIds, makeProblemDir)
}
