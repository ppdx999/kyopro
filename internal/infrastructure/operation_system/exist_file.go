package operation_system

import "os"

func (o *OperationSystem) ExistFile(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
