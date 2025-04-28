package language_resource

import (
	"github.com/ppdx999/kyopro/internal/domain/model"
)

type Cpp struct {
	runner OsCmdRunner
	rm     FileRemover
	reader FileReader
}

func (c *Cpp) Name() string {
	return "c++"
}

func (c *Cpp) SourceCode() *model.SourceCode {
	path := "main.cpp"
	code, _ := c.reader.ReadFileIfExist(path)
	return &model.SourceCode{Path: path, Code: code}
}

func (c *Cpp) Build(p *model.Pipeline) error {
	return c.runner.Run([]string{"g++", c.SourceCode().Path}, p)
}

func (c *Cpp) Run(p *model.Pipeline) error {
	return c.runner.Run([]string{"./a.out"}, p)
}

func (c *Cpp) Clean(p *model.Pipeline) error {
	return c.rm.Remove("a.out")
}
