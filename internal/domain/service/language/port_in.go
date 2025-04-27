package language

import "github.com/ppdx999/kyopro/internal/domain/model"

type LanguageDetector interface {
	DetectLanguage() (*model.Language, error)
}

func NewLanguageDetector(
	allLanguagesFetcher AllLanguagesFetcher,
	fileExister FileExister,
) LanguageDetector {
	return &detector{
		allLanguagesFetcher: allLanguagesFetcher,
		fileExister:         fileExister,
	}
}

type LanguageTestCaseRunner interface {
	Run(l *model.Language, tc *model.TestCase) ([]byte, []byte, error)
}

func NewLanguageTestCaseRunner() LanguageTestCaseRunner {
	return &testcaseRunner{}
}
