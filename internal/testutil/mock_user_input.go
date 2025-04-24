package testutil

type MockUserInput struct {
	Input string
	Err   error
}

func (m *MockUserInput) UserInput() (string, error) {
	if m.Err != nil {
		return "", m.Err
	}
	return m.Input, nil
}
