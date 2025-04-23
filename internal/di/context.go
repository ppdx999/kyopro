package di

import "github.com/ppdx999/kyopro/internal/context"

func CurrentProblemLoaderImpl() *context.CurrentProblemLoaderImpl {
	var WorkspaceGetter = func() context.GetWd {
		return FsImpl()
	}

	return context.NewCurrentProblemLoaderImpl(
		WorkspaceGetter(),
	)
}
