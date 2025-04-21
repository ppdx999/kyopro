package console

import (
	"io"
	"os"
)

type Console interface {
	In() io.Reader
	Out() io.Writer
	Err() io.Writer
	SetIn(r io.Reader)
	SetOut(w io.Writer)
	SetErr(w io.Writer)
	ReadIn(p []byte) (n int, err error)
	WriteOut(p []byte) (n int, err error)
	WriteErr(p []byte) (n int, err error)
}

type ConsoleImpl struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

func NewConsoleImpl() *ConsoleImpl {
	return &ConsoleImpl{
		in:  os.Stdin,
		out: os.Stdout,
		err: os.Stderr,
	}
}

func (c *ConsoleImpl) In() io.Reader {
	return c.in
}

func (c *ConsoleImpl) Out() io.Writer {
	return c.out
}

func (c *ConsoleImpl) Err() io.Writer {
	return c.err
}

func (c *ConsoleImpl) SetIn(r io.Reader) {
	c.in = r
}

func (c *ConsoleImpl) SetOut(w io.Writer) {
	c.out = w
}

func (c *ConsoleImpl) SetErr(w io.Writer) {
	c.err = w
}

func (c *ConsoleImpl) ReadIn(p []byte) (n int, err error) {
	return c.in.Read(p)
}

func (c *ConsoleImpl) WriteOut(p []byte) (n int, err error) {
	return c.out.Write(p)
}

func (c *ConsoleImpl) WriteErr(p []byte) (n int, err error) {
	return c.err.Write(p)
}
