package repository_language

import (
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/infrastructure/repository/language/language_resource"
)

func (r *RepositoryLanguage) FetchAllLanguages() ([]model.Language, error) {
	var langs = []model.Language{
		language_resource.NewCpp(r.osCmdRunner, r.fileRemover, r.fileReader),
		language_resource.NewPython(r.osCmdRunner, r.fileReader),
	}
	return langs, nil
}
