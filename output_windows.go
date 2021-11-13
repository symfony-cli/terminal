package terminal

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
)

func defaultOutputs() (io.Writer, io.Writer) {
	if HasPosixColorSupport() {
		return os.Stdout, os.Stderr
	}

	return colorable.NewColorableStdout(), colorable.NewColorableStderr()
}
