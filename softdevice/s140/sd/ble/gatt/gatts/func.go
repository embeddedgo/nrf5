// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package gatts

import (
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"
)

// AddService adds a service.
func AddService(typ ServiceType, uuid *ble.UUID) (h Handle, err error) {
	err = mkerr(addService(typ, uuid, &h))
	return
}

func AddCharacteristic(service Handle, charMeta *CharMeta, charVal *Attr) (handles CharHandles, err error) {
	err = mkerr(addCharacteristic(service, charMeta, charVal, &handles))
	return
}

// HVX handles value notification or indication. If len(data) != 0 the
// attribute will be updated before sending the notification or indication. HVX
// returns the number of bytes written to attribute.
func HVX(connHandle ble.Handle, params *HVXParams, data ...byte) (int, error) {
	p := hvxParams{params.Handle, params.Type, params.Offset, nil, nil}
	if len(data) >= 1<<16 {
		panic("gatts: HVX data too long")
	}
	n := uint16(len(data))
	if n != 0 {
		if params.Len != 0 && n > params.Len {
			n = params.Len
		}
		p.len = &n
		p.data = &data[0]
	}
	e := hvx(connHandle, &p)
	return int(n), mkerr(e)
}

type SysAttrFlags uint32

const (
	SysAttrSys SysAttrFlags = 1 << 0 // restrict system attributes to system services
	SysAttrUsr SysAttrFlags = 1 << 1 // restrict system attributes to user services
)

// SetSysAttr sets the persistent system attributes for a connection.
func SetSysAttr(connHandle ble.Handle, data []byte, flags SysAttrFlags) error {
	var p *byte
	if len(data) != 0 {
		p = &data[0]
	}
	return mkerr(setSysAttr(connHandle, p, uint16(len(data)), flags))
}

// GetSysAttr retrieves the persistent system attributes.
func GetSysAttr(connHandle ble.Handle, data []byte, flags SysAttrFlags) (int, error) {
	var p *byte
	var n uint16
	if len(data) != 0 {
		p = &data[0]
		n = uint16(len(data))
	}
	e := getSysAttr(connHandle, p, &n, flags)
	return int(n), mkerr(e)
}
