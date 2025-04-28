package language

import (
	"errors"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

type detector struct {
	allLanguagesFetcher AllLanguagesFetcher
	fileExister         FileExister
}

func (d *detector) DetectLanguage() (model.Language, error) {
	langs, err := d.allLanguagesFetcher.FetchAllLanguages()
	if err != nil {
		return nil, err
	}
	for _, lang := range langs {
		if d.fileExister.ExistFile(lang.SourceCode().Path) {
			return lang, nil
		}
	}

	return nil, errors.New("language not found")
}
