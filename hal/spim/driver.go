// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spim

import (
	"embedded/mmio"
	"embedded/rtos"
	"runtime"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/p/mmap"
)

type Driver struct {
	p    *Periph
	done rtos.Note
}

// NewDriver returns a new driver for p.
func NewDriver(p *Periph) *Driver {
	return &Driver{p: p}
}

func (d *Driver) Periph() *Periph {
	return d.p
}

// Enable enables SPIM peripheral.
func (d *Driver) Enable() {
	d.p.StoreENABLE(true)
	d.p.Event(END).Clear()
}

// Disable disables SPIM peripheral.
func (d *Driver) Disable() {
	d.p.StoreENABLE(false)
	if uintptr(unsafe.Pointer(d.p)) == mmap.SPIM3_BASE {
		// see errata 195: SPIM3 continues to draw current after disable
		(*mmio.U32)(unsafe.Pointer(uintptr(0x4002F004))).Store(1)
	}
}

// UsePin configurs the specified pin to be used as signal s.
func (d *Driver) UsePin(pin gpio.Pin, s Signal) {
	if s == MISO {
		pin.Setup(gpio.ModeIn)
	} else {
		switch s {
		case SCK:
			if d.p.LoadCONFIG()&CPOL1 != 0 {
				pin.Set()
			} else {
				pin.Clear()
			}
		case MOSI:
			pin.Clear()
		case CSN:
			pin.Store(^d.p.LoadCSNPOL())
		default: // DCX
			pin.Set()
		}
		pin.Setup(gpio.ModeOut)
	}
	d.p.StorePSEL(s, pin.PSEL(), true)
}

// Setup sets the SPI mode and clock frequency.
func (d *Driver) Setup(cfg Config, f Freq) {
	// TODO: set gpio output for SCK according to CPOL/CPHA
	d.p.StoreCONFIG(cfg)
	d.p.StoreFREQUENCY(f)
}

// IRQ returns the interrupt assigned to UART peripheral used by driver.
func (d *Driver) IRQ() rtos.IRQ {
	return d.p.IRQ()
}

func (d *Driver) ISR() {
	d.p.Event(END).DisableIRQ()
	d.done.Wakeup()
}

const minIRQLen = 8

func (d *Driver) WriteRead(out, in []byte) int {
	rxn := len(in)
	for len(out)|len(in) != 0 {
		n := len(out)
		if n != 0 {
			if n > 0xFFFF {
				n = 0xFFFF
			}
			d.p.StoreTXDPTR(unsafe.Pointer(&out[0]))
			out = out[n:]
		}
		d.p.StoreTXDMAXCNT(n)
		n = len(in)
		if n != 0 {
			if n > 0xFFFF {
				n = 0xFFFF
			}
			d.p.StoreRXDPTR(unsafe.Pointer(&in[0]))
			in = in[n:]
		}
		d.p.StoreRXDMAXCNT(n)
		d.done.Clear()
		if n < 8 {
			// avoid interrupts for small n
			d.p.Task(START).Trigger()
			for !d.p.Event(END).IsSet() {
				if n > 1 {
					n--
					runtime.Gosched()
				}
			}
		} else {
			d.p.Event(END).EnableIRQ()
			d.p.Task(START).Trigger()
			d.done.Sleep(-1)
		}
		d.p.Event(END).Clear()
	}
	return rxn
}
