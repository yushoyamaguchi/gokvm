package eventfd

import (
	"encoding/binary"
	"syscall"
)

const EFD_NONBLOCK = 0x800

type Eventfd struct {
	fd int
}

func Create() (*Eventfd, error) {
	// Use EFD_NONBLOCK flag for non-blocking mode
	fd, _, errno := syscall.Syscall(syscall.SYS_EVENTFD2, 0, EFD_NONBLOCK, 0)
	if errno != 0 {
		return nil, errno
	}
	return &Eventfd{int(fd)}, nil
}

func (e *Eventfd) Close() error {
	return syscall.Close(e.fd)
}

func (e *Eventfd) Read() (uint64, error) {
	var buf [8]byte
	_, err := syscall.Read(e.fd, buf[:])
	if err != nil {
		if err == syscall.EAGAIN {
			// No data to read in non-blocking mode
			return 0, nil
		}
		return 0, err
	}
	return binary.LittleEndian.Uint64(buf[:]), nil
}

func (e *Eventfd) Write(value uint64) error {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], value)
	_, err := syscall.Write(e.fd, buf[:])
	return err
}
