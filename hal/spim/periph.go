// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spim

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/internal"
	"github.com/embeddedgo/nrf5/hal/internal/psel"
	"github.com/embeddedgo/nrf5/hal/te"
	"github.com/embeddedgo/nrf5/p/mmap"
)

type Periph struct {
	te.Regs

	stallstat       mmio.U32
	_               [63]uint32
	enable          mmio.U32
	_               uint32
	psel            [4]psel.Reg
	_               [3]uint32
	frequency       mmio.U32
	_               [3]uint32
	rxdptr          mmio.U32
	rxdmaxcnt       mmio.U32
	rxdamount       mmio.U32
	rxdlist         mmio.U32
	txdptr          mmio.U32
	txdmaxcnt       mmio.U32
	txdamount       mmio.U32
	txdlist         mmio.U32
	config          mmio.U32
	_               [2]uint32
	iftimingrxdelay mmio.U32
	iftimingcsndur  mmio.U32
	csnpol          mmio.U32
	pseldcx         psel.Reg
	dcxcnt          mmio.U32
	_               [19]uint32
	orc             mmio.U32
}

func SPIM(n int) *Periph {
	switch n {
	case 0:
		return (*Periph)(unsafe.Pointer(mmap.SPIM0_BASE))
	case 1:
		return (*Periph)(unsafe.Pointer(mmap.SPIM1_BASE))
	case 2:
		return (*Periph)(unsafe.Pointer(mmap.SPIM2_BASE))
	case 3:
		return (*Periph)(unsafe.Pointer(mmap.SPIM3_BASE))
	}
	return nil
}

type Task uint8

const (
	START   Task = 4 // Start SPI transaction.
	STOP    Task = 5 // Stop SPI transaction.
	SUSPEND Task = 7 // Suspend SPI transaction.
	RESUME  Task = 8 // Resume SPI transaction.
)

type Event uint8

const (
	STOPPED Event = 1  // SPI transaction has stopped.
	ENDRX   Event = 4  // End of RXD buffer reached.
	END     Event = 6  // End of RXD buffer and TXD buffer reached.
	ENDTX   Event = 8  // End of TXD buffer reached.
	STARTED Event = 19 // Transaction started.
)

func (p *Periph) Task(t Task) *te.Task    { return p.Regs.Task(int(t)) }
func (p *Periph) Event(e Event) *te.Event { return p.Regs.Event(int(e)) }

type Shorts uint32

const (
	END_START Shorts = 1 << 17
)

func (p *Periph) LoadSHORTS() Shorts   { return Shorts(p.Regs.LoadSHORTS()) }
func (p *Periph) StoreSHORTS(s Shorts) { p.Regs.StoreSHORTS(uint32(s)) }

// LoadSTALSTAT returns stall status for EasyDMA acesses. SPIM3 only.
func (p *Periph) LoadSTALSTAT() (rx, tx bool) {
	stallstat := p.stallstat.Load()
	return stallstat&2 != 0, stallstat&1 != 0
}

// ClearSTALSTAT clears stall status. SPIM3 only.
func (p *Periph) ClearSTALSTAT() {
	p.stallstat.Store(0)
}

// LoadENABLE reports whether the UART peripheral is enabled.
func (p *Periph) LoadENABLE() bool {
	return p.enable.Load() == 7
}

// StoreENABLE enables or disables UART peripheral.
func (p *Periph) StoreENABLE(en bool) {
	p.enable.Store(7 * internal.BoolToUint32(en))
}

type Signal uint8

const (
	SCK  Signal = 0
	MOSI Signal = 1
	MISO Signal = 2
	CSN  Signal = 3 // SPIM3 only.
	DCX  Signal = 4 // SPIM3 only.
)

// LoadPSEL returns configuration of signal s.
func (p *Periph) LoadPSEL(s Signal) (psel gpio.PSEL, en bool) {
	if s == DCX {
		return p.pseldcx.Load()
	}
	return p.psel[s].Load()
}

// StorePSEL configures signal s.
func (p *Periph) StorePSEL(s Signal, psel gpio.PSEL, en bool) {
	if s == DCX {
		p.pseldcx.Store(psel, en)
	} else {
		p.psel[s].Store(psel, en)
	}
}

type Freq uint32

const (
	F125kHz Freq = 0x02000000
	F250kHz Freq = 0x04000000
	F500kHz Freq = 0x08000000
	F1MHz   Freq = 0x10000000
	F2MHz   Freq = 0x20000000
	F4MHz   Freq = 0x40000000
	F8MHz   Freq = 0x80000000
	F16MHz  Freq = 0x0A000000 // SPIM3 only.
	F32MHz  Freq = 0x14000000 // SPIM3 only.
)

// LoadFREQUENCY returns SPI frequency (SCK clock).
func (p *Periph) LoadFREQUENCY() Freq {
	return Freq(p.frequency.Load())
}

// StoreFREQUENCY sets SPI frequency (SCK clock).
func (p *Periph) StoreFREQUENCY(f Freq) {
	p.frequency.Store(uint32(f))
}

// LoadRXDPTR returns the Rx data pointer.
func (p *Periph) LoadRXDPTR() uintptr {
	return uintptr(p.rxdptr.Load())
}

// StoreRXDPTR sets the Rx data pointer.
func (p *Periph) StoreRXDPTR(ptr unsafe.Pointer) {
	p.rxdptr.Store(uint32(uintptr(ptr)))
}

