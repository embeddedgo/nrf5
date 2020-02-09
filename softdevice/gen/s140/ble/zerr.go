// Code generated from ../nordic/s140/headers/ble_err.h; DO NOT EDIT.

package ble

import (
	"github.com/embeddedgo/nrf5/sd/s140/nrf"
)

const ERROR_NOT_ENABLED = (nrf.ERROR_STK_BASE_NUM + 0x001)            // @ref sd_ble_enable has not been called.
const ERROR_INVALID_CONN_HANDLE = (nrf.ERROR_STK_BASE_NUM + 0x002)    // Invalid connection handle.
const ERROR_INVALID_ATTR_HANDLE = (nrf.ERROR_STK_BASE_NUM + 0x003)    // Invalid attribute handle.
const ERROR_INVALID_ADV_HANDLE = (nrf.ERROR_STK_BASE_NUM + 0x004)     // Invalid advertising handle.
const ERROR_INVALID_ROLE = (nrf.ERROR_STK_BASE_NUM + 0x005)           // Invalid role.
const ERROR_BLOCKED_BY_OTHER_LINKS = (nrf.ERROR_STK_BASE_NUM + 0x006) // The attempt to change link settings failed due to the scheduling of other links.
const L2CAP_ERR_BASE = (nrf.ERROR_STK_BASE_NUM + 0x100)               // L2CAP specific errors.
const GAP_ERR_BASE = (nrf.ERROR_STK_BASE_NUM + 0x200)                 // GAP specific errors.
const GATTC_ERR_BASE = (nrf.ERROR_STK_BASE_NUM + 0x300)               // GATT client specific errors.
const GATTS_ERR_BASE = (nrf.ERROR_STK_BASE_NUM + 0x400)               // GATT server specific errors.
