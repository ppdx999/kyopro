package testutil

type MockExistFile struct {
	Exist bool
}

func (m *MockExistFile) ExistFile(path string) bool {
	return m.Exist
}
