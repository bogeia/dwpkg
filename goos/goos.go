package goos

import (
	"runtime"
)

const IsAix = "aix"
const IsAndroid = "android"
const IsDarwin = "darwin"
const IsDragonfly = "dragonfly"
const IsFreebsd = "freebsd"
const IsHurd = "hurd"
const IsIllumos = "illumos"
const IsIos = "ios"
const IsJs = "js"
const IsLinux = "linux"
const IsNacl = "nacl"
const IsNetbsd = "netbsd"
const IsOpenbsd = "openbsd"
const IsPlan9 = "plan9"
const IsSolaris = "solaris"
const IsWindows = "windows"
const IsZos = "zos"

// IsSpecifiedPlatform checks for the specified platform type.
func IsSpecifiedPlatform(platform string) bool {
	if platform == runtime.GOOS {
		return true
	}
	return false
}
