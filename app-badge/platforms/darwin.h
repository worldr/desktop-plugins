#import <Cocoa/Cocoa.h>
int
PlatformSetWindowTitle(NSString* value) {
	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
	window.title = value;
  return 0;
}
NSString*
PlatformGetWindowTitle() {
	NSWindow *window = [[[NSApplication sharedApplication] windows] objectAtIndex:0];
  return window.title;
}
int
PlatformSetBadge(int value) {
	NSString* str = [NSString stringWithFormat:@"%i", value];
	NSDockTile* tile = [[NSApplication sharedApplication] dockTile];
	[tile setBadgeLabel:str];	
  return 0;
}
