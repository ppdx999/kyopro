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
