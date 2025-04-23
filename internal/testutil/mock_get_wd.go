package testutil

type MockGetWd struct {
	Wd  string
	Err error
}

func (m *MockGetWd) GetWd() (string, error) {
	if m.Err != nil {
		return "", m.Err
	}
	return m.Wd, nil
}
