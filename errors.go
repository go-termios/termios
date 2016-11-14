package termios

import "errors"

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrInvalidAction  = errors.New("invalid action")
)
