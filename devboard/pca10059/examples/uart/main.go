// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to use the UART peripheral.
package main

import (
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/uart"
	"github.com/embeddedgo/nrf5/hal/uart/uart0"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/init"
)

func main() {
	p1 := gpio.P(1)
	rxpin := p1.Pin(10)
	txpin := p1.Pin(13)

	tts := uart0.Driver()
	tts.UsePin(rxpin, uart.RXD)
	tts.UsePin(txpin, uart.TXD)
	tts.SetBaudrate(uart.Baud115200)
	tts.Enable()
	tts.EnableRx(make([]byte, 128))

	tts.WriteString("Ready!\r\n")

	buf := make([]byte, 64)
	for {
		n, err := tts.Read(buf)
		if n != 0 {
			tts.WriteString("inp: ")
			tts.Write(buf[:n])
			tts.WriteString("\r\n")
		}
		if err != nil {
			tts.WriteString("err: " + err.Error() + "\r\n")
		}
	}
}
