package godarkmode

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#include "darkmode.h"
*/
import "C"

// DarkModeEnabled returns whether or not the dark theme is enabled on macOS
func IsDarkModeEnabled() bool {
	if C.IsDarkMode() {
		return true
	}
	return false
}
