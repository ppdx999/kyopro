package language

import "github.com/ppdx999/kyopro/internal/domain/model"

type AllLanguagesFetcher interface {
	FetchAllLanguages() ([]*model.Language, error)
}

type FileExister interface {
	ExistFile(path string) bool
}
