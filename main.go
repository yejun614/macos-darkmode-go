package main

// [Original Source Code]
// https://github.com/mafredri/macos-darkmode/blob/master/cmd/darkmode/main.go
//
// [Useful References]
// https://pkg.go.dev/cmd/cgo

/*
#cgo LDFLAGS: -F/System/Library/PrivateFrameworks -framework SkyLight

#import <objc/objc.h>

// Private APIs for Dark Mode.
void SLSSetAppearanceThemeLegacy(char);
char SLSGetAppearanceThemeLegacy();
*/
import "C"

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func darkModeEnabled() bool {
	return C.SLSGetAppearanceThemeLegacy() == 1
}

func setDarkMode(on bool) {
	var b int = 0
	if on {
		b = 1
	}
	C.SLSSetAppearanceThemeLegacy(C.char(b))
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [on]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [off]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [toggle]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := os.Args[1:]
	switch len(args) {
	case 1:
		switch strings.ToLower(args[0]) {
		case "on":
			setDarkMode(true)
		case "off":
			setDarkMode(false)
		case "toggle":
			setDarkMode(!darkModeEnabled())
		default:
			usage()
		}
	case 0:
		if darkModeEnabled() {
			fmt.Println("on")
		} else {
			fmt.Println("off")
		}
	default:
		usage()
	}
}
