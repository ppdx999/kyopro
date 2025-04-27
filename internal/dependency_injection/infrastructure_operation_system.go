package di

import "github.com/ppdx999/kyopro/internal/infrastructure/operation_system"

func OperationSystem() *operation_system.OperationSystem {
	return operation_system.NewOperationSystem()
}
