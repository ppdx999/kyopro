package filesystem

import "os"

func (fs *FileSystem) WriteSecretFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0600)
}
