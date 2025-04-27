package di

import "github.com/ppdx999/kyopro/internal/domain/service/testcase"

func TestCaseCurrentLoader() testcase.TestCaseCurrentLoader {
	var WdGetter = func() testcase.WdGetter {
		return OperationSystem()
	}
	var ChildFileNamesGetter = func() testcase.ChildFileNamesGetter {
		return OperationSystem()
	}
	var PublicFileReader = func() testcase.PublicFileReader {
		return OperationSystem()
	}

	return testcase.NewTestCaseCurrentLoader(
		WdGetter(),
		ChildFileNamesGetter(),
		PublicFileReader(),
	)
}

func TestCaseGetter() testcase.TestCasesGetter {
	return Atcoder()
}

func TestCaseSaver() testcase.TestCaseSaver {
	var WorkspaceGetter = func() testcase.WdGetter {
		return OperationSystem()
	}
	var PublicDirMaker = func() testcase.PublicDirMaker {
		return OperationSystem()
	}
	var PublicFileWriter = func() testcase.PublicFileWriter {
		return OperationSystem()
	}

	return testcase.NewTestCaseSaver(
		WorkspaceGetter(),
		PublicDirMaker(),
		PublicFileWriter(),
	)
}
