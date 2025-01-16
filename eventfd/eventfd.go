package eventfd

import (
	"syscall"
)

type Eventfd struct {
	fd int
}

func Create() (*Eventfd, error) {
	fd, _, errno := syscall.Syscall(syscall.SYS_EVENTFD2, 0, 0, 0)
	if errno != 0 {
		return nil, errno
	}
	return &Eventfd{int(fd)}, nil
}
