package appbadge

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "platforms/darwin.h"
*/
import "C"

type AppBadgeDarwin struct{}

func (*AppBadgeDarwin) SetBadge(value int) error {
	r1 := PlatformSetWindowTitle(formatWindowTitle(PlatformGetWindowTitle(), value))
	if r1 != 0 {
		return newError("Failed to set window title")
	}
	r2 := PlatformSetBadge(value)
	if r2 != 0 {
		return newError("Failed to set app badge value")
	}
	return nil
}

func (*AppBadgeDarwin) ClearBadge() error {
	return SetBadge(0)
}

func init() {
	Api = &AppBadgeDarwin{}
}
