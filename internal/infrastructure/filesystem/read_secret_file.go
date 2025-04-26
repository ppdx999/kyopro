package filesystem

import "os"

func (fs *FileSystem) ReadSecretFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
