package operation_system

import "os"

func (o *OperationSystem) ReadSecretFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
