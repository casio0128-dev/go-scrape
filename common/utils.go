package common

import (
	"runtime"
	"strings"
)

func IsWindowsOS() bool {
	return strings.EqualFold(runtime.GOOS, "windows")
}

func IsMacOS() bool {
	return strings.EqualFold(runtime.GOOS, "darwin")
}

func IsLinuxOS() bool {
	return strings.EqualFold(runtime.GOOS, "linux")
}
