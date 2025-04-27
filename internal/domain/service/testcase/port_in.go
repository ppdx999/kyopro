package testcase

import "github.com/ppdx999/kyopro/internal/domain/model"

type TestCaseCurrentLoader interface {
	LoadCurrentTestCases() ([]*model.TestCase, error)
}

func NewTestCaseCurrentLoader(
	wdGetter WdGetter,
	childFileNamesGetter ChildFileNamesGetter,
	fileReader PublicFileReader,
) TestCaseCurrentLoader {
	return &testCaseCurrentLoader{
		wdGetter:             wdGetter,
		childFileNamesGetter: childFileNamesGetter,
		publicFileReader:     fileReader,
	}
}

type TestCasesGetter interface {
	GetTestCases(
		contestId model.ContestId,
		problemId model.ProblemId,
	) (
		[]*model.TestCase,
		error,
	)
}

type TestCaseSaver interface {
	SaveTestCase(t *model.TestCase) error
}

func NewTestCaseSaver(wd WdGetter, dirMaker PublicDirMaker, fileWriter PublicFileWriter) TestCaseSaver {
	return &testCaseSaver{
		wd:         wd,
		dirMaker:   dirMaker,
		fileWriter: fileWriter,
	}
}
