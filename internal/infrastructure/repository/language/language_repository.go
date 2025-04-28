package repository_language

import "github.com/ppdx999/kyopro/internal/infrastructure/repository/language/language_resource"

type RepositoryLanguage struct {
	osCmdRunner language_resource.OsCmdRunner
	fileRemover language_resource.FileRemover
	fileReader  language_resource.FileReader
}

func NewRepositoryLanguage(
	osCmdRunner language_resource.OsCmdRunner,
	fileRemover language_resource.FileRemover,
	fileReader language_resource.FileReader,
) *RepositoryLanguage {
	return &RepositoryLanguage{
		osCmdRunner: osCmdRunner,
		fileRemover: fileRemover,
		fileReader:  fileReader,
	}
}
