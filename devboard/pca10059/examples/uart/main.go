// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"

	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/uart"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/init"
)

var tts *uart.Driver

func main() {
	p0 := gpio.P(0)
	rxpin := p0.Pin(22)
	txpin := p0.Pin(24)

	rxpin.Setup(gpio.ModeIn)
	txpin.Set()
	txpin.Setup(gpio.ModeOut)

	tts = uart.NewDriver(uart.UART(0), make([]byte, 16))
	tts.UsePin(uart.RXD, rxpin)
	tts.UsePin(uart.TXD, txpin)
	tts.IRQ().Enable(rtos.IntPrioLow)
	tts.SetBaudrate(uart.Baud115200)
	tts.Enable()
	tts.EnableRx()
	tts.EnableTx()

	buf := make([]byte, 256)
	for {
		n, err := tts.Read(buf)
		if err == nil {
			tts.WriteString("in:  ")
			tts.Write(buf[:n])
		} else {
			tts.WriteString("err: " + err.Error())
		}
		tts.WriteString("\r\n")
	}
}

//go:interrupthandler
func UARTE0_UART0_Handler() { tts.ISR() }
