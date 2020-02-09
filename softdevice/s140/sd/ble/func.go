// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package ble

// Enable enables and initializes the BLE stack.
func Enable(ramBase uintptr) (minBase uintptr, err error) {
	e := enable(&ramBase)
	return ramBase, mkerr(e)
}

// AddVSUUID add a Vendor Specific base UUID.
func AddVSUUID(vsuuid *UUID128) (t UUIDType, err error) {
	err = mkerr(addUUIDVS(vsuuid, &t))
	return
}

// EncodeUUID encode UUID bytes into provided buffer ant returns its size in
// bytes. Set buf to nil to only compute encoded UUID size.
func EncodeUUID(uuid *UUID, buf []byte) (int, error) {
	var p *byte
	if buf != nil {
		if len(buf) < 2 || len(buf) < 16 && uuid.Type > UUID16 {
			panic("ble: buffer too short for UUID")
		}
		p = &buf[0]
	}
	var n uint8
	e := encodeUUID(uuid, &n, p)
	return int(n), mkerr(e)
}

func DataBytes(data []byte) Data {
	if len(data) == 0 {
		return Data{}
	}
	if len(data) > 0xFFFF {
		panic("ble: data too long")
	}
	return Data{&data[0], uint16(len(data))}
}

// GetEvt gets an event bytes from the pending events queue. It copies event
// data to dest and returns its size (in bytes). Set dest to nil to only check
// the event size.
func GetEvt(dest []uint32) (uintptr, error) {
	n := uint16(len(dest) * 4)
	var p *uint32
	if dest != nil {
		p = &dest[0]
	}
	e := getEvt(p, &n)
	return uintptr(n), mkerr(e)
}
