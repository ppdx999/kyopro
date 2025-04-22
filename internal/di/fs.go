package di

import (
	"github.com/ppdx999/kyopro/internal/fs"
)

func FsImpl() *fs.FsImpl {
	return fs.NewFsImpl()
}
