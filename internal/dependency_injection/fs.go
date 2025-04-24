package di

import (
	fs "github.com/ppdx999/kyopro/internal/infrastructure/filesystem"
)

func FsImpl() *fs.FsImpl {
	return fs.NewFsImpl()
}
