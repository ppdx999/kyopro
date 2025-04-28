package language_resource

import "github.com/ppdx999/kyopro/internal/domain/model"

func NewCpp(
	runner OsCmdRunner,
	rm FileRemover,
	reader FileReader,
) model.Language {
	return &Cpp{
		runner: runner,
		rm:     rm,
		reader: reader,
	}
}

func NewPython(runner OsCmdRunner, reader FileReader) model.Language {
	return &Python{
		runner: runner,
		reader: reader,
	}
}
