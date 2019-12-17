// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leds

import (
	"github.com/embeddedgo/nrf5/hal/gpio"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/init"
)

// Onboard LEDs
const (
	LD1  LED = 0x06 // P0.06
	LD2R LED = 0x08 // P0.08
	LD2G LED = 0x19 // P1.09
	LD2B LED = 0x0C // P0.12

	User  = LD1
	Red   = LD2R
	Green = LD2G
	Blue  = LD2B
)

type LED uint8

func (d LED) prt() int  { return int(d) >> 4 }
func (d LED) pin() uint { return uint(d) & 15 }

func (d LED) SetOn() {
	gpio.P(d.prt()).ClearPins(1 << d.pin())
}
func (d LED) SetOff() {
	gpio.P(d.prt()).SetPins(1 << d.pin())
}
func (d LED) Set(on int) {
	port := gpio.P(d.prt())
	if on&1 == 0 {
		port.SetPins(1 << d.pin())
	} else {
		port.ClearPins(1 << d.pin())
	}
}
func (d LED) Get() int {
	return int(gpio.P(d.prt()).LoadOut()>>d.pin())&1 ^ 1
}
func (d LED) Pin() gpio.Pin {
	return gpio.P(d.prt()).Pin(int(d.pin()))
}

func init() {
	cfg := gpio.ModeOut | gpio.DriveH0D1
	LD1.SetOff()
	LD2R.SetOff()
	LD2G.SetOff()
	LD2B.SetOff()
	LD1.Pin().Setup(cfg)
	LD2R.Pin().Setup(cfg)
	LD2G.Pin().Setup(cfg)
	LD2B.Pin().Setup(cfg)
}
