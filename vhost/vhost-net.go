package vhost

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	VHOSTVIRTIO                = 0xAF
	VHOST_GET_BACKEND_FEATURES = IOR(
		uintptr(VHOSTVIRTIO),
		0x26,
		uintptr(unsafe.Sizeof(uint64(0))),
	)
)

type VhostNet struct {
	vhostfd int
}

func open(name string, flags int) (int, error) {
	res, err := syscall.Open(name, flags, 0)
	if err != nil {
		return res, err
	}
	return res, nil
}

func ioctl(fd, op, arg uintptr) (uintptr, error) {
	res, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL, fd, op, arg)
	if errno != 0 {
		return res, errno
	}

	return res, nil
}

func NewVhostNet() (*VhostNet, error) {
	fmt.Println("NewVhostNet")
	vhost := &VhostNet{}
	vhostfd, err := open("/dev/vhost-net", syscall.O_RDWR)
	if err != nil {
		fmt.Printf("open: %v\n", err)
		return nil, err
	}
	vhost.vhostfd = vhostfd
	fmt.Printf("vhostfd: %v\n", vhost.vhostfd)
	feature := uint64(0)
	_, err = ioctl(uintptr(vhost.vhostfd), VHOST_GET_BACKEND_FEATURES, uintptr(unsafe.Pointer(&feature)))
	if err != nil {
		fmt.Printf("ioctl: %v\n", err)
		return nil, err
	}
	fmt.Printf("feature: %v\n", feature)
	return vhost, nil
}
