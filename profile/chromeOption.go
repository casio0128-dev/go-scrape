package profile

import "fmt"

func sizeBy(width, height int) string {
	return fmt.Sprintf("--window-size=%d,%d", width, height)
}

func userDataBy(userDataDirectory string) string {
	return fmt.Sprintf("--user-data-dir=%s", userDataDirectory)
}

func isHeadless() string {
	return "--headless"
}
