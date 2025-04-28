package operation_system

import "os"

func (o *OperationSystem) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
