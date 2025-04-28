package language_resource

import "github.com/ppdx999/kyopro/internal/domain/model"

type Python struct {
	runner OsCmdRunner
}

func (py *Python) Name() string {
	return "python"
}

func (py *Python) MainFile() string {
	return "main.py"
}

func (py *Python) Build(p *model.Pipeline) error {
	return nil
}

func (py *Python) Run(p *model.Pipeline) error {
	return py.runner.Run([]string{"python", py.MainFile()}, p)
}

func (py *Python) Clean(p *model.Pipeline) error {
	return nil
}
