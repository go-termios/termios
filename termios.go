package termios // import "gopkg.in/termios.v0"

import (
	"fmt"
	"sort"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	// from bits/termios.h
	// You could just use the ones from syscall (aka sys/unix) directly,
	// but having them here is more convenient (syscall is too big).

	/* c_cflag bit meaning */
	CBAUD    = uint32(unix.CBAUD)
	B0       = uint32(unix.B0)
	B50      = uint32(unix.B50)
	B75      = uint32(unix.B75)
	B110     = uint32(unix.B110)
	B134     = uint32(unix.B134)
	B150     = uint32(unix.B150)
	B200     = uint32(unix.B200)
	B300     = uint32(unix.B300)
	B600     = uint32(unix.B600)
	B1200    = uint32(unix.B1200)
	B1800    = uint32(unix.B1800)
	B2400    = uint32(unix.B2400)
	B4800    = uint32(unix.B4800)
	B9600    = uint32(unix.B9600)
	B19200   = uint32(unix.B19200)
	B38400   = uint32(unix.B38400)
	EXTA     = uint32(unix.EXTA)
	EXTB     = uint32(unix.EXTB)
	CSIZE    = uint32(unix.CSIZE)
	CS5      = uint32(unix.CS5)
	CS6      = uint32(unix.CS6)
	CS7      = uint32(unix.CS7)
	CS8      = uint32(unix.CS8)
	CSTOPB   = uint32(unix.CSTOPB)
	CREAD    = uint32(unix.CREAD)
	PARENB   = uint32(unix.PARENB)
	PARODD   = uint32(unix.PARODD)
	HUPCL    = uint32(unix.HUPCL)
	CLOCAL   = uint32(unix.CLOCAL)
	CBAUDEX  = uint32(unix.CBAUDEX)
	B57600   = uint32(unix.B57600)
	B115200  = uint32(unix.B115200)
	B230400  = uint32(unix.B230400)
	B460800  = uint32(unix.B460800)
	B500000  = uint32(unix.B500000)
	B576000  = uint32(unix.B576000)
	B921600  = uint32(unix.B921600)
	B1000000 = uint32(unix.B1000000)
	B1152000 = uint32(unix.B1152000)
	B1500000 = uint32(unix.B1500000)
	B2000000 = uint32(unix.B2000000)
	B2500000 = uint32(unix.B2500000)
	B3000000 = uint32(unix.B3000000)
	B3500000 = uint32(unix.B3500000)
	B4000000 = uint32(unix.B4000000)
	CIBAUD   = uint32(unix.CIBAUD)
	CMSPAR   = uint32(unix.CMSPAR)
	CRTSCTS  = uint32(unix.CRTSCTS)

	/* c_iflag bits */
	IGNBRK  = uint32(unix.IGNBRK)
	BRKINT  = uint32(unix.BRKINT)
	IGNPAR  = uint32(unix.IGNPAR)
	PARMRK  = uint32(unix.PARMRK)
	INPCK   = uint32(unix.INPCK)
	ISTRIP  = uint32(unix.ISTRIP)
	INLCR   = uint32(unix.INLCR)
	IGNCR   = uint32(unix.IGNCR)
	ICRNL   = uint32(unix.ICRNL)
	IUCLC   = uint32(unix.IUCLC)
	IXON    = uint32(unix.IXON)
	IXANY   = uint32(unix.IXANY)
	IXOFF   = uint32(unix.IXOFF)
	IMAXBEL = uint32(unix.IMAXBEL)
	IUTF8   = uint32(unix.IUTF8)

	/* c_oflag bits */
	OPOST  = uint32(unix.OPOST)
	OLCUC  = uint32(unix.OLCUC)
	ONLCR  = uint32(unix.ONLCR)
	OCRNL  = uint32(unix.OCRNL)
	ONOCR  = uint32(unix.ONOCR)
	ONLRET = uint32(unix.ONLRET)
	OFILL  = uint32(unix.OFILL)
	OFDEL  = uint32(unix.OFDEL)
	NLDLY  = uint32(unix.NLDLY)
	NL0    = uint32(unix.NL0)
	NL1    = uint32(unix.NL1)
	CRDLY  = uint32(unix.CRDLY)
	CR0    = uint32(unix.CR0)
	CR1    = uint32(unix.CR1)
	CR2    = uint32(unix.CR2)
	CR3    = uint32(unix.CR3)
	TABDLY = uint32(unix.TABDLY)
	TAB0   = uint32(unix.TAB0)
	TAB1   = uint32(unix.TAB1)
	TAB2   = uint32(unix.TAB2)
	TAB3   = uint32(unix.TAB3)
	BSDLY  = uint32(unix.BSDLY)
	BS0    = uint32(unix.BS0)
	BS1    = uint32(unix.BS1)
	FFDLY  = uint32(unix.FFDLY)
	FF0    = uint32(unix.FF0)
	FF1    = uint32(unix.FF1)

	/* c_lflag bits */
	ISIG    = uint32(unix.ISIG)
	ICANON  = uint32(unix.ICANON)
	XCASE   = uint32(unix.XCASE)
	ECHO    = uint32(unix.ECHO)
	ECHOE   = uint32(unix.ECHOE)
	ECHOK   = uint32(unix.ECHOK)
	ECHONL  = uint32(unix.ECHONL)
	NOFLSH  = uint32(unix.NOFLSH)
	TOSTOP  = uint32(unix.TOSTOP)
	ECHOCTL = uint32(unix.ECHOCTL)
	ECHOPRT = uint32(unix.ECHOPRT)
	ECHOKE  = uint32(unix.ECHOKE)
	FLUSHO  = uint32(unix.FLUSHO)
	PENDIN  = uint32(unix.PENDIN)
	IEXTEN  = uint32(unix.IEXTEN)
	EXTPROC = uint32(unix.EXTPROC)
)

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
	if err := ioctl(fd, unix.TCGETS, unsafe.Pointer(&tio)); err != nil {
		return nil, err
	}

	return &tio, nil
}

