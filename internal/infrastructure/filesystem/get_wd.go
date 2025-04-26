package filesystem

import "os"

func (fs *FileSystem) GetWd() (string, error) {
	return os.Getwd()
}
