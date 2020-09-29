package darwin

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
#include "conversions.h"
*/
import "C"

// NSString -> C string
// NOTE: free memory manually on the Go side!
// defer C.free(unsafe.Pointer(cs))
func cString(s *C.NSString) *C.char { return C.nsstring2cstring(s) }

// NSString -> Go string
func GoString(s *C.NSString) (string, *C.char) {
	str := cString(s)
	return C.GoString(str), str
}

// NSNumber -> Go int
func GoInt(i *C.NSNumber) int { return int(C.nsnumber2int(i)) }
