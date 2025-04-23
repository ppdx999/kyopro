package downlaod

import "github.com/ppdx999/kyopro/internal/model"

type CurrentProblemLoader interface {
	LoadCurrentProblem() (*model.Problem, error)
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
