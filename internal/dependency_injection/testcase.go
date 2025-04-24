package di

import "github.com/ppdx999/kyopro/internal/domain/service/testcase"

func TestCaseFsSaver() *testcase.TestCaseFsSaver {
	var WorkspaceGetter = func() testcase.GetWd {
		return FsImpl()
	}
	var PublicDirMaker = func() testcase.PublicDirMaker {
		return FsImpl()
	}
	var PublicFileWriter = func() testcase.PublicFileWriter {
		return FsImpl()
	}

	return testcase.NewTestCaseFsSaver(
		WorkspaceGetter(),
		PublicDirMaker(),
		PublicFileWriter(),
	)
}
