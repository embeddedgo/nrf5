// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package gatts

import (
	"github.com/embeddedgo/nrf5/softdevice/s140/sd"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"
)

const ErrorBase = ble.ErrorBase + 0x400

type Error uint32

const (
	ErrInvalidAttrType Error = ErrorBase + 0 // invalid attribute type
	ErrSysAttrMissing  Error = ErrorBase + 1 // system attributes missing
)

var errStr = [...]string{
	"invalid attribute type",
	"system attributes missing",
}

func (e Error) Error() string {
	s := "???"
	if e >= ErrorBase {
		if e -= ErrorBase; int(e) < len(errStr) {
			s = errStr[e]
		}
	}
	return "gatts: " + s
}

func mkerr(e uint32) error {
	switch {
	case e == 0:
		return nil
	case e >= ErrorBase:
		return Error(e)
	case e >= ble.ErrorBase:
		return ble.Error(e)
	default:
		return sd.Error(e)
	}
}
