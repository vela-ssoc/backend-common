////go:build !windows && !solaris && !aix

package teletype

import "syscall"

func ioctl(fd, cmd, ptr uintptr) error {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptr); err != 0 {
		return err
	}
	return nil
}
