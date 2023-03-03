package teletype

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"unsafe"
)

func Open(ctx context.Context) (io.ReadWriteCloser, error) {
	pty, err := openPTY()
	if err != nil {
		return nil, err
	}

	name := shellCmd()
	cmd := exec.CommandContext(ctx, name)
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	cmd.SysProcAttr.Setsid = true
	cmd.SysProcAttr.Setctty = true
	cmd.Stdin = pty.slave
	cmd.Stdout = pty.slave
	cmd.Stderr = pty.slave
	if err = cmd.Run(); err != nil {
		_ = pty.Close()
		return nil, err
	}
	return pty, nil
}

func shellCmd() string {
	name := os.Getenv("SHELL")
	if name == "" {
		name = "/bin/bash"
	}
	return name
}

func openPTY() (*linuxPTY, error) {
	master, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = master.Close()
		}
	}()

	fd := master.Fd()
	sname, err := ptsname(fd)
	if err != nil {
		return nil, err
	}
	if err = unlockpt(fd); err != nil {
		return nil, err
	}

	slave, err := os.OpenFile(sname, os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		return nil, err
	}

	pty := &linuxPTY{
		master: master,
		slave:  slave,
	}

	return pty, nil
}

func ptsname(fd uintptr) (string, error) {
	var n uint32
	if err := ioctl(fd, syscall.TIOCGPTN, uintptr(unsafe.Pointer(&n))); err != nil {
		return "", err
	}
	return "/dev/pts/" + strconv.Itoa(int(n)), nil
}

func unlockpt(fd uintptr) error {
	var n uint32
	return ioctl(fd, syscall.TIOCSPTLCK, uintptr(unsafe.Pointer(&n)))
}

type linuxPTY struct {
	master *os.File
	slave  *os.File
}

func (pt *linuxPTY) Read(p []byte) (int, error) {
	return pt.master.Write(p)
}

func (pt *linuxPTY) Write(p []byte) (int, error) {
	return pt.master.Write(p)
}

func (pt *linuxPTY) Close() error {
	_ = pt.slave.Close()
	return pt.master.Close()
}
