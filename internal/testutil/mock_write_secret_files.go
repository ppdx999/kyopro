package testutil

type MockWriteSecretFile struct {
	nCalled      int
	CreatedFiles map[string][]byte
	Errs         []error
}

func (m *MockWriteSecretFile) WriteSecretFile(path string, data []byte) error {
	err := m.Errs[m.nCalled]
	m.nCalled++
	if err != nil {
		return err
	}
	if m.CreatedFiles == nil {
		m.CreatedFiles = make(map[string][]byte)
	}
	m.CreatedFiles[path] = data
	return nil
}
