package di

import "github.com/ppdx999/kyopro/internal/domain/service/testcase"

func TestCaseGetter() testcase.TestCasesGetter {
	return Atcoder()
}

func TestCaseSaver() testcase.TestCaseSaver {
	var WorkspaceGetter = func() testcase.GetWd {
		return OperationSystem()
	}
	var PublicDirMaker = func() testcase.PublicDirMaker {
		return OperationSystem()
	}
	var PublicFileWriter = func() testcase.PublicFileWriter {
		return OperationSystem()
	}

	return testcase.NewTestCaseFsSaver(
		WorkspaceGetter(),
		PublicDirMaker(),
		PublicFileWriter(),
	)
}
