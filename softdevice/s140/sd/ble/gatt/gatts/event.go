// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package gatts

import (
	"unsafe"

	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"
)

const (
	EvtBase = 0x50
	EvtLast = 0x6F
)

const (
	EvtWrite          = EvtBase + 0 // write operation performed
	EvtRWAuthorizeReq = EvtBase + 1 // read/write authorization request
	EvtSysAttrMissing = EvtBase + 2 // persistent sys. attr. access is pending
	EvtHVC            = EvtBase + 3 // handle value confirmation
	EvtSCConfirm      = EvtBase + 4 // service changed confirmation
	EvtExchangeMTUReq = EvtBase + 5 // exchange MTU request
	EvtTimeout        = EvtBase + 6 // peer failed to resp. to ATT req. in time
	EvtHVNTxComplete  = EvtBase + 7 // handle val. notification Tx complete
)

type Write struct {
	ble.EvtHdr
	Handle       Handle   // attribute handle
	UUID         ble.UUID // attribute UUID
	Op           Op       // type of write operation
	AuthRequired bool     // writing operation deferred due to auth. requirement
	Offset       uint16   // offset for the write operation
	len          uint16   // length of the received data
	data         [1]uint8 // received data
}

func (w *Write) Data() []byte {
	return (*[1 << 16]byte)(unsafe.Pointer(&w.data[0]))[:w.len:w.len]
}
