package terminal

import (
	"os"
)

func IsInteractive(stream interface{}) bool {
	if !IsTerminal(stream) {
		return false
	}

	if IsCI() && os.Getenv("SHELL_INTERACTIVE") == "" {
		return false
	}

	return true
}