// LoadRXDMAXCNT returns the maximum number of bytes in receive buffer.
func (p *Periph) LoadRXDMAXCNT() int {
	return int(p.rxdmaxcnt.Load())
}

// StoreRXDMAXCNT sets the maximum number of bytes in receive buffer.
func (p *Periph) StoreRXDMAXCNT(n int) {
	p.rxdmaxcnt.Store(uint32(n))
}

// LoadRXDAMOUNT returns the Number of bytes transferred in the last Rx
// transaction
func (p *Periph) LoadRXDAMOUNT() int {
	return int(p.rxdamount.Load())
}

// LoadRXDLIST reports EasyDMA array list mode for received data.
func (p *Periph) LoadRXDLIST() bool {
	return p.rxdlist.Load() != 0
}

// StoreRXDLIST enables/disables EasyDMA array list mode for received data.
func (p *Periph) StoreRXDLIST(arrayListMode bool) {
	p.rxdlist.Store(internal.BoolToUint32(arrayListMode))
}

// StoreTXDPTR sets the Tx data pointer.
func (p *Periph) StoreTXDPTR(ptr unsafe.Pointer) {
	p.txdptr.Store(uint32(uintptr(ptr)))
}

// LoadTXDMAXCNT returns the maximum number of bytes in transmit buffer.
func (p *Periph) LoadTXDMAXCNT() int {
	return int(p.txdmaxcnt.Load())
}

// StoreTXDMAXCNT sets the maximum number of bytes in transmit buffer.
func (p *Periph) StoreTXDMAXCNT(n int) {
	p.txdmaxcnt.Store(uint32(n))
}

// LoadTXDAMOUNT returns the number of bytes transferred in the last Tx
// transaction
func (p *Periph) LoadTXDAMOUNT() int {
	return int(p.txdamount.Load())
}

// LoadTXDLIST reports EasyDMA array list mode for transmitted data.
func (p *Periph) LoadTXDLIST() bool {
	return p.txdlist.Load() != 0
}

// StoreTXDLIST enables/disables EasyDMA array list mode for transmitted data.
func (p *Periph) StoreTXDLIST(arrayListMode bool) {
	p.txdlist.Store(internal.BoolToUint32(arrayListMode))
}

type Config uint8

const (
	MSBF  Config = 0 << 0 // Most significant bit first.
	LSBF  Config = 1 << 0 // Least significant bit first.
	CPHA0 Config = 0 << 1 // Sample on leading edge.
	CPHA1 Config = 1 << 1 // Sample on trailing edge.
	CPOL0 Config = 0 << 2 // Clock idle state is 0.
	CPOL1 Config = 1 << 2 // Clock idle state is 1.
)

// LoadCONFIG returns the SPI protocol configuration.
func (p *Periph) LoadCONFIG() Config {
	return Config(p.config.Load())
}

// StoreCONFIG sets the SPI protocol configuration.
func (p *Periph) StoreCONFIG(cfg Config) {
	p.config.Store(uint32(cfg))
}

// LoadRXDELAY returns the number of 64 MHz clock cycles delay from sampling
// edge of SCK until the input serial data is sampled. SPIM3 only.
func (p *Periph) LoadRXDELAY() int {
	return int(p.iftimingrxdelay.Load())
}

// StoreRXDELAY sets the number of 64 MHz clock cycles [0..7] delay from
// sampling edge of SCK until the input serial data is sampled. SPIM3 only.
func (p *Periph) StoreRXDELAY(rxdelay int) {
	p.iftimingrxdelay.Store(uint32(rxdelay))
}

// LoadCSNDUR returns the number of 64 MHz clock cycles of minimum duration
// between edge of CSN and edge of SCK and minimum duration CSN must stay high
// between transactions. SPIM3 only.
func (p *Periph) LoadCSNDUR() int {
	return int(p.iftimingcsndur.Load())
}

// StoreCSNDUR sets the number of 64 MHz clock cycles [0..255] of minimum
// duration between edge of CSN and edge of SCK and minimum duration CSn must
// stay high between transactions. SPIM3 only.
func (p *Periph) StoreCSNDUR(csndur int) {
	p.iftimingcsndur.Store(uint32(csndur))
}

// LoadCSNPOL returns CSN polarity (0: active low, 1 active high). SPIM3 only.
func (p *Periph) LoadCSNPOL() int {
	return int(p.csnpol.Load())
}

// StoreCSNPOL sets CSN polarity (0: active low, 1 active high). SPIM3 only.
func (p *Periph) StoreCSNPOL(csnpol int) {
	p.csnpol.Store(uint32(csnpol))
}

// LoadDCXCNT returns the number of command bytes preceding the data bytes.
// SPIM3 only.
func (p *Periph) LoadDCXCNT() int {
	return int(p.dcxcnt.Load())
}

// StoreDCXCNT sets the number of command bytes preceding the data bytes. SPIM3
// only.
func (p *Periph) StoreDCXCNT(dcxcnt int) {
	p.dcxcnt.Store(uint32(dcxcnt))
}

// LoadORC returns the value transmitted after TXD.MAXCNT bytes have been
// transmitted in the case when RXD.MAXCNT is greater than TXD.MAXCNT.
func (p *Periph) LoadORC() byte {
	return byte(p.orc.Load())
}

// StoreORC sets the value transmitted after TXD.MAXCNT bytes have been
// transmitted in the case when RXD.MAXCNT is greater than TXD.MAXCNT.
func (p *Periph) StoreORC(orc byte) {
	p.orc.Store(uint32(orc))
}
