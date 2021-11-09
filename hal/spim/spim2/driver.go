// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spim0

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/nrf5/hal/spim"
)

var driver *spim.Driver

// Driver returns ready to use driver for SPIM2 peripheral.
func Driver() *spim.Driver {
	if driver == nil {
		driver = spim.NewDriver(spim.SPIM(2))
		driver.IRQ().Enable(rtos.IntPrioLow, 0)
	}
	return driver
}

//go:interrupthandler
func _SPIM2_SPIS2_SPI2_Handler() { driver.ISR() }

//go:linkname _SPIM2_SPIS2_SPI2_Handler IRQ35_Handler
