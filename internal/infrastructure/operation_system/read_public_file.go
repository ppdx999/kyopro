package operation_system

import "os"

func (o *OperationSystem) ReadPublicFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
