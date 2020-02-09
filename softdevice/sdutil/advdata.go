// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdutil

import (
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble/gap"
)

type AdvData struct {
	buf []byte
}

func NewAdvData(maxLen int) *AdvData {
	return &AdvData{make([]byte, 0, maxLen)}
}

func (d *AdvData) Data() ble.Data {
	return ble.DataBytes(d.buf)
}

func (d *AdvData) Reset() {
	d.buf = d.buf[:0]
}

func (d *AdvData) AppendBytes(t gap.AdvDataType, b ...byte) {
	n := 1 + len(b)
	i := len(d.buf)
	newLen := i + n + 1
	if newLen > cap(d.buf) {
		panic("BLE data too long")
	}
	d.buf = d.buf[:newLen]
	d.buf[i] = byte(n)
	d.buf[i+1] = byte(t)
	copy(d.buf[i+2:], b)
}

func (d *AdvData) AppendString(t gap.AdvDataType, s string) {
	d.AppendBytes(t, []byte(s)...)
}

func (d *AdvData) AppendUUIDs(t gap.AdvDataType, size int, uuids ...*ble.UUID) {
	n := 0
	for _, uuid := range uuids {
		m, err := ble.EncodeUUID(uuid, nil)
		if err != nil {
			panic(err.Error())
		}
		if m == size {
			n += m
		}
	}
	if n == 0 {
		return
	}
	n += 1
	i := len(d.buf)
	newLen := i + n + 1
	if newLen > cap(d.buf) {
		panic("BLE data too long")
	}
	d.buf = d.buf[:newLen]
	d.buf[i] = byte(n)
	d.buf[i+1] = byte(t)
	i += 2
	for _, uuid := range uuids {
		if m, _ := ble.EncodeUUID(uuid, nil); m != size {
			continue
		}
		m, _ := ble.EncodeUUID(uuid, d.buf[i:])
		i += m
	}
}
