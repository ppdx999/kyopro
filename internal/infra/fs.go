package infra

import "os"

type FsImpl struct{}

func NewFsImple() *FsImpl {
	return &FsImpl{}
}

func (fs *FsImpl) MakePublicDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	return nil
}

func (fs *FsImpl) GetWd() (string, error) {
	return os.Getwd()
}
