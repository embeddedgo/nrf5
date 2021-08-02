// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spim

import (
	"embedded/mmio"
	"embedded/rtos"
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
	p := d.p
	p.StoreENABLE(true)
	p.Event(END).Clear()
	p.Event(END).EnableIRQ()
}

// Disable disables SPIM peripheral.
func (d *Driver) Disable() {
	p := d.p
	p.Event(END).DisableIRQ()
	p.StoreENABLE(false)
	if uintptr(unsafe.Pointer(p)) == mmap.SPIM3_BASE {
		// see errata 195: SPIM3 continues to draw current after disable
		(*mmio.U32)(unsafe.Pointer(uintptr(0x4002F004))).Store(1)
	}
}

// UsePin configurs the specified pin to be used as signal s.
func (d *Driver) UsePin(pin gpio.Pin, s Signal) {
	p := d.p
	if s == MISO {
		pin.Setup(gpio.ModeIn)
	} else {
		switch s {
		case SCK:
			if p.LoadCONFIG()&CPOL1 != 0 {
				pin.Set()
			} else {
				pin.Clear()
			}
		case MOSI:
			pin.Clear()
		case CSN:
			pin.Store(^p.LoadCSNPOL())
		default: // DCX
			pin.Set()
		}
		pin.Setup(gpio.ModeOut)
	}
	p.StorePSEL(s, pin.PSEL(), true)
}

// SetFreq sets SPI frequency (SCK clock).
func (d *Driver) SetFreq(f Freq) {
	d.p.StoreFREQUENCY(f)
}

// SetConfig sets the SPI protocol configuration.
func (d *Driver) SetConfig(cfg Config) {
	// TODO: set gpio output for SCK
	d.p.StoreCONFIG(cfg)
}

// IRQ returns the interrupt assigned to UART peripheral used by driver.
func (d *Driver) IRQ() rtos.IRQ {
	return d.p.IRQ()
}

func (d *Driver) ISR() {
	d.p.Event(END).Clear()
	d.done.Wakeup()
}

func (d *Driver) WriteRead(out, in []byte) int {
	rxn := len(in)
	p := d.p
	for len(out)|len(in) != 0 {
		n := len(out)
		if n != 0 {
			if n > 0xFFFF {
				n = 0xFFFF
			}
			p.StoreTXDPTR(unsafe.Pointer(&out[0]))
			out = out[n:]
		}
		p.StoreTXDMAXCNT(n)
		n = len(in)
		if n != 0 {
			if n > 0xFFFF {
				n = 0xFFFF
			}
			p.StoreRXDPTR(unsafe.Pointer(&in[0]))
			in = in[n:]
		}
		p.StoreRXDMAXCNT(n)
		d.done.Clear()
		p.Task(START).Trigger()
		d.done.Sleep(-1)
	}
	return rxn
}
