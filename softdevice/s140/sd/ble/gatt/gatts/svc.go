// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package gatts

import "github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"

//go:noescape
func addService(typ ServiceType, uuid *ble.UUID, handle *Handle) uint32

// attr_char_value.Value can escape
func addCharacteristic(service_handle Handle, char_md *CharMeta, attr_char_value *Attr, handles *CharHandles) uint32

//go:noescape
func hvx(conn_handle ble.Handle, hvx_params *hvxParams) uint32
