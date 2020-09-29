#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
int
SetWindowTitle(String value) {
	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
	window.title = value;
  return 0;
}