// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdutil

import (
	"errors"
	"unsafe"

	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble/gap"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble/gatt/gatts"
)

//func bleEvtWordsMax(attmtu int) int {
//	bytes := 12 + (attmtu-1)/4*8 // based on BLE_EVT_LEN_MAX(ATT_MTU)
//	return (bytes + 3) / 4
//}

// NextEvent returns next available softdevice event. Internally it calls
// ble.GetEvt twice: first to determine the event size and next to dequeue it.
// For this reason it cannot be used concurently by multiple gorutines.
func NextEvent() (Event, error) {
	n, err := ble.GetEvt(nil)
	if err != nil {
		return nil, err
	}
	if n < 6 {
		return nil, errors.New("sdutil: event too short")
	}
	buf := make([]uint32, (n+3)/4)
	n, err = ble.GetEvt(buf)
	if err != nil {
		return nil, err
	}
	p := unsafe.Pointer(&buf[0])
	//print("\n[ ")
	//for _, b := range (*[1 << 16]byte)(p)[:n] {
	//	print(b)
	//	print(" ")
	//}
	//print("]\n")
	hdr := (*ble.EvtHdr)(p)
	if uintptr(hdr.Len) != n {
		return nil, errors.New("sdutil: bad event data size")
	}
	switch {
	case ble.EvtBase <= hdr.ID && hdr.ID <= ble.EvtLast:
		switch hdr.ID {
		case ble.EvtUserMemRequest:
			return (*ble.UserMemRequest)(p), nil
		case ble.EvtUserMemRelease:
			return (*ble.UserMemRelease)(p), nil
		}
	case gap.EvtBase <= hdr.ID && hdr.ID <= gap.EvtLast:
		switch hdr.ID {
		case gap.EvtConnected:
			return (*gap.Connected)(p), nil
		case gap.EvtDisconnected:
			return (*gap.Disconnected)(p), nil
		case gap.EvtConnParamUpdate:
			return (*gap.ConnParamUpdate)(p), nil
		}
	case gatts.EvtBase <= hdr.ID && hdr.ID <= gatts.EvtLast:
		switch hdr.ID {
		case gatts.EvtWrite:
			return (*gatts.Write)(p), nil
		}
	}
	return hdr, nil // unknown event
}

type Event interface {
	EventID() uint16
}
