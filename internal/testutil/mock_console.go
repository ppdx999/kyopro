package testutil

import (
	"bytes"
	"io"
)

type MockConsole struct {
	inBuf  *bytes.Buffer
	outBuf *bytes.Buffer
	errBuf *bytes.Buffer
}

func NewMockConsole() *MockConsole {
	return &MockConsole{
		inBuf:  &bytes.Buffer{},
		outBuf: &bytes.Buffer{},
		errBuf: &bytes.Buffer{},
	}
}

func (m *MockConsole) In() io.Reader                  { return m.inBuf }
func (m *MockConsole) Out() io.Writer                 { return m.outBuf }
func (m *MockConsole) Err() io.Writer                 { return m.errBuf }
func (m *MockConsole) SetIn(r io.Reader)              {}
func (m *MockConsole) SetOut(w io.Writer)             {}
func (m *MockConsole) SetErr(w io.Writer)             {}
func (m *MockConsole) ReadIn(p []byte) (int, error)   { return m.inBuf.Read(p) }
func (m *MockConsole) WriteOut(p []byte) (int, error) { return m.outBuf.Write(p) }
func (m *MockConsole) WriteErr(p []byte) (int, error) { return m.errBuf.Write(p) }
