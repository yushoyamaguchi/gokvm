package vhost

import "unsafe"

const (
	iocNBits    = 8
	iocTypeBits = 8
	iocSizeBits = 14
	iocDirBits  = 2

	// 各ビットのシフト量
	iocNRShift   = 0
	iocTypeShift = iocNRShift + iocNBits      // 8
	iocSizeShift = iocTypeShift + iocTypeBits // 16
	iocDirShift  = iocSizeShift + iocSizeBits // 30

	// _IOC_READ, _IOC_WRITE 等
	IOC_NONE  = 0
	IOC_WRITE = 1
	IOC_READ  = 2

	DefaultNVQs = 2
)

var (
	VHOSTVIRTIO                = 0xAF
	VHOST_GET_BACKEND_FEATURES = IOR(
		uintptr(VHOSTVIRTIO),
		0x26,
		uintptr(unsafe.Sizeof(uint64(0))),
	)
	VHOST_SET_OWNER = IO(
		uintptr(VHOSTVIRTIO),
		0x01,
	)
	VHOST_SET_VRING_CALL = IOW(
		uintptr(VHOSTVIRTIO),
		0x21,
		uintptr(unsafe.Sizeof(VhostVringFile{})),
	)
)

// _IO() 相当のラッパ (データの受け渡しが不要な場合)
func IO(t, nr uintptr) uintptr {
	return IOC(IOC_NONE, t, nr, 0)
}

// _IOC() 相当の関数
func IOC(dir, t, nr, size uintptr) uintptr {
	return (dir << iocDirShift) |
		(t << iocTypeShift) |
		(nr << iocNRShift) |
		(size << iocSizeShift)
}

// _IOR() 相当のラッパ
func IOR(t, nr, size uintptr) uintptr {
	return IOC(IOC_READ, t, nr, size)
}

// _IOW() 相当のラッパ (読み取り不要の場合)
func IOW(t, nr, size uintptr) uintptr {
	return IOC(IOC_WRITE, t, nr, size)
}

// _IOWR() 相当のラッパ (読み書き両方の場合)
func IOWR(t, nr, size uintptr) uintptr {
	return IOC(IOC_READ|IOC_WRITE, t, nr, size)
}
