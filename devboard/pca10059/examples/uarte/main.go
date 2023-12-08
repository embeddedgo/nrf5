// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to use the UARTE peripheral.
package main

import (
	"time"

	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/uarte"
	"github.com/embeddedgo/nrf5/hal/uarte/uarte1"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/system"
)

func main() {
	p1 := gpio.P(1)
	rxpin := p1.Pin(10)
	txpin := p1.Pin(13)

	tts := uarte1.Driver()
	tts.UsePin(rxpin, uarte.RXD)
	tts.UsePin(txpin, uarte.TXD)
	tts.SetBaudrate(uarte.Baud115200)
	tts.Enable()

	for {
		tts.Write([]byte("Hello World!"))
		tts.WriteByte('\n')
		tts.WriteByte('\r')
		time.Sleep(time.Second)
	}
}
