package browser

import "fmt"

func SizeBy(width, height int) string {
	return fmt.Sprintf("--window-size=%d,%d", width, height)
}

func UserDataBy(userDataDirectory string) string {
	return fmt.Sprintf("--user-data-dir=%s", userDataDirectory)
}

func IsHeadless() string {
	return "--headless"
}
