package helpers

import (
	"golang.org/x/sys/windows"
	"log"
	"syscall"
	"unsafe"
)

var (
	u32                       = windows.NewLazySystemDLL("user32.dll")
	k32                       = windows.NewLazySystemDLL("kernel32.dll")
	pGetWindowTitle           = u32.NewProc("GetWindowTextW")
	pSetWindowTitle           = u32.NewProc("SetWindowTextW")
	pIsWindow                 = u32.NewProc("IsWindow")
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

func GetWindowText(hwnd HWND) string {
	textLen := 32

	buf := make([]uint16, textLen)
	pGetWindowTitle.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func GetWindowHandle() (result uintptr) {
	result = 0
	var prevWindow HWND = 0
	processId, _, err := pGetCurrentProcessId.Call()
	if processId == 0 {
		log.Printf("Cannot get current process id: %v", err)
		return
	}
	log.Printf("Current process ID: %v", processId)
	desktopWindow, err := GetDesktopWindow()
	if err != nil {
		log.Printf("Desktop Window error: %s", err)
		return
	}
	log.Printf("Desktop Window handle: %v", desktopWindow)

	for i := 0; i < 100; i++ {
		nextWindow, _, _ := pFindWindowEx.Call(uintptr(desktopWindow), uintptr(prevWindow), 0, 0)
		if nextWindow == 0 {
			log.Printf("NextWindow error: %s", err)
			break
		}
		//log.Printf("NextWindow handle: %v (%s)", nextWindow, GetWindowText(HWND(nextWindow)))
		var cpid uintptr
		r1, _, _ := pGetWindowThreadProcessId.Call(nextWindow, uintptr(unsafe.Pointer(&cpid)))
		if r1 == 0 {
			log.Printf("Cannot get process id of %v", nextWindow)
			break
		}
		//log.Printf("R1, ProcessId: %v, %v", r1, cpid)
		if cpid == processId {
			windowText := GetWindowText(HWND(nextWindow))
			log.Printf("FOUND: %v, %s", cpid, windowText)
			isw, _, _ := pIsWindow.Call(uintptr(nextWindow))
			log.Printf("Is Window: %v", isw)
			parentHandle, _, _ := pGetParent.Call(uintptr(nextWindow))
			log.Printf("Parent: %v", parentHandle)
			if (parentHandle == 0) && (windowText != "") {
				log.Printf("Proper Window: %s", windowText)
				return nextWindow
			}
		}
		prevWindow = HWND(nextWindow)
	}
	return
}

func GetDesktopWindow() (h windows.Handle, err error) {
	wh, _, err := pGetDesktopWindow.Call()
	//if err != nil {
	//	log.Printf("Cannot get desktop window: %s", err)
	//}
	return windows.Handle(wh), nil
}
