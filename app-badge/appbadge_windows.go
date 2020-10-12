package appbadge

import (
	"golang.org/x/sys/windows"
	"log"
	"syscall"
)

import (
	"fmt"
	"os"
	"unsafe"
)

var (
	u32                       = windows.NewLazySystemDLL("user32.dll")
	k32                       = windows.NewLazySystemDLL("kernel32.dll")
	pGetWindowTitle           = u32.NewProc("GetWindowTextW")
	pSetWindowTitle           = u32.NewProc("SetWindowTextW")
	pEnumWindows              = u32.NewProc("EnumThreadWindows")
	pGetParent                = u32.NewProc("GetParent")
	pCurrentProcess           = k32.NewProc("GetCurrentProcess")
	pGetWindowThreadProcessId = u32.NewProc("GetWindowThreadProcessId")
)

type AppBadgeWindows struct{}

type (
	HANDLE uintptr
	HWND   HANDLE
	LPARAM uintptr
	DWORD  uint32
)

func EnumWindows(tid uintptr, enumFunc uintptr, lparam uintptr) (err error) {
	r1, r2, e1 := pEnumWindows.Call(tid, enumFunc, lparam)
	log.Printf("Enumerated window: r1 = %v, r2 = %v, lparam = %v", r1, r2, lparam)
	if r1 == 0 {
		if e1 != nil {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func FindWindow() error {

	getWindow()

	tid := windows.GetCurrentThreadId()
	log.Printf("Thread id: %v", tid)
	cb := windows.NewCallback(func(h windows.Handle, p uintptr) uintptr {
		chw := HWND(h)
		wt := GetWindowText(chw)
		//if wt != "" {
		//s, _ := windows.UTF16PtrFromString(fmt.Sprintf("Counter: %d", 123123 ))
		//SetWindowText(chw, s)
		//}
		var pid uintptr
		gotthid, _, err := pGetWindowThreadProcessId.Call(uintptr(chw), uintptr(unsafe.Pointer(&pid)))
		if err != nil {
		}
		log.Printf("Gotthid: %v (%v), PID: %v", gotthid, h, pid)
		log.Printf("Window: %s (%v) ptr: %v", wt, chw, p)
		return 1 // continue enumeration
	})
	EnumWindows(uintptr(tid), cb, 0)
	return nil
}

func echoCP() {
	pid, _ := windows.GetCurrentProcess()
	log.Printf("Current process: %v", pid)
}

func getWindow() uintptr {

	proc := u32.NewProc("GetForegroundWindow")
	hwnd, _, _ := proc.Call()
	hwnd2, herr := windows.GetCurrentProcess()
	chw := HWND(hwnd)
	log.Printf("Foreground window: %s (%v)", GetWindowText(chw), chw)

	iswow := k32.NewProc("IsWow64Process")

	//hwnd2, _, herr := syscall.Syscall(pCurrentProcess.Addr(), 0, 0, 0, 0)
	if herr != nil {
		log.Printf("Hwnd err: %s", herr)
	}
	log.Printf("GetCurrentProcess: %v (%v)", hwnd2, HWND(hwnd2))

	var result bool
	z1, z2, res := iswow.Call(uintptr(hwnd2), uintptr(unsafe.Pointer(&result)))

	log.Printf("IsWow64: %v (%v) ((%v)) %v", res, z1, z2, result)

	//
	//hwndParent, _, errp := pGetParent.Call(uintptr(hwnd))
	//
	//if errp != nil {
	//	log.Printf("Error parent: %v", errp)
	//}
	//
	//log.Printf("HwndParent: %v", hwndParent)

	return hwnd
}

func GetWindowText(hwnd HWND) string {
	textLen := 32

	buf := make([]uint16, textLen)
	pGetWindowTitle.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func SetWindowText(hwnd HWND, s *uint16) string {
	textLen := 32

	buf := make([]uint16, textLen)
	pSetWindowTitle.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func (*AppBadgeWindows) SetBadge(value int32) error {
	log.Printf("TEST: %d", value)

	s, err := windows.UTF16PtrFromString(fmt.Sprintf("Counter: %d", value))
	log.Printf("String Ptr: %v", s)
	if err != nil {
		log.Println("ERR: " + err.Error())
	}
	//log.Println("SET: " + SetWindowText(HWND(getWindow()), s))

	FindWindow()

	return nil
}

func (*AppBadgeWindows) ClearBadge() error {
	return ErrNotImplemented
}

func init() {
	f, err := os.OpenFile("log.txt", os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("Cannot open logfilw")
		os.Exit(666)
	}
	log.SetOutput(f)

	Api = &AppBadgeWindows{}
}
