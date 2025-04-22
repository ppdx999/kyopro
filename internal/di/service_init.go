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
		return init_.NewGetWorkspaceImpl(GetWd())
	}

	var MakeProblemDir = func() init_.MakeProblemDir {
		return init_.NewMakeProblemDirImpl(
			MakePublicDir(),
			GetWorkspace(),
		)
	}

	return init_.NewInitServiceImpl(
		GetProblemIds(),
		MakeProblemDir(),
	)

}
