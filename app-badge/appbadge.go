package appbadge

import (
	"errors"
	"runtime"
)

type AppBadge interface {
	SetBadge(value int) error
	ClearBadge() error
}

var ErrUnsupportedPlatform = errors.New("Unsupported platform: " + runtime.GOOS)
var ErrNotImplemented = errors.New("Not implemented for platform: " + runtime.GOOS)
var api AppBadge = &AppBadgeFallback{}

type AppBadgeFallback struct{}

func (*AppBadgeFallback) SetBadge(value int) error {
	return ErrUnsupportedPlatform
}

func (*AppBadgeFallback) ClearBadge() error {
	return ErrUnsupportedPlatform
}
