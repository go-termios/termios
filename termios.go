package termios // import "gopkg.in/termios.v0"

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

// // From termios.h
// // You could just use the ones from syscall (aka sys/unix) directly,
// // but having them here is more convenient (syscall is too big).
// const (
// 	// c_cflag bits:
// 	B0      = uint32(unix.B0)
// 	B50     = uint32(unix.B50)
// 	B75     = uint32(unix.B75)
// 	B110    = uint32(unix.B110)
// 	B134    = uint32(unix.B134)
// 	B150    = uint32(unix.B150)
// 	B200    = uint32(unix.B200)
// 	B300    = uint32(unix.B300)
// 	B600    = uint32(unix.B600)
// 	B1200   = uint32(unix.B1200)
// 	B1800   = uint32(unix.B1800)
// 	B2400   = uint32(unix.B2400)
// 	B4800   = uint32(unix.B4800)
// 	B9600   = uint32(unix.B9600)
// 	B19200  = uint32(unix.B19200)
// 	B38400  = uint32(unix.B38400)
// 	EXTA    = uint32(unix.EXTA)
// 	EXTB    = uint32(unix.EXTB)
// 	CSIZE   = uint32(unix.CSIZE)
// 	CS5     = uint32(unix.CS5)
// 	CS6     = uint32(unix.CS6)
// 	CS7     = uint32(unix.CS7)
// 	CS8     = uint32(unix.CS8)
// 	CSTOPB  = uint32(unix.CSTOPB)
// 	CREAD   = uint32(unix.CREAD)
// 	PARENB  = uint32(unix.PARENB)
// 	PARODD  = uint32(unix.PARODD)
// 	HUPCL   = uint32(unix.HUPCL)
// 	CLOCAL  = uint32(unix.CLOCAL)
// 	B57600  = uint32(unix.B57600)
// 	B115200 = uint32(unix.B115200)
// 	B230400 = uint32(unix.B230400)
// 	B460800 = uint32(unix.B460800)
// 	B921600 = uint32(unix.B921600)

// 	// c_iflag bits:
// 	IGNBRK  = uint32(unix.IGNBRK)
// 	BRKINT  = uint32(unix.BRKINT)
// 	IGNPAR  = uint32(unix.IGNPAR)
// 	PARMRK  = uint32(unix.PARMRK)
// 	INPCK   = uint32(unix.INPCK)
// 	ISTRIP  = uint32(unix.ISTRIP)
// 	INLCR   = uint32(unix.INLCR)
// 	IGNCR   = uint32(unix.IGNCR)
// 	ICRNL   = uint32(unix.ICRNL)
// 	IXON    = uint32(unix.IXON)
// 	IXANY   = uint32(unix.IXANY)
// 	IXOFF   = uint32(unix.IXOFF)
// 	IMAXBEL = uint32(unix.IMAXBEL)

// 	// c_oflag bits:
// 	OPOST  = uint32(unix.OPOST)
// 	ONLCR  = uint32(unix.ONLCR)
// 	OCRNL  = uint32(unix.OCRNL)
// 	ONOCR  = uint32(unix.ONOCR)
// 	ONLRET = uint32(unix.ONLRET)

// 	// c_lflag bits:
// 	ISIG    = uint32(unix.ISIG)
// 	ICANON  = uint32(unix.ICANON)
// 	ECHO    = uint32(unix.ECHO)
// 	ECHOE   = uint32(unix.ECHOE)
// 	ECHOK   = uint32(unix.ECHOK)
// 	ECHONL  = uint32(unix.ECHONL)
// 	NOFLSH  = uint32(unix.NOFLSH)
// 	TOSTOP  = uint32(unix.TOSTOP)
// 	ECHOCTL = uint32(unix.ECHOCTL)
// 	ECHOPRT = uint32(unix.ECHOPRT)
// 	ECHOKE  = uint32(unix.ECHOKE)
// 	FLUSHO  = uint32(unix.FLUSHO)
// 	PENDIN  = uint32(unix.PENDIN)
// 	IEXTEN  = uint32(unix.IEXTEN)
// 	EXTPROC = uint32(unix.EXTPROC)
// )

