// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package gatts

import (
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble/gap"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble/gatt"
)

type Handle uint16

type ServiceType uint8

const (
	InvalidService   ServiceType = 0
	PrimaryService   ServiceType = 1
	SecondaryService ServiceType = 2
)

type AttrOptions uint8

const (
	VLenAttr  AttrOptions = 1 << 0 // variable length attribute
	ValLoc    AttrOptions = 3 << 1 // attribute value location
	ValStack  AttrOptions = 1 << 1 // attr. value located in stack memory
	ValUser   AttrOptions = 2 << 1 // attr. value located in user memory
	ReadAuth  AttrOptions = 1 << 3 // authoriz. requested on every read
	WriteAuth AttrOptions = 1 << 4 // authoriz. requested on every write
)

type AttrMeta struct {
	ReadPerm  gap.ConnSecMode
	WritePerm gap.ConnSecMode
	Options   AttrOptions
}

type Attr struct {
	UUID     *ble.UUID // attribute UUID.
	Meta     *AttrMeta // metadata
	InitLen  uint16    // initial attribute value length in bytes
	InitOffs uint16    // initial attribute value offset in bytes
	MaxLen   uint16    // maximum attribute value length in bytes
	Value    *byte     // pointer to the attribute data
}

type CharPF struct {
	Format    gatt.PF // format of the value
	Exponent  int8    // exponent for integer data types
	Unit      uint16  // unit from Bluetooth Assigned Numbers
	Namespace uint8   // namespace from Bluetooth Assigned Numbers
	Desc      uint16  // namespace description from Bluetooth Assigned Numbers
}

type CharMeta struct {
	Props           gatt.CharProps    // characteristic properties
	ExtProps        gatt.CharExtProps // characteristic extended properties.
	UserDesc        *byte             // pointer to a UTF-8 encoded string
	UserDescMaxSize uint16
	UserDescSize    uint16
	PF              *CharPF   // presentation format (CPF descriptor)
	UserDescMeta    *AttrMeta // user description descriptor metadata
	CCCDMeta        *AttrMeta // client char. conf. descriptor metadeta
	SCCDMeta        *AttrMeta // server char. conf. descriptor metadata
}

type CharHandles struct {
	Value    Handle // handle to the characteristic value.
	UserDesc Handle // handle to the user description descriptor
	CCCD     Handle // handle to the client char. conf. descriptor
	SCCD     Handle // handle to the server char. conf. descriptor
}

type Op uint8

const (
	OpInvalid            Op = 0 // invalid operation
	OpWriteReq           Op = 1 // write request
	OpWriteCmd           Op = 2 // write command
	OpSignWriteCmd       Op = 3 // signed write command
	OpPrepWriteReq       Op = 4 // prepare write request
	OpExecWriteReqCancel Op = 5 // cancel all prepared writes
	OpExecWriteReqNow    Op = 6 // immediately execute all prepared writes
)

type HVXParams struct {
	Handle Handle       // characteristic value handle
	Type   gatt.HVXType // indication or notification
	Offset uint16       // field offset within the attribute value
	Len    uint16       // field length in bytes (0 means variable length)
}

type hvxParams struct {
	handle Handle
	typ    gatt.HVXType
	offset uint16
	len    *uint16
	data   *byte
}
