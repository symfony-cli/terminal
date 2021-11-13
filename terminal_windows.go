package terminal

import (
	"os"

	isattypkg "github.com/mattn/go-isatty"
	"github.com/pkg/errors"
	"golang.org/x/sys/windows"
	"golang.org/x/term"
)

func IsTTY(fd uintptr) bool {
	return isattypkg.IsTerminal(fd) || IsCygwinTTY(fd)
}

func IsCygwinTTY(fd uintptr) bool {
	return isattypkg.IsCygwinTerminal(fd)
}

// term.GetSize does not work with MinGW
// But this implementation does
func GetSize() (width, height int) {
	fd := int(os.Stdout.Fd())

	if IsCygwinTTY(os.Stdout.Fd()) {
		h, err := openWindowsHandle(fd)
		if err != nil {
			return defaultWidth, defaultHeight
		}
		defer windows.CloseHandle(h)
		fd = int(h)
	}

	if w, h, err := term.GetSize(fd); err == nil {
		return w, h
	}

	return defaultWidth, defaultHeight
}

func MakeRaw(fd int) (*term.State, error) {
	if IsCygwinTTY(uintptr(fd)) {
		return nil, ErrNotATTY
	}

	return term.MakeRaw(fd)
}

func Restore(fd int, state *term.State) error {
	if IsCygwinTTY(uintptr(fd)) {
		return ErrNotATTY
	}

	return term.Restore(fd, state)
}

//https://rosettacode.org/wiki/Terminal_control/Dimensions#Windows
func openWindowsHandle(fd int) (windows.Handle, error) {
	path := ""
	switch fd {
	case int(os.Stdin.Fd()):
		path = "CONIN$"
	case int(os.Stdout.Fd()), int(os.Stderr.Fd()):
		path = "CONOUT$"
	default:
		return windows.Handle(0), errors.Errorf("unsupported handler %v", fd)
	}

	return windows.CreateFile(
		windows.StringToUTF16Ptr(path),
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		0,
		0,
	)
}
