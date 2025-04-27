package di

import (
	"github.com/ppdx999/kyopro/internal/infrastructure/filesystem"
)

func FileSystem() *filesystem.FileSystem {
	return filesystem.NewFileSystem()
}
