// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package timer provides interface to manage nRF5 timers.
package timer

import (
	"embedded/mmio"

	"github.com/embeddedgo/nrf5/hal/te"
)

// Periph represents timer/counter peripheral.
type Periph struct {
	te.Regs

	_         [65]uint32
	mode      mmio.U32 // timer mode selection
	bitmode   mmio.U32 // number of bits used
	_         uint32
	prescaler mmio.U32
	_         [11]uint32
	cc        [6]mmio.U32 // capture/compare registers.
}

// TIMER returns n-th instance of TIMER peripheral or nil if this instance is
// not implemented.
func TIMER(n int) *Periph { return timer(n) }

type Task byte

const (
	START Task = 0 // Start Timer.
	STOP  Task = 1 // Stop Timer.
	COUNT Task = 2 // Increment Timer (Counter mode only).
	CLEAR Task = 3 // Clear timer.
)

// CAPTURE returns Capture task for CCn register.
func CAPTURE(n int) Task {
	return Task(16 + n)
}

type Event byte

// COMPARE returns Compare event for CCn register.
func COMPARE(n int) Event {
	return Event(16 + n)
}

func (p *Periph) Task(t Task) *te.Task    { return p.Regs.Task(int(t)) }
func (p *Periph) Event(e Event) *te.Event { return p.Regs.Event(int(e)) }

type Shorts uint32

const (
	COMPARE0_CLEAR Shorts = 1 << 0
	COMPARE1_CLEAR Shorts = 1 << 1
	COMPARE2_CLEAR Shorts = 1 << 2
	COMPARE3_CLEAR Shorts = 1 << 3
	COMPARE4_CLEAR Shorts = 1 << 4
	COMPARE5_CLEAR Shorts = 1 << 5
	COMPARE0_STOP  Shorts = 1 << 8
	COMPARE1_STOP  Shorts = 1 << 9
	COMPARE2_STOP  Shorts = 1 << 10
	COMPARE3_STOP  Shorts = 1 << 11
	COMPARE4_STOP  Shorts = 1 << 12
	COMPARE5_STOP  Shorts = 1 << 13
)

func (p *Periph) LoadSHORTS() Shorts   { return Shorts(p.Regs.LoadSHORTS()) }
func (p *Periph) StoreSHORTS(s Shorts) { p.Regs.StoreSHORTS(uint32(s)) }

type Mode uint8

const (
	Timer           Mode = 0
	Counter         Mode = 1
	LowPowerCounter Mode = 2
)

func (p *Periph) LoadMODE() Mode {
	return Mode(p.mode.Load())
}

func (p *Periph) StoreMODE(m Mode) {
	p.mode.Store(uint32(m))
}

type Bitmode uint8

const (
	Bit8  Bitmode = 1
	Bit16 Bitmode = 0
	Bit24 Bitmode = 2
	Bit32 Bitmode = 3
)

func (p *Periph) LoadBITMODE() Bitmode {
	return Bitmode(p.bitmode.Load())
}

func (p *Periph) StoreBITMODE(m Bitmode) {
	p.bitmode.Store(uint32(m))
}

func (p *Periph) LoadPRESCALER() int {
	return int(p.prescaler.Load())
}

// StorePRESCALER sets prescaler to exp (freq = 16MHz/2^exp). Must only be used
// when the timer is stopped.
func (p *Periph) StorePRESCALER(exp int) {
	p.prescaler.Store(uint32(exp))
}

// LoadCC returns value of n-th Capture/Compare register. nRF51/nRF52 has 4/6 CC
// registers
func (p *Periph) LoadCC(n int) uint32 {
	return p.cc[n].Load()
}

// StoreCC stores cc into n-th Capture/Compare register. nRF51/nRF52 has 4/6 CC
// registers
func (p *Periph) StoreCC(n int, cc uint32) {
	p.cc[n].Store(cc)
}
