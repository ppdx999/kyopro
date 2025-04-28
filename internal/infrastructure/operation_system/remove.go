package operation_system

import "os"

func (o *OperationSystem) Remove(path string) error {
	return os.Remove(path)
}
