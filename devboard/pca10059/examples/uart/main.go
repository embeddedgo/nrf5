// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"

	"github.com/embeddedgo/nrf5/hal/clock"
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/uart"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/init"
)

var tts *uart.Driver

func main() {
	clock.StoreTRACECONFIG(clock.T4MHz, clock.Serial) // enable SWO on P1.00

	p1 := gpio.P(1)
	rxpin := p1.Pin(10)
	txpin := p1.Pin(13)

	rxpin.Setup(gpio.ModeIn)
	txpin.Set()
	txpin.Setup(gpio.ModeOut)

	tts = uart.NewDriver(uart.UART(0), make([]byte, 256))
	tts.UsePin(uart.RXD, rxpin)
	tts.UsePin(uart.TXD, txpin)
	tts.IRQ().Enable(rtos.IntPrioLow)
	tts.SetBaudrate(uart.Baud1M)
	tts.Enable()
	tts.EnableRx()
	tts.EnableTx()

	buf := make([]byte, 128)
	for {
		n, err := tts.Read(buf)
		if n != 0 {
			tts.WriteString("inp:  ")
			tts.Write(buf[:n])
			tts.WriteString("\r\n")
		}
		if err != nil {
			tts.WriteString("err: " + err.Error() + "\r\n")
		}
	}
}

//go:interrupthandler
func UARTE0_UART0_Handler() { tts.ISR() }
