// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"time"

	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/gpiote"
	"github.com/embeddedgo/nrf5/hal/te"
)

type Button struct {
	C     <-chan bool
	Event *te.Event

	pin      gpio.Pin
	note     rtos.Note
	released int
}

func NewButton(pin gpio.Pin, released int) *Button {
	ch := gpiote.Alloc()
	ch.Setup(pin, gpiote.ModeEvent|gpiote.PolarityToggle)
	c := make(chan bool, 3)
	b := &Button{
		C:        c,
		Event:    ch.IN().Event(),
		pin:      pin,
		released: released,
	}
	go b.run(c)
	return b
}

func (b *Button) wait(state int, long bool) bool {
	for {
		b.note.Clear()
		b.Event.EnableIRQ()
		timeout := time.Duration(-1)
		if b.pin.Load() == state {
			timeout = 50 * time.Millisecond // we need 50 ms of stable state
		}
		if !b.note.Sleep(timeout) {
			if long {
				// wait +1 second to detect the long state
				long = !b.note.Sleep(time.Second)
			}
			b.Event.DisableIRQ()
			b.Event.Clear()
			return long
		}
	}
}

func (b *Button) run(c chan<- bool) {
	for {
		b.wait(b.released, false)
		c <- b.wait(b.released^1, true)
	}
}

func (b *Button) ISR() {
	b.Event.DisableIRQ()
	b.Event.Clear()
	b.note.Wakeup()
}
