package terminal

import (
	"io"
)

const (
	ErrNotATTY                  = termErr(0)
	defaultWidth, defaultHeight = 80, 20
)

type FdHolder interface {
	Fd() uintptr
}

func IsTerminal(stream interface{}) bool {
	output, streamIsFile := stream.(FdHolder)
	return streamIsFile && IsTTY(output.Fd())
}

type FdReader interface {
	io.Reader
	FdHolder
}

type termErr int

func (e termErr) Error() string {
	switch e {
	case 0:
		return "not a TTY"
	}
	return "undefined terminal error"
}
