//go:build linux && !appengine
// +build linux,!appengine

package ansitty

import (
	"syscall"
	"unsafe"
)

func setEnabled(fd uintptr, enabled bool) bool {
	if !enabled {
		return false
	}

	var termios syscall.Termios
	_, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, fd, syscall.TCGETS, uintptr(unsafe.Pointer(&termios)), 0, 0, 0) //nolint:gosec

	return errno == 0
}
