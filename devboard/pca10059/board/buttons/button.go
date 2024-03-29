// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buttons

import (
	"github.com/embeddedgo/nrf5/hal/gpio"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/system"
)

// Onboard buttons
const (
	SW1 Button = 0x16 // P1.06

	User = SW1
)

type Button uint8

func (b Button) prt() int  { return int(b) >> 4 }
func (b Button) pin() uint { return uint(b) & 15 }

func (b Button) Read() int {
	return int(gpio.P(b.prt()).Load()>>b.pin())&1 ^ 1
}
func (b Button) Pin() gpio.Pin {
	return gpio.P(b.prt()).Pin(int(b.pin()))
}

func init() {
	SW1.Pin().Setup(gpio.ModeIn | gpio.PullUp)
}
