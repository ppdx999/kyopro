package filesystem

import "os"

func (fs *FileSystem) ExistFile(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
