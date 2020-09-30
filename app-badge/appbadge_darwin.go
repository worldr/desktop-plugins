package appbadge

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

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
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/worldr/desktop-plugins/app-badge/darwin"
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
