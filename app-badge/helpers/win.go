package helpers

import (
	"golang.org/x/sys/windows"
	//"log"
	"syscall"
	"unsafe"
)

var (
	u32                       = windows.NewLazySystemDLL("user32.dll")
	k32                       = windows.NewLazySystemDLL("kernel32.dll")
	pGetWindowTitle           = u32.NewProc("GetWindowTextW")
	pSetWindowTitle           = u32.NewProc("SetWindowTextW")
	pIsWindow                 = u32.NewProc("IsWindow")
	pIsIconic                 = u32.NewProc("IsIconic")
	pIsWindowVisible          = u32.NewProc("IsWindowVisible")
	pGetParent                = u32.NewProc("GetParent")
	pGetCurrentProcessId      = k32.NewProc("GetCurrentProcessId")
	pGetDesktopWindow         = u32.NewProc("GetDesktopWindow")
	pFindWindowEx             = u32.NewProc("FindWindowExW")
	pGetWindowThreadProcessId = u32.NewProc("GetWindowThreadProcessId")
)

type (
	HANDLE uintptr
	HWND   HANDLE
	LPARAM uintptr
	DWORD  uint32
)

func GetWindowText(hwnd uintptr) string {
	textLen := 32

	buf := make([]uint16, textLen)
	pGetWindowTitle.Call(
		hwnd,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func GetWindowHandle() (result uintptr) {
	result = 0
	var prevWindow uintptr = 0
	processId, _, err := pGetCurrentProcessId.Call()
	if processId == 0 {
		//log.Printf("Cannot get current process id: %v", err)
		return
	}
	//log.Printf("Current process ID: %v", processId)
	desktopWindow, err := GetDesktopWindow()
	if err != nil {
		//log.Printf("Desktop Window error: %s", err)
		return
	}
	//log.Printf("Desktop Window handle: %v", desktopWindow)

	for i := 0; i < 2000; i++ {
		nextWindow, _, _ := pFindWindowEx.Call(uintptr(desktopWindow), prevWindow, 0, 0)
		if nextWindow == 0 {
			//log.Printf("NextWindow error: %s", err)
			break
		}
		var cpid uintptr
		r1, _, _ := pGetWindowThreadProcessId.Call(nextWindow, uintptr(unsafe.Pointer(&cpid)))
		if r1 == 0 {
			//log.Printf("Cannot get process id of %v", nextWindow)
			break
		}
		//log.Printf("R1, ProcessId: %v, %v", r1, cpid)
		if cpid == processId {
			windowText := GetWindowText(nextWindow)
			//log.Printf("FOUND: %v, %s", cpid, windowText)
			//isw, _, _ := pIsWindow.Call(nextWindow)
			//log.Printf("Is Window: %v", isw)
			parentHandle, _, _ := pGetParent.Call(nextWindow)
			//log.Printf("Parent: %v", parentHandle)
			isVisible, _, _ := pIsWindowVisible.Call(nextWindow)
			//log.Printf("Is Visible: %v", isVisible)
			//isIconic, _, _ := pIsIconic.Call(nextWindow)
			//log.Printf("Is Iconic: %v", isIconic)
			if (parentHandle == 0) && (windowText != "") && isVisible != 0 {
				//log.Printf("Proper Window: %s", windowText)
				return nextWindow
			}
		}
		prevWindow = nextWindow
	}
	return
}

func GetDesktopWindow() (h windows.Handle, err error) {
	wh, _, err := pGetDesktopWindow.Call()
	return windows.Handle(wh), nil
}
