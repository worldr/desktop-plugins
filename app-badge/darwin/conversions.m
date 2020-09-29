#import <Cocoa/Cocoa.h>

const char*
nsstring2cstring(void* s) {
    if (s == NULL) { return NULL; }
		NSString *nss = *((__unsafe_unretained NSString **)(s));
    const char *cstr = [nss UTF8String];
    return cstr;
}

int
nsnumber2cint(void* i) {
    if (i == NULL) { return 0; }
		NSNumber* nsi = *((__unsafe_unretained NSNumber **)(i));
    return nsi.intValue;
}