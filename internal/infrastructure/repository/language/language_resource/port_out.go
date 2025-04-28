package language_resource

import "github.com/ppdx999/kyopro/internal/domain/model"

type OsCmdRunner interface {
	Run(cmd []string, p *model.Pipeline) error
}

type FileRemover interface {
	Remove(path string) error
}

type FileExister interface {
	ExistFile(path string) bool
}

type FileReader interface {
	ReadFileIfExist(path string) ([]byte, error)
}
