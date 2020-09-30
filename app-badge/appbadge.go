package appbadge

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type AppBadge interface {
	SetBadge(value int32) error
	ClearBadge() error
}

var ErrUnsupportedPlatform = errors.New("Unsupported platform: " + runtime.GOOS)
var ErrNotImplemented = errors.New("Not implemented for platform: " + runtime.GOOS)
var Api AppBadge = &AppBadgeFallback{}

type AppBadgeFallback struct{}

func (*AppBadgeFallback) SetBadge(value int32) error {
	return ErrUnsupportedPlatform
}

func (*AppBadgeFallback) ClearBadge() error {
	return ErrUnsupportedPlatform
}

func formatWindowTitle(current string, badgeValue int32) string {
	t := current
	if open := strings.Index(current, "("); open == 0 {
		if close := strings.Index(current, ") "); close > 0 {
			t = current[close+2:]
		}
	}
	if badgeValue > 0 {
		if badgeValue > 99 {
			return fmt.Sprintf("(99+) %v", t)
		} else {
			return fmt.Sprintf("(%v) %v", strconv.Itoa(int(badgeValue)), t)
		}
	}
	return t
}
