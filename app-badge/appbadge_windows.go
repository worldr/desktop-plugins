package appbadge

import (
	"fmt"
	"github.com/worldr/desktop-plugins/app-badge/helpers"
	"golang.org/x/sys/windows"
	"log"
	"regexp"
	"syscall"
)

import (
	"os"
	"unsafe"
)

var (
	u32             = windows.NewLazySystemDLL("user32.dll")
	pSetWindowTitle = u32.NewProc("SetWindowTextW")
	pFlash          = u32.NewProc("FlashWindow")
)

type AppBadgeWindows struct{}

type (
	HANDLE  uintptr
	HWND    HANDLE
	LPARAM  uintptr
	DWORD   uint32
	handler struct {
		is *bool
	}
)

func SetWindowText(hwnd helpers.HWND, txt string) string {
	s, _ := windows.UTF16PtrFromString(txt)
	textLen := len(txt)

	if textLen == 0 {
		log.Printf("Zero length string: %s", txt)
		return ""
	}

	buf := make([]uint16, textLen)
	pSetWindowTitle.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func (me *AppBadgeWindows) SetBadge(value int32) error {
	handle := helpers.GetWindowHandle()
	if handle == 0 {
		return nil
	}
	currentText := helpers.GetWindowText(helpers.HWND(handle))
	r, _ := regexp.Compile("^([^ ]+).*$")
	if value != 0 {
		SetWindowText(helpers.HWND(handle), fmt.Sprintf(r.ReplaceAllString(currentText, "$1")+" (%v)", value))
		f := handler{is: func(b bool) *bool { return &b }(true)}
		pFlash.Call(uintptr(handle), uintptr(unsafe.Pointer(&f.is)))
	}

	return nil
}

func (me *AppBadgeWindows) ClearBadge() error {
	handle := helpers.GetWindowHandle()
	currentText := helpers.GetWindowText(helpers.HWND(handle))
	r, _ := regexp.Compile("^([^ ]+).*$")
	SetWindowText(helpers.HWND(handle), r.ReplaceAllString(currentText, "$1"))
	return nil
}

func init() {
	f, err := os.OpenFile("log.txt", os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("Cannot open logfilw")
		os.Exit(1)
	}
	log.SetOutput(f)

	Api = &AppBadgeWindows{}
}
