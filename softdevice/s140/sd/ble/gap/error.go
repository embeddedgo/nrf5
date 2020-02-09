// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package gap

import (
	"github.com/embeddedgo/nrf5/softdevice/s140/sd"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"
)

const ErrorBase = ble.ErrorBase + 0x200

type Error uint32

const (
	// UUID list does not contain an integral number of UUIDs
	ErrUUIDListMismatch Error = ErrorBase + 0

	// Use of whitelist not permitted with discoverable advertising.
	ErrDiscoverableWithWhitelist Error = ErrorBase + 1

	// The upper two bits of the address do not correspond to the specified
	// address type.
	ErrInvalidBLEAddr Error = ErrorBase + 2

	// Attempt to modify the whitelist while already in use by another
	// operation.
	ErrWhitelistInUse Error = ErrorBase + 3

	// Attempt to modify the device identity list while already in use by
	// another operation.
	ErrDeviceIdentitiesInUse Error = ErrorBase + 4

	// The device identity list contains entries with duplicate identity
	// addresses.
	ErrDeviceIdentitiesDuplicate Error = ErrorBase + 5
)

var errStr = [...]string{
	"UUID list mismatch",
	"discoverable with whitelist",
	"invalid BLE addr",
	"whitelist in use",
	"device identities in use",
	"device identities duplicate",
}

func (e Error) Error() string {
	s := "???"
	if e >= ErrorBase {
		if e -= ErrorBase; int(e) < len(errStr) {
			s = errStr[e]
		}
	}
	return "gap: " + s
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
