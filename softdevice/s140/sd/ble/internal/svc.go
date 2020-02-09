// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package internal

import "unsafe"

const CONN_CFG_BASE = 0x20

const (
	CONN_CFG_GAP   = CONN_CFG_BASE + 0 // BLE GAP specific connection configuration.
	CONN_CFG_GATTC = CONN_CFG_BASE + 1 // BLE GATTC specific connection configuration.
	CONN_CFG_GATTS = CONN_CFG_BASE + 2 // BLE GATTS specific connection configuration.
	CONN_CFG_GATT  = CONN_CFG_BASE + 3 // BLE GATT specific connection configuration.
	CONN_CFG_L2CAP = CONN_CFG_BASE + 4 // BLE L2CAP specific connection configuration.
)

const GAP_CFG_BASE = 0x40

const (
	GAP_CFG_ROLE_COUNT       = GAP_CFG_BASE + 0 // Role count configuration.
	GAP_CFG_DEVICE_NAME      = GAP_CFG_BASE + 1 // Device name configuration.
	GAP_CFG_PPCP_INCL_CONFIG = GAP_CFG_BASE + 2 // Peripheral Preferred Connection Parameters characteristic inclusion configuration.
	GAP_CFG_CAR_INCL_CONFIG  = GAP_CFG_BASE + 3 // Central Address Resolution characteristic inclusion configuration.
)

//go:noescape
func SetCfg(cfg_id uint32, cfg unsafe.Pointer, app_ram_base uintptr) uint32
