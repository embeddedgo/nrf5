// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package gatt

type CharProps uint8

const (
	PropBroadcast    CharProps = 1 << 0 // broadcasting permitted
	PropRead         CharProps = 1 << 1 // reading permitted
	PropWriteWoResp  CharProps = 1 << 2 // write command permitted
	PropWrite        CharProps = 1 << 3 // write request permitted
	PropNotify       CharProps = 1 << 4 // notification permitted
	PropIndicate     CharProps = 1 << 5 // indications permitted
	PropAuthSignedWr CharProps = 1 << 6 // signed write command permitted
)

type CharExtProps uint8

const (
	PropRealiableWr CharExtProps = 1 << 0 // queued write permitted
	PropWrAux       CharExtProps = 1 << 1 // write Char.User Descr. descriptor permitted
)

// Presentation Format
type PF uint8

const (
	FormatRFU     PF = 0x00 // reserved for future use
	FormatBoolean PF = 0x01 // boolean
	Format2bit    PF = 0x02 // unsigned 2-bit integer
	FormatNibble  PF = 0x03 // unsigned 4-bit integer
	FormatUint8   PF = 0x04 // unsigned 8-bit integer
	FormatUint12  PF = 0x05 // unsigned 12-bit integer
	FormatUint16  PF = 0x06 // unsigned 16-bit integer
	FormatUint24  PF = 0x07 // unsigned 24-bit integer
	FormatUint32  PF = 0x08 // unsigned 32-bit integer
	FormatUint48  PF = 0x09 // unsigned 48-bit integer
	FormatUint64  PF = 0x0A // unsigned 64-bit integer
	FormatUint128 PF = 0x0B // unsigned 128-bit integer
	FormatSint8   PF = 0x0C // signed 8-bit integer
	FormatSint12  PF = 0x0D // signed 12-bit integer
	FormatSint16  PF = 0x0E // signed 16-bit integer
	FormatSint24  PF = 0x0F // signed 24-bit integer
	FormatSint32  PF = 0x10 // signed 32-bit integer
	FormatSint48  PF = 0x11 // signed 48-bit integer
	FormatSint64  PF = 0x12 // signed 64-bit integer
	FormatSint128 PF = 0x13 // signed 128-bit integer
	FormatFloat32 PF = 0x14 // IEEE-754 32-bit floating point
	FormatFloat64 PF = 0x15 // IEEE-754 64-bit floating point
	FormatSFloat  PF = 0x16 // IEEE-11073 16-bit SFLOAT
	FormatFloat   PF = 0x17 // IEEE-11073 32-bit FLOAT
	FormatDUint16 PF = 0x18 // IEEE-20601 format
	FormatUTF8S   PF = 0x19 // UTF-8 string
	FormatUTF16S  PF = 0x1A // UTF-16 string
	FormatStruct  PF = 0x1B // opaque Structure
)

type HVXType uint8

const (
	HVXInvalid   HVXType = 0 // invalid operation
	Notification HVXType = 1 // handle value notification
	Indication   HVXType = 2 // handle value indication
)
