package certmgr

import (
	"errors"
)

var (
	ErrNotImplemented = errors.New("not implemented function")
	ErrUnsupported    = errors.ErrUnsupported
)
