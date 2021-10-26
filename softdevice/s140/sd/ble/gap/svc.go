// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package gap

import "github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"

//go:noescape
func setDeviceName(write_perm *ConnSecMode, devName string) uint32

//go:noescape
func setPPCP(conn_params *ConnParams) uint32

//go:noescape
func configureAdvSet(adv_handle *AdvSetHandle, adv_data *AdvData, adv_params *AdvParams) uint32

func startAdv(adv_handle AdvSetHandle, conn_cfg_tag uint8) uint32

func stopAdv(adv_handle AdvSetHandle) uint32

//go:noescape
func updateConnParam(conn_handle ble.Handle, conn_params *ConnParams) uint32

//go:noescape
func updateDataLength(conn_handle ble.Handle, dl_params *DataLengthParams, dl_limitation *DataLengthLimitation) uint32
