#import <Cocoa/Cocoa.h>
int
PlatformSetWindowTitle(String value) {
	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
	window.title = value;
  return 0;
}
String
PlatformGetWindowTitle() {
	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
  return window.title;
}
int
PlatformSetBadge(int value) {
	[UIApplication sharedApplication].applicationIconBadgeNumber = value;
  return 0;
}
