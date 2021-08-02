// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package uarte provides access to the registers of UARTE peripheral.
// It also provides the driver that implements io.ReadWriter interface.
package uarte

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

	_         [32]uint32
	errorsrc  mmio.U32
	_         [31]uint32
	enable    mmio.U32
	_         uint32
	psel      [4]psel.Reg
	_         [3]uint32
	baudrate  mmio.U32
	_         [3]uint32
	rxdptr    mmio.U32
	rxdmaxcnt mmio.U32
	rxdamount mmio.U32
	_         uint32
	txdptr    mmio.U32
	txdmaxcnt mmio.U32
	txdamount mmio.U32
	_         [7]uint32
	config    mmio.U32
}

func UARTE(n int) *Periph {
	switch n {
	case 0:
		return (*Periph)(unsafe.Pointer(mmap.UARTE0_BASE))
	case 1:
		return (*Periph)(unsafe.Pointer(mmap.UARTE1_BASE))
	}
	return nil
}

type Task uint8

const (
	STARTRX Task = 0  // Start UART receiver.
	STOPRX  Task = 1  // Stop UART receiver.
	STARTTX Task = 2  // Start UART transmitter.
	STOPTX  Task = 3  // Stop UART transmitter.
	FLUSHRX Task = 11 // Flush RX FIFO into RX buffer.
)

type Event uint8

const (
	CTS       Event = 0  // CTS is activated (set low).
	NCTS      Event = 1  // CTS is deactivated (set high).
	RXDRDY    Event = 2  // Data received in RXD.
	ENDRX     Event = 4  // Receive buffer is filled up.
	TXDRDY    Event = 7  // Data sent from TXD.
	ENDTX     Event = 8  // Last TX byte transmitted.
	ERROR     Event = 9  // Error detected.
	RXTO      Event = 17 // Receiver timeout.
	RXSTARTED Event = 19 // UART receiver has started.
	TXSTARTED Event = 20 // UART transmitter has started.
	TXSTOPPED Event = 22 // Transmitter stopped.
)

func (p *Periph) Task(t Task) *te.Task    { return p.Regs.Task(int(t)) }
func (p *Periph) Event(e Event) *te.Event { return p.Regs.Event(int(e)) }

type Shorts uint32

const (
	ENDRX_STARTRX Shorts = 1 << 5
	ENDRX_STOPRX  Shorts = 1 << 6
)

func (p *Periph) LoadSHORTS() Shorts   { return Shorts(p.Regs.LoadSHORTS()) }
func (p *Periph) StoreSHORTS(s Shorts) { p.Regs.StoreSHORTS(uint32(s)) }

// ErrorBits is a bitfield that lists detected errors.
type ErrorBits uint8

const (
	EOVERRUN ErrorBits = 1 << 0
	EPARITY  ErrorBits = 1 << 1
	EFRAMING ErrorBits = 1 << 2
	EBREAK   ErrorBits = 1 << 3

	EALL = EOVERRUN | EPARITY | EFRAMING | EBREAK
)

func (e ErrorBits) Error() string {
	if e == 0 {
		return ""
	}
	s := "uarte:"
	if e&EOVERRUN != 0 {
		s += " overrun"
	}
	if e&EPARITY != 0 {
		s += " parity"
	}
	if e&EFRAMING != 0 {
		s += " framing"
	}
	if e&EBREAK != 0 {
		s += " break"
	}
	return s
}

// LoadERRORSRC returns error flags.
func (p *Periph) LoadERRORSRC() ErrorBits {
	return ErrorBits(p.errorsrc.Load())
}

// ClearERRORSRC clears specfied error flags.
func (p *Periph) ClearERRORSRC(e ErrorBits) {
	p.errorsrc.Store(uint32(e))
}

// LoadENABLE reports whether the UART peripheral is enabled.
func (p *Periph) LoadENABLE() bool {
	return p.enable.Load() == 8
}

