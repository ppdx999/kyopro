package testcase

import "github.com/ppdx999/kyopro/internal/domain/model"

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

func NewTestCaseSaver(wd GetWd, dirMaker PublicDirMaker, fileWriter PublicFileWriter) TestCaseSaver {
	return &testCaseSaver{
		wd:         wd,
		dirMaker:   dirMaker,
		fileWriter: fileWriter,
	}
}
