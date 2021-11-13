//go:build !windows
// +build !windows

package terminal

import (
	"io"
	"os"
)

func defaultOutputs() (io.Writer, io.Writer) {
	return os.Stdout, os.Stderr
}
