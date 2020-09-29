package appbadge

/*
#include "platform/darwin.h"
*/
import "C"

type AppBadgeDarwin struct{}

func (*AppBadgeDarwin) SetBadge(value int) error {
	r1 := SetWindowTitle(formatWindowTitle(GetWindowTitle(), value))
	if r1 != 0 {
		return newError("Failed to set window title")
	}
	r2 := SetBadgeValue(value)
	if r2 != 0 {
		return newError("Failed to set app badge value")
	}
	return nil
}

func (*AppBadgeDarwin) ClearBadge() error {
	return SetBadge(0)
}

func init() {
	api = &AppBadgeDarwin{}
}
