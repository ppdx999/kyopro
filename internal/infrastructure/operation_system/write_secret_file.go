package operation_system

import "os"

func (o *OperationSystem) WriteSecretFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0600)
}
