// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package ble

type Handle uint16

const (
	ConnInvalid Handle = 0xFFFF // invalid connection handle
	ConnAll     Handle = 0xFFFE // applies to all connection handles
)

type UUID128 [16]byte // little-endian UUID bytes

type UUIDType uint8

const (
	UUIDXX     UUIDType = 0 // unknown/invalid UUID type
	UUID16     UUIDType = 1 // bluetooth SIG 16-bit UUID
	VSUUIDBase UUIDType = 2 // vendor-spcific UUID types start at this index
)

type UUID struct {
	Value uint16
	Type  UUIDType
}

type Data struct {
	p   *byte
	len uint16
}

type MemType uint8

const (
	UserMemInvalid           MemType = 0 // invalid user memory types
	UserMemGATTSQueuedWrites MemType = 1 // User memory for GATTS queued writes
)
