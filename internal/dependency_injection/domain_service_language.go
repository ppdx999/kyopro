package di

import (
	"github.com/ppdx999/kyopro/internal/domain/service/language"
	repository_language "github.com/ppdx999/kyopro/internal/infrastructure/repository/language"
)

func LanguageDetector() language.LanguageDetector {
	var AllLanguagesFetcher = func() language.AllLanguagesFetcher {
		return repository_language.NewRepositoryLanguage()
	}

	var FileExister = func() language.FileExister {
		return OperationSystem()
	}

	return language.NewLanguageDetector(
		AllLanguagesFetcher(),
		FileExister(),
	)
}

func LanguageTestCaseRunner() language.LanguageTestCaseRunner {
	return language.NewLanguageTestCaseRunner()
}
