package common

const (
	DATE_REGEXP = `^[0-9]{4}\/((0[1-9]|[1-9])|1[0-2])\/((0[1-9]|[1-9])|[12][0-9]|3[01])日?$`
	TIME_REGEXP = `^([0-1]\d|2[0-3]|\d):[0-5]\d?$`

	DATE_SEPARATER = "/"
	TIME_SEPARATER = ":"

	WINDOWS_ENV_SEPARATER = ";"
	LINUX_ENV_SEPARATER   = ":"
	MACOS_ENV_SEPARATER   = ":"
)
