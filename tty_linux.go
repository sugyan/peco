// +build linux

package peco

import (
	"os"
	"syscall"
	"unsafe"
)

func IsTty() bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

func TtyReady() error {
	return nil
}

func TtyTerm() {
}
