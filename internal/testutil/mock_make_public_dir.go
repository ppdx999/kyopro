package testutil

type MockMakePublicDir struct {
	nCalled     int
	Errs        []error
	CreatedDirs map[string]bool
}

func (m *MockMakePublicDir) MakePublicDir(path string) error {
	err := m.Errs[m.nCalled]
	m.nCalled++

	if err != nil {
		return err
	}

	if m.CreatedDirs == nil {
		m.CreatedDirs = make(map[string]bool)
	}

	m.CreatedDirs[path] = true
	return nil
}
