package session_test

type MockSessionPath struct {
	path string
	err  error
}

func (m *MockSessionPath) SessionPath() (string, error) {
	return m.path, m.err
}
