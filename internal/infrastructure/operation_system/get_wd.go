package operation_system

import "os"

func (o *OperationSystem) GetWd() (string, error) {
	return os.Getwd()
}
