package model

type TestCaseID string

type TestCase struct {
	ID    TestCaseID
	Input []byte
	Want  []byte
}

func NewTestCase(id string) *TestCase {
	if id == "" {
		panic("test case id is empty")
	}

	return &TestCase{
		ID: TestCaseID(id),
	}
}
