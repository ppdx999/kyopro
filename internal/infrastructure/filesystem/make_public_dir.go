package filesystem

import "os"

func (fs *FileSystem) MakePublicDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	return nil
}
