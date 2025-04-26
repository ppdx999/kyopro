package filesystem

import "os"

func (fs *FileSystem) WritePublicFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
