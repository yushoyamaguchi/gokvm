package vhost

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/bobuhiro11/gokvm/eventfd"
)

type VhostDev struct {
}

type VhostNet struct {
	dev     VhostDev
	vhostfd int
	nvqs    uint
	vq      [DefaultNVQs]VhostVirtqueue
}

type VhostVirtqueue struct {
	maskedNotifier eventfd.EventNotifier
}

type VhostVringFile struct {
	Index uint32
	FD    int32
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

func (v *VhostNet) vhostVirtqueueInit(vq *VhostVirtqueue, index uint) error {
	fmt.Println("vhostVirtqueueInit")
	err := vq.maskedNotifier.EventNotifierInit()
	if err != nil {
		fmt.Printf("EventNotifierInit: %v\n", err)
	}
	file := VhostVringFile{
		Index: uint32(index),
	}
	file.FD = int32(vq.maskedNotifier.GetWfd())
	_, err = ioctl(uintptr(v.vhostfd), VHOST_SET_VRING_CALL, uintptr(unsafe.Pointer(&file)))
	if err != nil {
		fmt.Printf("ioctl VHOST_SET_VRING_CALL: %v\n", err)
	}
	return nil
}

func NewVhostNet() (*VhostNet, error) {
	fmt.Println("NewVhostNet")
	vhost := &VhostNet{
		dev:  VhostDev{},
		nvqs: DefaultNVQs,
	}
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
		fmt.Printf("ioctl VHOST_GET_BACKEND_FEATURES: %v\n", err)
		return nil, err
	}
	fmt.Printf("feature: %v\n", feature)
	_, err = ioctl(uintptr(vhost.vhostfd), VHOST_SET_OWNER, 0)
	if err != nil {
		fmt.Printf("ioctl VHOST_SET_OWNER: %v\n", err)
		return nil, err
	}

	for i := uint(0); i < vhost.nvqs; i++ {
		vhost.vhostVirtqueueInit(&vhost.vq[i], i)
	}

	return vhost, nil
}
