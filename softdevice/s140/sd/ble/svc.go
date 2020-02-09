// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package ble

//go:noescape
func enable(app_ram_base *uintptr) uint32

//go:noescape
func addUUIDVS(vs_uuid *UUID128, uuid_type *UUIDType) uint32

//go:noescape
func encodeUUID(uuid *UUID, uuid_le_len *uint8, uuid_le *uint8) uint32

//go:noescape
func getEvt(dest *uint32, len *uint16) uint32
