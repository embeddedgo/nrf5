// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package spi provides access to the registers of SPI peripheral.
// It also provides the driver that implements asynchronous and synchronous I/O.
package spi

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/internal"
	"github.com/embeddedgo/nrf5/hal/te"
	"github.com/embeddedgo/nrf5/p/mmap"
)

type Periph struct {
	te.Regs

	_         [64]mmio.U32
	enable    mmio.U32
	_         mmio.U32
	psel      [3]mmio.U32
	_         mmio.U32
	rxd       mmio.U32
	txd       mmio.U32
	_         mmio.U32
	frequency mmio.U32
	_         [11]mmio.U32
	config    mmio.U32
}

func SPI(n int) *Periph {
	switch n {
	case 0:
		return (*Periph)(unsafe.Pointer(mmap.SPI0_BASE))
	case 1:
		return (*Periph)(unsafe.Pointer(mmap.SPI1_BASE))
	case 2:
		return (*Periph)(unsafe.Pointer(mmap.SPI2_BASE))
	}
	return nil
}

type Event byte

const (
	READY Event = 2 // TXD byte sent and RXD byte received.
)

func (p *Periph) Event(e Event) *te.Event { return p.Regs.Event(int(e)) }

// LoadENABLE reports whether the SPI peripheral is enabled.
func (p *Periph) LoadENABLE() bool {
	return p.enable.Load()&1 != 0
}

// StoreENABLE enables or disables SPI peripheral.
func (p *Periph) StoreENABLE(en bool) {
	p.enable.Store(internal.Bool2uint32(en))
}

type Signal byte

const (
	SCK  Signal = 0
	MOSI Signal = 1
	MISO Signal = 2
)

// LoadPSEL returns configuration of signal s.
func (p *Periph) LoadPSEL(s Signal) (pin gpio.Pin, connected bool) {
	return p.psel[s].Pin()
}

// StorePSEL configures signal s.
func (p *Periph) StorePSEL(s Signal, pin gpio.Pin, connected bool) {
	p.psel[s].SetPin(pin, connected)
}

// LoadRX returns data received (double buffered).
func (p *Periph) LoadRXD() byte {
	return byte(p.rxd.Load())
}

// StoreTXD stores data to send (double buffered).
func (p *Periph) StoreTXD(b byte) {
	p.txd.Store(uint32(b))
}

// Freq sets
type Freq uint32

const (
	Freq125k Freq = 0x02000000 // 125 kbps
	Freq250k Freq = 0x04000000 // 250 kbps
	Freq500k Freq = 0x08000000 // 500 kbps
	Freq1M   Freq = 0x10000000 // 1 Mbps
	Freq2M   Freq = 0x20000000 // 2 Mbps
	Freq4M   Freq = 0x40000000 // 4 Mbps
	Freq8M   Freq = 0x80000000 // 8 Mbps
)

// LoadFREQUENCY returns configured SCK frequency.
func (p *Periph) LoadFREQUENCY() Freq {
	return Freq(p.frequency.Load())
}

// StoreFREQUENCY stores SCK frequency.
func (p *Periph) StoreFREQUENCY(f Freq) {
	p.frequency.Store(uint32(f))
}

// Config is a bitfield that describes SPI configuration.
type Config byte

const (
	MSBF  Config = 0      // Most significant bit shifted out first.
	LSBF  Config = 1 << 0 // Most significant bit shifted out first.
	CPHA0 Config = 0      // Sample on leading SCK edge, shift data on trailing edge.
	CPHA1 Config = 1 << 1 // Sample on trailing SCK edge, shift data on leading edge.
	CPOL0 Config = 0      // SCK polarity: active high.
	CPOL1 Config = 1 << 2 // SCK polarity: active low.
)

// LoadCONFIG returns current SPI configuration.
func (p *Periph) LoadCONFIG() Config {
	return Config(p.config.Load())
}

// StoreCONFIG stores SPI configuration..
func (p *Periph) StoreCONFIG(c Config) {
	p.config.Store(uint32(c))
}
