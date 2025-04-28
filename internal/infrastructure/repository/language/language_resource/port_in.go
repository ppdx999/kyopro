package language_resource

import "github.com/ppdx999/kyopro/internal/domain/model"

func NewCpp(
	runner OsCmdRunner,
	rm FileRemover,
) model.Language {
	return &Cpp{
		runner: runner,
		rm:     rm,
	}
}

func NewPython(runner OsCmdRunner) model.Language {
	return &Python{
		runner: runner,
	}
}
