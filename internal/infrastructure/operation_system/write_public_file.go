package operation_system

import "os"

func (o *OperationSystem) WritePublicFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
