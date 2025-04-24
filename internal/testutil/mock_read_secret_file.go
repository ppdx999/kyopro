package testutil

type MockReadSecretFile struct {
	nCalled   int
	ReadFiles []string
	Datas     [][]byte
	Errs      []error
}

func (m *MockReadSecretFile) ReadSecretFile(path string) ([]byte, error) {
	err := m.Errs[m.nCalled]
	data := m.Datas[m.nCalled]
	m.nCalled++

	if err != nil {
		return nil, err
	}

	return data, nil
}
