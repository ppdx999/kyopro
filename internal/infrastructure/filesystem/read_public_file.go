package filesystem

import "os"

func (fs *FileSystem) ReadPublicFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