// GetLock gets the locking status of the termios structure of the given
// terminal. See tty_ioctl(4).
func GetLock(fd uintptr) (*Termios, error) {
	var tio Termios
	if err := ioctl(fd, unix.TIOCGLCKTRMIOS, unsafe.Pointer(&tio)); err != nil {
		return nil, err
	}

	return &tio, nil
}

// SetLock sets the locking status of the termios structure of the given
// terminal. Needs CAP_SYS_ADMIN. See tty_ioctl(4).
func (tio *Termios) SetLock(fd uintptr) error {
	return ioctl(fd, unix.TIOCSLCKTRMIOS, unsafe.Pointer(&tio))
}

// SetAttr sets the attributes of the given terminal from this termios
// structure, immediately. See tcsetattr(3).
func (tio *Termios) SetAttr(fd uintptr) error {
	return ioctl(fd, unix.TCSETS, unsafe.Pointer(&tio))
}

// DrainAndSetAttr sets the attributes of the given terminal from this termios
// structure, after all output written to fd has been transmitted. This option
// should be used when changing parameters that affect output. See tcsetattr(3).
func (tio *Termios) DrainAndSetAttr(fd uintptr) error {
	return ioctl(fd, unix.TCSETSW, unsafe.Pointer(&tio))
}

// FlushAndSetAttr sets the attributes of the given terminal from this termios
// structure, “after all output written to fd has been transmitted, and all
// input that has been received but not read will be discarded before the
// change is made”. See tcsetattr(3).
func (tio *Termios) FlushAndSetAttr(fd uintptr) error {
	return ioctl(fd, unix.TCSETSF, unsafe.Pointer(&tio))
}

