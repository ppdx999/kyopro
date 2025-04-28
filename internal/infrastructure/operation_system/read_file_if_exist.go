package operation_system

func (o *OperationSystem) ReadFileIfExist(path string) ([]byte, error) {
	if o.ExistFile(path) {
		return o.ReadFile(path)
	} else {
		return nil, nil
	}
}
