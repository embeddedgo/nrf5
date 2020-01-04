// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uart0

import (
	"embedded/rtos"
	_ "unsafe"

	"github.com/embeddedgo/nrf5/hal/uart"
)

var driver *uart.Driver

// Driver returns ready to use driver for UART.
// information.
func Driver() *uart.Driver {
	if driver == nil {
		driver = uart.NewDriver(uart.UART(0))
		driver.IRQ().Enable(rtos.IntPrioLow)
	}
	return driver
}

//go:interrupthandler
func _UARTE0_UART0_Handler() { driver.ISR() }

//go:linkname _UARTE0_UART0_Handler IRQ2_Handler
