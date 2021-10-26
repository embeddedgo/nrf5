// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2019 Michal Derkacz. All rights reserved.

package gap

import (
	"unsafe"

	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble/internal"
)

func SetConnCfg(cfg *ConnCfg, ramBase uintptr) error {
	return mkerr(internal.SetCfg(internal.CONN_CFG_GAP, unsafe.Pointer(cfg),
		ramBase))
}

// SetRoleCountCfg sets the role count configuration.
func SetRoleCountCfg(cfg *RoleCountCfg, ramBase uintptr) error {
	return mkerr(internal.SetCfg(internal.GAP_CFG_ROLE_COUNT, unsafe.Pointer(cfg), ramBase))
}

// SetDeviceName sets device name and security mode.
func SetDeviceName(writePerm ConnSecMode, devName string) error {
	return mkerr(setDeviceName(&writePerm, devName))
}

// SetPPCP sets Peripheral Preferred Connection Parameters.
func SetPPCP(connParams *ConnParams) error {
	return mkerr(setPPCP(connParams))
}

// ConfigureAdvSet configures an advertising set.
func ConfigureAdvSet(handle AdvSetHandle, data *AdvData, params *AdvParams) (AdvSetHandle, error) {
	e := configureAdvSet(&handle, data, params)
	return handle, mkerr(e)
}

// StartAdv starts advertising.
func StartAdv(advHandle AdvSetHandle, connCfgTag uint8) error {
	return mkerr(startAdv(advHandle, connCfgTag))
}

// StopAdv stops advertising.
func StopAdv(advHandle AdvSetHandle) error {
	return mkerr(stopAdv(advHandle))
}

// UpdateConnParam updates connection parameters.
func UpdateConnParam(connHandle ble.Handle, connParams *ConnParams) error {
	return mkerr(updateConnParam(connHandle, connParams))
}

// UpdateDataLength initiated or respondd to a Data Length Update Procedure.
func UpdateDataLength(connHandle ble.Handle, dlParams *DataLengthParams, dlLimitation *DataLengthLimitation) error {
	return mkerr(updateDataLength(connHandle, dlParams, dlLimitation))
}
