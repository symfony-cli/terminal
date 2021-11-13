//go:build !windows
// +build !windows

package terminal

func HasNativeColorSupport(stream interface{}) bool {
	return IsTerminal(stream)
}