type Termios unix.Termios

func ioctl(fd, cmd uintptr, arg unsafe.Pointer) error {
	return ioctlu(fd, cmd, uintptr(arg))
}

func ioctlu(fd, cmd, arg uintptr) error {
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, fd, cmd, arg)
	if errno == 0 {
		return nil
	}
	return errno
}

// GetAttr gets the attributes of the given terminal.
func GetAttr(fd uintptr) (*Termios, error) {
	var tio Termios
	if err := ioctl(fd, getAttrIOCTL, unsafe.Pointer(&tio)); err != nil {
		return nil, err
	}

	return &tio, nil
}

// SetAttr sets the attributes of the given terminal from this termios
// structure, immediately. See tcsetattr(3).
func (tio *Termios) SetAttr(fd uintptr) error {
	return ioctl(fd, setAttrNowIOCTL, unsafe.Pointer(&tio))
}

// DrainAndSetAttr sets the attributes of the given terminal from this termios
// structure, after all output written to fd has been transmitted. This option
// should be used when changing parameters that affect output. See tcsetattr(3).
func (tio *Termios) DrainAndSetAttr(fd uintptr) error {
	return ioctl(fd, setAttrDrainIOCTL, unsafe.Pointer(&tio))
}

// FlushAndSetAttr sets the attributes of the given terminal from this termios
// structure, “after all output written to fd has been transmitted, and all
// input that has been received but not read will be discarded before the
// change is made”. See tcsetattr(3).
func (tio *Termios) FlushAndSetAttr(fd uintptr) error {
	return ioctl(fd, setAttrFlushIOCTL, unsafe.Pointer(&tio))
}

// MakeRaw returns a copy of the termios structure with the flags set to a
// state disabling all input and output processing, giving a “raw I/O
// path”. See cfmakeraw(3). Note the exact flags are platform-dependent.
//
// Note also this does *not* apply it to a terminal; for that call one of the
// SetAttr methods on the returned struct.
func (tio *Termios) MakeRaw() *Termios {
	copy := *tio

	// from cfmakeraw(3)
	copy.Iflag = (copy.Iflag & rawImaskOff) | rawImaskOn
	copy.Oflag = (copy.Oflag & rawOmaskOff) | rawOmaskOn
	copy.Lflag = (copy.Lflag & rawLmaskOff) | rawLmaskOn
	copy.Cflag = (copy.Cflag & rawCmaskOff) | rawCmaskOn
	copy.Cc[unix.VMIN] = 1
	copy.Cc[unix.VTIME] = 0

	return &copy
}

type winSize struct {
	row    uint16
	col    uint16
	xpixel uint16 // unused
	Ypixel uint16 // unused
}

// GetWinSize gets the current size of the terminal referred to by fd.
func GetWinSize(fd uintptr) (width int, height int, err error) {
	// from tty_ioctl(4)
	ws := winSize{}

	if err := ioctl(fd, unix.TIOCGWINSZ, unsafe.Pointer(&ws)); err != nil {
		return -1, -1, err
	}

	return int(ws.col), int(ws.row), nil
}

// SetWinSize sets the size of the terminal referred to by fd.
func SetWinSize(fd uintptr, width, height int) error {
	ws := winSize{
		col: uint16(width),
		row: uint16(height),
	}

	return ioctl(fd, unix.TIOCSWINSZ, unsafe.Pointer(&ws))
}

type Queue uint8

const (
	AnyQueue Queue = iota
	InputQueue
	OutputQueue
)

// Flush discards data written to the terminal but not transmitted, or
// received from the terminal but not read, depending on the queue.
//
// InputQueue   received but not read.
// OutputQueue  written but not transmitted.
// AnyQueue     both
func Flush(fd uintptr, queue Queue) error {
	return ioctlu(fd, flushIOCTL, queue.bits())
}
