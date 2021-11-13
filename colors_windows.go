package terminal

import (
	"strconv"

	"golang.org/x/sys/windows/registry"
)

func HasNativeColorSupport(stream interface{}) bool {
	return isWindows10orMore()
}

func isWindows10orMore() bool {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()

	cv, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		return false
	}
	version, err := strconv.ParseFloat(cv, 32)
	if err != nil {
		return false
	}

	// 6.1	Windows 7 / Windows Server 2008 R2
	// 6.2	Windows 8 / Windows Server 2012
	// 6.3	Windows 8.1 / Windows Server 2012 R2
	// 10.0	Windows 10
	// But some Windows 10 systems (at least mine) return 6.3 ...
	return version >= 6.3
}
