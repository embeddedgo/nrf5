// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"

	"github.com/embeddedgo/nrf5/devboard/pca10059/board/buttons"
	"github.com/embeddedgo/nrf5/devboard/pca10059/board/leds"
	"github.com/embeddedgo/nrf5/hal/gpiote"
	"github.com/embeddedgo/nrf5/hal/ppi"
)

func main() {
	led := gpiote.ChanAlloc()
	led.Setup(leds.User.Pin(), gpiote.ModeTask|gpiote.PolarityToggle)

	btn := gpiote.ChanAlloc()
	btn.Setup(buttons.User.Pin(), gpiote.ModeEvent|gpiote.PolarityHiToLo)

	btnled := ppi.ChanAlloc()
	btnled.SetTEP(led.OUT().Task())
	btnled.SetEEP(btn.IN().Event())
	btnled.Enable()

	for {
		leds.Blue.SetOff()
		leds.Red.SetOn()
		rtos.Nanosleep(500e6)

		leds.Red.SetOff()
		leds.Green.SetOn()
		rtos.Nanosleep(500e6)

		leds.Green.SetOff()
		leds.Blue.SetOn()
		rtos.Nanosleep(500e6)
	}
}
