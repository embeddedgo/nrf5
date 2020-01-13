// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpiote

import (
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/internal"
	"github.com/embeddedgo/nrf5/hal/te"
)

// ChanNum is the number of implemented GPIOTE channels (4 in case of nRF51, 8
// in case of nRF52).
const ChanNum = chanNum

// Chan represents GPIOTE channel.
type Chan int8

var unusedChannels uint32 = 1<<chanNum - 1

// ChanAlloc returns first unused GPIOTE channel. It returns negative number if
// there is no free channels.
func ChanAlloc() Chan {
	return Chan(internal.BitAlloc(&unusedChannels))
}

// Free disables channel and returns it to the pool of unused channels.
func (c Chan) Free() {
	internal.BitFree(&unusedChannels, 1<<c)
}

type Task byte

// OUT returns task for writing to pin associated with channel c.
func (c Chan) OUT() Task {
	return Task(c)
}

// SET returns task for set pin associated with channel c. nRF52.
func (c Chan) SET() Task {
	return Task(c + 12)
}

// CLR returns task for clear pin associated with channel c. nRF52.
func (c Chan) CLR() Task {
	return Task(c + 24)
}

type Event byte

const PORT Event = 31 // From multiple input pins with SENSE mechanism enabled.

// IN returns event generated from pin associated with channel c.
func (c Chan) IN() Event {
	return Event(c)
}

func (t Task) Task() *te.Task    { return r().Regs.Task(int(t)) }
func (e Event) Event() *te.Event { return r().Regs.Event(int(e)) }

type Config uint32

const (
	ModeDiscon Config = 0 // Disconnect pin from GPIOTE module.
	ModeEvent  Config = 1 // Pin generates IN event.
	ModeTask   Config = 3 // Pin controlled by OUT, SET, CLR task.

	PolarityNone   Config = 0 << 16 // No task on pin, no event from pin.
	PolarityLoToHi Config = 1 << 16 // OUT sets pin, IN when rising edge.
	PolarityHiToLo Config = 2 << 16 // OUT clears pin, IN when falling edge.
	PolarityToggle Config = 3 << 16 // OUT toggles pin, IN when any change.

	OutInitLow  Config = 0       // Low initial output value.
	OutInitHigh Config = 1 << 20 // High initial output value.
)

// Config returns current configuration of channel c.
func (c Chan) Config() (gpio.Pin, Config) {
	v := r().config[c].Load()
	const psel = 0x7F << 8
	return gpio.PSEL(v & psel >> 8).Pin(), Config(v &^ psel)
}

// Setup setups channel c to use pin and cfg configuration.
func (c Chan) Setup(pin gpio.Pin, cfg Config) {
	r().config[c].Store(uint32(pin.PSEL())<<8 | uint32(cfg))
}
