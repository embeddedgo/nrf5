// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uarte1

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/nrf5/hal/uarte"
)

var driver *uarte.Driver

// Driver returns ready to use driver for UART.
func Driver() *uarte.Driver {
	if driver == nil {
		driver = uarte.NewDriver(uarte.UARTE(1))
		driver.IRQ().Enable(rtos.IntPrioLow, -1)
	}
	return driver
}

//go:interrupthandler
func _UARTE1_Handler() { driver.ISR() }

//go:linkname _UARTE1_Handler IRQ40_Handler
