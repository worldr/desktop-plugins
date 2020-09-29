package darwin

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
#include "conversions.h"
*/
import "C"
import "unsafe"

// NSString -> C string
// NOTE: free memory manually on the Go side!
// defer C.free(unsafe.Pointer(cs))
func cString(s interface{}) *C.char {
	return C.nsstring2cstring(unsafe.Pointer(s))
}

// NSString -> Go string
func GoString(s interface{}) (string, *C.char) {
	str := cString(s)
	return C.GoString(str), str
}

// NSNumber -> Go int
func GoInt(i interface{}) int {
	return int(C.nsnumber2cint(unsafe.Pointer(i)))
}
