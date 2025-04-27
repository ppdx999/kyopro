package di

import "github.com/ppdx999/kyopro/internal/domain/service/testcase"

func TestCaseGetter() testcase.TestCasesGetter {
	return Atcoder()
}

func TestCaseSaver() testcase.TestCaseSaver {
	var WorkspaceGetter = func() testcase.GetWd {
		return FileSystem()
	}
	var PublicDirMaker = func() testcase.PublicDirMaker {
		return FileSystem()
	}
	var PublicFileWriter = func() testcase.PublicFileWriter {
		return FileSystem()
	}

	return testcase.NewTestCaseFsSaver(
		WorkspaceGetter(),
		PublicDirMaker(),
		PublicFileWriter(),
	)
}
