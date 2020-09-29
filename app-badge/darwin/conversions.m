#import <Cocoa/Cocoa.h>

const char*
nsstring2cstring(NSString *s) {
    if (s == NULL) { return NULL; }
    const char *cstr = [s UTF8String];
    return cstr;
}

int
nsnumber2cint(NSNumber *i) {
    if (i == NULL) { return 0; }
    return i.intValue;
}