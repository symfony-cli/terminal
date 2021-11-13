package terminal

import "os"

func HasPosixColorSupport() bool {
	return os.Getenv("ANSICON") != "" || os.Getenv("ConEmuANSI") == "ON" || os.Getenv("TERM") == "xterm" || os.Getenv("SHLVL") != ""
}
