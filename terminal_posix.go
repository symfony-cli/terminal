//go:build !windows
// +build !windows

package terminal

import (
	"os"

	isattypkg "github.com/mattn/go-isatty"
	"golang.org/x/term"
)

var (
	IsTTY   = isattypkg.IsTerminal
	MakeRaw = term.MakeRaw
	Restore = term.Restore
)

func GetSize() (width, height int) {
	for _, f := range []*os.File{os.Stderr, os.Stdout, os.Stdin} {
		w, h, err := term.GetSize(int(f.Fd()))
		if err != nil {
			continue
		}

		if w > 0 && h > 0 {
			return w, h
		}
	}

	return defaultWidth, defaultHeight
}

func IsCygwinTTY(fd uintptr) bool {
	return false
}
