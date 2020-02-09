// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package ble

import "github.com/embeddedgo/nrf5/softdevice/s140/sd"

const ErrorBase = 0x3000

type Error uint32

const (
	ErrNotEnabled          Error = ErrorBase + 1 // BLE not enabled
	ErrInvalidConnHandle   Error = ErrorBase + 2 // invalid connection handle
	ErrInvalidAttrHandle   Error = ErrorBase + 3 // invalid connection handle
	ErrInvalidAdvHandle    Error = ErrorBase + 4 // invalid advertising handle
	ErrInvalidRole         Error = ErrorBase + 5 // invalid advertising handle
	ErrBlockedByOtherLinks Error = ErrorBase + 6 // change link settings failed due to the scheduling of other links
)

var errStr = [...]string{
	"???",
	"BLE not enabled",
	"invalid connection handle",
	"invalid connection handle",
	"invalid advertising handle",
	"invalid advertising handle",
	"blocked by other links",
}

func (e Error) Error() string {
	s := "???"
	if e >= ErrorBase {
		if e -= ErrorBase; int(e) < len(errStr) {
			s = errStr[e]
		}
	}
	return "ble: " + s
}

func mkerr(e uint32) error {
	switch {
	case e == 0:
		return nil
	case e >= ErrorBase:
		return Error(e)
	default:
		return sd.Error(e)
	}
}
