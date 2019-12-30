# godarkmode

godarkmode is a lib used to determine whether or not dark theme is enabled on macOS.

it will query UserDefaults through C/Objective-C headers, get the AppleInterfaceStyles value and parse it to a go boolean type.

## Installation

install via go get as shown below:

```bash
go get github.com/InfeCtlll3/godarkmode
```

## Usage
#### test.go:

```go
package main
import (
	"github.com/InfeCtlll3/godarkmode"
	"fmt"
)

func main() {
    fmt.Println("Dark mode is:", godarkmode.IsDarkModeEnabled())
}
```

### Output
```bash
% go run test.go

Dark mode is:  true
```

## Issues
In case on any issues please either open a ticket or e-mail me at contato.carmando@gmail.com

## Known issues
#### Clang link errors:
```bash
# command-line-arguments
Undefined symbols for architecture x86_64:
  "_IsDarkMode", referenced from:
      __cgo_e3e2416a22ae_Cfunc_IsDarkMode in _x002.o
     (maybe you meant: __cgo_e3e2416a22ae_Cfunc_IsDarkMode)
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```
#### How to solve:
Sometimes due to some weird clang/xcode configuration this error happens. Make sure you clean the cache of your go build with ```go clean -cache``` and try to build it using ```go build -x```
