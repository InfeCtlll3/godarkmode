#import <Foundation/Foundation.h>

bool IsDarkMode(){
    const NSString *mode = [[NSUserDefaults standardUserDefaults] stringForKey:@"AppleInterfaceStyle"];
    if (mode == NULL) {
        return false;
    }
    return true;
}