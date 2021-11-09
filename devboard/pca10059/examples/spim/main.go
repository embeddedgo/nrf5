// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example tests SPIM peripheral. Make a loop between MISO and MOSI.
// Observe the SWO output. Try disconnect and connect back the loop a few times.
package main

import (
	"embedded/rtos"
	"fmt"

	"github.com/embeddedgo/nrf5/hal/clock"
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/spim"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/init"
)

var drv *spim.Driver

func main() {
	clock.StoreTRACECONFIG(clock.T4MHz, clock.Serial) // enable SWO on P1.00

	println("START!")
	p0 := gpio.P(0)
	sck := p0.Pin(13) // required to make SPIM working, even if we don't use it
	miso := p0.Pin(15)
	mosi := p0.Pin(17)

	drv = spim.NewDriver(spim.SPIM(0))
	drv.UsePin(sck, spim.SCK)
	drv.UsePin(miso, spim.MISO)
	drv.UsePin(mosi, spim.MOSI)
	drv.Setup(spim.CPOL0|spim.CPHA0, spim.F1MHz)
	drv.Periph().StoreORC('.')
	drv.Enable()
	drv.IRQ().Enable(rtos.IntPrioLow, 0)

	txbuf := []byte("1234567890abcdefghijklmnoprstuvwxyz")
	rxbuf := make([]byte, len(txbuf)+3)

	for {
		for i := range rxbuf {
			rxbuf[i] = ' '
		}
		n := drv.WriteRead(txbuf, rxbuf)
		println(fmt.Sprintf("%d: %s", n, rxbuf))
	}
}

//go:interrupthandler
func SPIM0_SPIS0_TWIM0_TWIS0_SPI0_TWI0_Handler() {
	drv.ISR()
}
