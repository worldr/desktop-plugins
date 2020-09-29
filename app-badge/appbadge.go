package appbadge

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type AppBadge interface {
	SetBadge(value int) error
	ClearBadge() error
}

var ErrUnsupportedPlatform = errors.New("Unsupported platform: " + runtime.GOOS)
var ErrNotImplemented = errors.New("Not implemented for platform: " + runtime.GOOS)
var Api AppBadge = &AppBadgeFallback{}

type AppBadgeFallback struct{}

func (*AppBadgeFallback) SetBadge(value int) error {
	return ErrUnsupportedPlatform
}

func (*AppBadgeFallback) ClearBadge() error {
	return ErrUnsupportedPlatform
}

func formatWindowTitle(current string, badgeValue int) string {
	t := current
	if open := strings.Index(current, "("); open == 0 {
		if close := strings.Index(current, ") "); close > 0 {
			t = current[0 : close+2]
		}
	}
	if badgeValue > 0 {
		return fmt.Sprintf("(%v) %v", strconv.Itoa(badgeValue), t)
	}
	return t
}
