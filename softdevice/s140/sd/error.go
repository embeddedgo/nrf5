// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package sd

type Error uint32

const (
	ErrSVC           Error = 1  // SVC handler is missing
	ErrNotEnabled    Error = 2  // softdevice not enabled
	ErrInternal      Error = 3  // internal error
	ErrNoMem         Error = 4  // no memory for operation
	ErrNotFound      Error = 5  // not found
	ErrNotSupported  Error = 6  // not supported
	ErrInvalidParam  Error = 7  // invalid parameter
	ErrInvalidState  Error = 8  // invalid state, operation disallowed in this state
	ErrInvalidLength Error = 9  // invalid length
	ErrInvalidFlags  Error = 10 // invalid flags
	ErrInvalidData   Error = 11 // invalid data
	ErrDataSize      Error = 12 // invalid data size
	ErrTimeout       Error = 13 // operation timed out
	ErrNil           Error = 14 // nil pointer
	ErrForbidden     Error = 15 // forbidden operation
	ErrInvalidAddr   Error = 16 // bad memory address
	ErrBusy          Error = 17 // busy
	ErrConnCount     Error = 18 // maximum connection count exceeded
	ErrResources     Error = 19 // not enough resources for operation
)

var errStr = [...]string{
	"???",
	"SVC handler is missing",
	"softdevice not enabled",
	"internal error",
	"no memory for operation",
	"not found",
	"not supported",
	"invalid parameter",
	"invalid state",
	"invalid length",
	"invalid flags",
	"invalid data",
	"invalid data size",
	"operation timed out",
	"nil pointer",
	"forbidden operation",
	"bad memory address",
	"busy",
	"maximum connection count exceeded",
	"not enough resources for operation",
}

func (e Error) Error() string {
	s := "???"
	if int(e) < len(errStr) {
		s = errStr[e]
	}
	return "sd: " + s
}
