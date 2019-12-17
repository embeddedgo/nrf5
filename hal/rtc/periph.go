// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rtc provides interface to nRF5 real time counters.
package rtc

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/internal/mmap"
	"github.com/embeddedgo/nrf5/hal/te"
)

// Periph represents Real Time Counter peripheral.
type Periph struct {
	te.Regs

	_         [65]mmio.U32
	counter   mmio.U32
	prescaler mmio.U32
	_         [13]mmio.U32
	cc        [4]mmio.U32
}

func RTC(n int) *Periph {
	switch n {
	case 0:
		return (*Periph)(unsafe.Pointer(mmap.APB_BASE + 0x0B000))
	case 1:
		return (*Periph)(unsafe.Pointer(mmap.APB_BASE + 0x11000))
	case 2:
		return (*Periph)(unsafe.Pointer(mmap.APB_BASE + 0x24000))
	default:
		return nil
	}
}

type Task byte

const (
	START      Task = 0 // Start RTC COUNTER.
	STOP       Task = 1 // Stop RTC COUNTER.
	CLEAR      Task = 2 // Clear RTC COUNTER.
	TRIGOVRFLW Task = 3 // Store COUNTER to 0xFFFFF0.
)

type Event byte

const (
	TICK   Event = 0 // Event on COUNTER increment.
	OVRFLW Event = 1 // Event on COUNTER overflow.
)

// COMPARE returns compare event on CC[n] match.
func COMPARE(n int) Event {
	return Event(16 + n)
}

func (p *Periph) Task(t Task) *te.Task    { return p.Regs.Task(int(t)) }
func (p *Periph) Event(e Event) *te.Event { return p.Regs.Event(int(e)) }

// LoadCOUNTER returns value of counter register.
func (p *Periph) LoadCOUNTER() uint32 {
	return p.counter.Load()
}

// LoadPRESCALER returns value of prescaler register.
func (p *Periph) LoadPRESCALER() uint32 {
	return p.prescaler.Load()
}

// StorePRESCALER stores prescaler to pr (freq = 32768Hz/(pr+1)). Must only be used
// when the timer is stopped.
func (p *Periph) StorePRESCALER(pr int) {
	p.prescaler.Store(uint32(pr))
}

// LoadCC returns value of n-th compare register.
func (p *Periph) LoadCC(n int) uint32 {
	return p.cc[n].Load()
}

// StoreCC stores n-th compare register to cc.
func (p *Periph) StoreCC(n int, cc uint32) {
	p.cc[n].Store(cc)
}
