package di

import (
	"github.com/ppdx999/kyopro/internal/fs"
)

func InitializeFsImpl() *fs.FsImpl {
	return fs.NewFsImpl()
}
