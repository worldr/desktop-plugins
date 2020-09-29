package appbadge

/*
#include "platform/darwin.h"
*/
import "C"

type AppBadgeDarwin struct{}

func (*AppBadgeDarwin) SetBadge(value int) error {
	SetWindowTitle("Worldr")
	return ErrUnsupportedPlatform
}

func (*AppBadgeDarwin) ClearBadge() error {
	return ErrUnsupportedPlatform
}

func init() {
	api = &AppBadgeDarwin{}
}