// MakeRaw returns a copy of the termios structure with the flags set to a
// state disabling all input and output processing, giving a “raw I/O
// path”. See cfmakeraw(3).
//
// Note this does *not* apply it to a terminal; for that call one of the
// SetAttr methods on the returned struct.
func (tio *Termios) MakeRaw() *Termios {
	copy := *tio

	// from cfmakeraw(3)
	copy.Iflag &= ^(IGNBRK | BRKINT | PARMRK | ISTRIP | INLCR | IGNCR | ICRNL | IXON)
	copy.Oflag &= ^OPOST
	copy.Lflag &= ^(ECHO | ECHONL | ICANON | ISIG | IEXTEN)
	copy.Cflag &= ^(CSIZE | PARENB)
	copy.Cflag |= CS8

	return &copy
}

// from the Bnnnnnn constants above
var speeds = [...]int{
	// first 16:
	0, 50, 75, 110, 134, 150, 200, 300, 600, 1200, 1800, 2400, 4800, 9600, 19200, 38400,
	// extended:
	57600, 115200, 230400, 460800, 500000, 576000, 921600, 1000000,
	1152000, 1500000, 2000000, 2500000, 3000000, 3500000, 4000000,
}

// GetSpeed gets the stored baud rate.
func (tio *Termios) GetSpeed() int {
	spID := int(tio.Cflag & CBAUD)
	if spID <= 0xf {
		return speeds[spID]
	}
	spID -= int(CBAUDEX) - 0xf
	if 0xf >= spID || spID >= len(speeds) {
		return -1
	}
	return speeds[spID]
}

// SetSpeed sets the baud rate.
func (tio *Termios) SetSpeed(speed int) error {
	spID := sort.SearchInts(speeds[:], speed)
	if spID >= len(speeds) || speed != speeds[spID] {
		return fmt.Errorf("%d is not a good rate, even for a baud rate", speed)
	}
	val := uint32(spID)
	if val > 0xf {
		val += CBAUDEX - 0xf
	}
	tio.Cflag = (tio.Cflag &^ CBAUD) | val

	return nil
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

	return int(ws.row), int(ws.col), nil
}

// GetWinSize sets the size of the terminal referred to by fd.
func SetWinSize(fd uintptr, width, height int) error {
	ws := winSize{
		row: uint16(width),
		col: uint16(height),
	}

	return ioctl(fd, unix.TIOCSWINSZ, unsafe.Pointer(&ws))
}

// SendBreak transmits a continuous stream of zeros for between 0.25 and 0.5
// seconds if duration is 0, or for the given number of deciseconds if not, if
// the terminal supports breaks.
//
// Note this is TCSBRKP in tty_ioctl(4); this one seems saner than
// TCSBRK/tcsendbreak(3).
func SendBreak(fd uintptr, duration int) error {
	return ioctlu(fd, unix.TCSBRKP, uintptr(duration))
}

// Drain waits until all output written to the terminal referenced by fd has
// been transmitted to the terminal.
func Drain(fd uintptr) error {
	// on linux tcdrain is TCSBRK with non-zero arg
	return ioctlu(fd, unix.TCSBRK, uintptr(1))
}

// Flush discards data written to the terminal but not transmitted, or
// received from the terminal but not read, depending on the queue:
//
// TCIFLUSH   received but not read.
// TCOFLUSH   written but not transmitted.
// TCIOFLUSH  both
func Flush(fd uintptr, queue int) error {
	return ioctlu(fd, unix.TCFLSH, uintptr(queue))
}

// Flow manages the suspending of data transmission or reception on the
// terminal referenced by fd.  The value of action must be one of the
// following:
// TCOOFF  Suspend output.
// TCOON   Restart suspended output.
// TCIOFF  Transmit a STOP character (the XOFF in XON/XOFF)
// TCION   Transmit a START character (the XON).
func Flow(fd uintptr, action int) error {
	return ioctlu(fd, unix.TCXONC, uintptr(action))
}
