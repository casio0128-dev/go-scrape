package common

import (
	"runtime"
	"strconv"
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

func IsInt(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func GetKeys(m map[string]interface{}) []string {
	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	return keys
}