// StoreENABLE enables or disables UART peripheral.
func (p *Periph) StoreENABLE(en bool) {
	p.enable.Store(8 * internal.BoolToUint32(en))
}

type Signal uint8

const (
	RTSn Signal = 0
	TXD  Signal = 1
	CTSn Signal = 2
	RXD  Signal = 3
)

// LoadPSEL returns configuration of signal s.
func (p *Periph) LoadPSEL(s Signal) (psel gpio.PSEL, en bool) {
	return p.psel[s].Load()
}

// StorePSEL configures signal s.
func (p *Periph) StorePSEL(s Signal, psel gpio.PSEL, en bool) {
	p.psel[s].Store(psel, en)
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

// LoadRXDAMOUNT returns the Number of bytes transferred in the last Rx transaction
func (p *Periph) LoadRXDAMOUNT() int {
	return int(p.rxdamount.Load())
}

// LoadTXDPTR returns the Tx data pointer.
func (p *Periph) LoadTXDPTR() uintptr {
	return uintptr(p.txdptr.Load())
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

// LoadTXDAMOUNT returns the number of bytes transferred in the last Tx transaction
func (p *Periph) LoadTXDAMOUNT() int {
	return int(p.txdamount.Load())
}

type Baudrate uint32

const (
	Baud1200   Baudrate = 0x0004F000 // Actual rate: 1205 baud.
	Baud2400   Baudrate = 0x0009D000 // Actual rate: 2396 baud.
	Baud4800   Baudrate = 0x0013B000 // Actual rate: 4808 baud.
	Baud9600   Baudrate = 0x00275000 // Actual rate: 9598 baud.
	Baud14400  Baudrate = 0x003AF000 // Actual rate: 14401 baud.
	Baud19200  Baudrate = 0x004EA000 // Actual rate: 19208 baud.
	Baud28800  Baudrate = 0x0075C000 // Actual rate: 28777 baud.
	Baud31250  Baudrate = 0x00800000
	Baud38400  Baudrate = 0x009D0000 // Actual rate: 38369 baud.
	Baud56000  Baudrate = 0x00E50000 // Actual rate: 55944 baud.
	Baud57600  Baudrate = 0x00EB0000 // Actual rate: 57554 baud.
	Baud76800  Baudrate = 0x013A9000 // Actual rate: 76923 baud.
	Baud115200 Baudrate = 0x01D60000 // Actual rate: 115108 baud.
	Baud230400 Baudrate = 0x03B00000 // Actual rate: 231884 baud.
	Baud250000 Baudrate = 0x04000000
	Baud460800 Baudrate = 0x07400000 // Actual rate: 457143 baud.
	Baud921600 Baudrate = 0x0F000000 // Actual rate: 941176 baud.
	Baud1M     Baudrate = 0x10000000
)

func Baud(baud int) Baudrate {
	return (Baudrate(uint64(baud)<<32/16e6) + 0x800) & 0xFFFFF000
}

func (br Baudrate) Baud() int {
	return int((uint64(br)*16e6 + 1<<31) >> 32)
}

// LoadBAUDRATE returns configured baudrate.
func (p *Periph) LoadBAUDRATE() Baudrate {
	return Baudrate(p.baudrate.Load())
}

// StoreBAUDRATE stores baudrate.
func (p *Periph) StoreBAUDRATE(br Baudrate) {
	p.baudrate.Store(uint32(br))
}

type Config uint32

const (
	HWFC   Config = 0x01 << 0 // Hardware flow control
	PARITY Config = 0x07 << 1 // Parity
	STOP2  Config = 0x01 << 4 // Two stop bits
)

// LoadCONFIG returns configuration. nRF52.
func (p *Periph) LoadCONFIG() Config {
	return Config(p.config.Load())
}

// StoreCONFIG stores baudrate. nRF52.
func (p *Periph) StoreCONFIG(cfg Config) {
	p.config.Store(uint32(cfg))
}
