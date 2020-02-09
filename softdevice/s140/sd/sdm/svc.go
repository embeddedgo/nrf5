// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package sdm

import "unsafe"

//go:noescape
func enableSoftdevice(lfclkc *LFCLKC) uint32

func disableSoftdevice() uint32

//go:noescape
func softdeviceEnabled(en *bool) uint32

//escape
func setSoftdeviceVectorTableBase(addr unsafe.Pointer) uint32
