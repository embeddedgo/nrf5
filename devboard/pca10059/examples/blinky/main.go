// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/embeddedgo/x/time"

	"github.com/embeddedgo/nrf5/devboard/pca10059/board/buttons"
	"github.com/embeddedgo/nrf5/devboard/pca10059/board/leds"
	"github.com/embeddedgo/nrf5/hal/clock"
)

func delay() {
	if buttons.User.Read() != 0 {
		time.Sleep(time.Second / 8)
	} else {
		time.Sleep(time.Second / 2)
	}
}

func main() {
	clock.StoreTRACECONFIG(clock.T4MHz, clock.Serial) // enable SWO on P1.00
	for {
		leds.Blue.SetOff()
		leds.User.SetOn()
		println("user")
		delay()

		leds.User.SetOff()
		leds.Red.SetOn()
		println("red")
		delay()

		leds.Red.SetOff()
		leds.Green.SetOn()
		println("green")
		delay()

		leds.Green.SetOff()
		leds.Blue.SetOn()
		println("blue")
		delay()
	}
}
