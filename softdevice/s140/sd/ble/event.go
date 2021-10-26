// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package ble

const (
	EvtBase = 1
	EvtLast = 0x0F
)

const (
	EvtUserMemReq     = EvtBase + 0 // user memory request
	EvtUserMemRelease = EvtBase + 1 // user memory release
)

type EvtHdr struct {
	ID         uint16 // event id (eg. EvtUserMemReq, gap.EvtConnected)
	Len        uint16 // length in octets including this header
	ConnHandle Handle // connection handle
}

func (e *EvtHdr) EventID() uint16 { return e.ID }

type UserMemReq struct {
	EvtHdr
	Type MemType // user memory type
}

type UserMemBlock struct {
	Mem *uint8 // Pointer to the start of the user memory block.
	Len uint16 // Length in bytes of the user memory block.
}

type UserMemRelease struct {
	EvtHdr
	Type     MemType      // user memory type
	MemBlock UserMemBlock // user memory block
}
