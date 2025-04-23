package fs

import "os"

type FsImpl struct{}

func NewFsImpl() *FsImpl {
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

func (fs *FsImpl) ExistFile(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func (fs *FsImpl) ReadSecretFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (fs *FsImpl) WriteSecretFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0600)
}

func (fs *FsImpl) ReadPublicFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (fs *FsImpl) WritePublicFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
