package testutil

type MockHome struct {
	HomeDir string
	Err     error
}

func (m *MockHome) Home() (string, error) {
	if m.Err != nil {
		return "", m.Err
	}
	return m.HomeDir, nil
}
