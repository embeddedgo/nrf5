// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/embeddedgo/nrf5/devboard/pca10059/board/buttons"
	"github.com/embeddedgo/nrf5/devboard/pca10059/board/leds"
	"github.com/embeddedgo/nrf5/hal/gpiote"
	"github.com/embeddedgo/nrf5/hal/te"
)

func main() {
	led := gpiote.Alloc()
	led.Setup(leds.User.Pin(), gpiote.ModeTask|gpiote.PolarityToggle)

	btn := gpiote.Alloc()
	btn.Setup(buttons.User.Pin(), gpiote.ModeEvent|gpiote.PolarityHiToLo)

	btnled := te.AllocChan()
	led.OUT().Task().SetChan(btnled, true)
	btn.IN().Event().SetChan(btnled, true)
	btnled.Enable()

	for {
		leds.Blue.SetOff()
		leds.Red.SetOn()
		time.Sleep(500 * time.Millisecond)

		leds.Red.SetOff()
		leds.Green.SetOn()
		time.Sleep(500 * time.Millisecond)

		leds.Green.SetOff()
		leds.Blue.SetOn()
		time.Sleep(500 * time.Millisecond)
	}
}
