package di

import (
	repository_language "github.com/ppdx999/kyopro/internal/infrastructure/repository/language"
	"github.com/ppdx999/kyopro/internal/infrastructure/repository/language/language_resource"
)

func RepositoryLanguage() *repository_language.RepositoryLanguage {
	var OsCmdRunner = func() language_resource.OsCmdRunner {
		return OperationSystem()
	}
	var FileRemover = func() language_resource.FileRemover {
		return OperationSystem()
	}

	return repository_language.NewRepositoryLanguage(
		OsCmdRunner(),
		FileRemover(),
	)
}
