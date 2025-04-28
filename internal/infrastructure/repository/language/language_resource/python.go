package language_resource

import "github.com/ppdx999/kyopro/internal/domain/model"

type Python struct {
	runner OsCmdRunner
	reader FileReader
}

func (py *Python) Name() string {
	return "python"
}

func (py *Python) SourceCode() *model.SourceCode {
	path := "main.py"
	code, _ := py.reader.ReadFileIfExist(path)
	return &model.SourceCode{Path: path, Code: code}
}

func (py *Python) Build(p *model.Pipeline) error {
	return nil
}

func (py *Python) Run(p *model.Pipeline) error {
	return py.runner.Run([]string{"python", py.SourceCode().Path}, p)
}

func (py *Python) Clean(p *model.Pipeline) error {
	return nil
}
