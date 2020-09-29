package appbadge

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
void
platformSetWindowTitle(char* value) {
	NSString* str = [NSString stringWithUTF8String:value];
	NSWindow* window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
	window.title = str;
}
NSString*
platformGetWindowTitle() {
	NSWindow* window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
  return window.title;
}
NSNumber*
platformSetBadge(int* value) {
	NSString* str = [NSString stringWithFormat:@"%i", *value];
	NSDockTile* tile = [[NSApplication sharedApplication] dockTile];
	[tile setBadgeLabel:str];
  return [NSNumber numberWithInt:0];
}
*/
import "C"
import (
	"unsafe"

	"github.com/worldr/desktop-plugins/app-badge/darwin"
)

type AppBadgeDarwin struct{}

func (*AppBadgeDarwin) SetBadge(value int) error {
	// get current title
	ns := C.platformGetWindowTitle()
	gs, cs1 := darwin.GoString(ns)
	defer C.free(unsafe.Pointer(cs1))

	// create new title with counter
	cs2 := C.CString(formatWindowTitle(gs, value))
	defer C.free(unsafe.Pointer(cs2))

	// set new title
	r1 := darwin.GoInt(C.platformSetWindowTitle(cs2))
	if r1 != 0 {
		return newError("Failed to set window title")
	}
	v := C.int(value)
	r2 := darwin.GoInt(C.platformSetBadge(&v))
	if r2 != 0 {
		return newError("Failed to set app badge value")
	}
	return nil
}

func (me *AppBadgeDarwin) ClearBadge() error {
	return me.SetBadge(0)
}

func init() {
	Api = &AppBadgeDarwin{}
}
