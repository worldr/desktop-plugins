package appbadge

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

const char*
nsString2cString(NSString* ns) {
    if (ns == NULL) { return NULL; }
    const char* cs = [ns UTF8String];
    return cs;
}

NSString*
cString2nsString(char* cs) {
    if (cs == NULL) { return NULL; }
    NSString* ns = [NSString stringWithUTF8String:cs];
    return ns;
}

void
platformSetWindowTitle(char* value) {
	NSString* str = cString2nsString(value);
	NSWindow* window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
	window.title = str;
}

char*
platformGetWindowTitle() {
	NSWindow* window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
  return nsString2cString(window.title);
}

void
platformSetBadge(int* value) {
	NSString* str = [NSString stringWithFormat:@"%i", *value];
	NSDockTile* tile = [[NSApplication sharedApplication] dockTile];
	[tile setBadgeLabel:str];
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

type AppBadgeDarwin struct{}

func (*AppBadgeDarwin) SetBadge(value int32) error {
	// get current title
	a := C.platformGetWindowTitle()
	fmt.Println(fmt.Sprintf("-- type %T", a))
	fmt.Println(fmt.Sprintf("-- reflect %v", reflect.TypeOf(a)))
	fmt.Println(fmt.Sprintf("-- reflect %v", reflect.TypeOf(a).String()))
	fmt.Println(fmt.Sprintf("-- value %v", a))

	fmt.Println(fmt.Sprintf("-- type %T", *a))
	fmt.Println(fmt.Sprintf("-- reflect %v", reflect.TypeOf(*a)))
	fmt.Println(fmt.Sprintf("-- reflect %v", reflect.TypeOf(*a).String()))
	fmt.Println(fmt.Sprintf("-- value %v", *a))

	gs := C.GoString(a)
	fmt.Println(fmt.Sprintf("-- type %T", gs))
	fmt.Println(fmt.Sprintf("-- reflect %v", reflect.TypeOf(gs)))
	fmt.Println(fmt.Sprintf("-- reflect %v", reflect.TypeOf(gs).String()))
	fmt.Println(fmt.Sprintf("-- value %v", gs))

	// create new title with counter
	cs2 := C.CString(formatWindowTitle(gs, value))
	defer C.free(unsafe.Pointer(cs2))

	// set new title
	C.platformSetWindowTitle(cs2)

	// Set badge number
	C.platformSetBadge(&v)

	return nil
}

func (me *AppBadgeDarwin) ClearBadge() error {
	return me.SetBadge(0)
}

func init() {
	Api = &AppBadgeDarwin{}
}
