package helpers

import (
	"golang.org/x/sys/windows"
)

var (
	// dlls
	u32 = windows.NewLazySystemDLL("user32.dll")

	// user32 functions
	pGetDesktopWindow = u32.NewProc("GetDesktopWindow")
)

func GetWindowHandle() (h windows.Handle) {
	return
}

func GetDesktopWindow() (h windows.Handle) {
	return
}
