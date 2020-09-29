package appbadge

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <Cocoa/Cocoa.h>
int
platformSetWindowTitle(NSString* value) {
	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
	window.title = value;
  return 0;
}
NSString*
platformGetWindowTitle() {
	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
  return window.title;
}
int
platformSetBadge(int value) {
	NSString* str = [NSString stringWithFormat:@"%i", value];
	NSDockTile* tile = [[NSApplication sharedApplication] dockTile];
	[tile setBadgeLabel:str];
  return 0;
}
*/
import "C"

type AppBadgeDarwin struct{}

func (*AppBadgeDarwin) SetBadge(value int) error {
	r1 := platformSetWindowTitle(formatWindowTitle(platformGetWindowTitle(), value))
	if r1 != 0 {
		return newError("Failed to set window title")
	}
	r2 := platformSetBadge(value)
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
