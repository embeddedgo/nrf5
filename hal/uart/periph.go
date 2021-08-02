// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package uart provides access to the registers of UART peripheral.
// It also provides the driver that implements io.ReadWriter interface.
package uart

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

	_        [32]uint32
	errorsrc mmio.U32
	_        [31]uint32
	enable   mmio.U32
	_        uint32
	psel     [4]psel.Reg
	rxd      mmio.U32
	txd      mmio.U32
	_        uint32
	baudrate mmio.U32
	_        [17]uint32
	config   mmio.U32
}

func UART(n int) *Periph {
	if n != 0 {
		return nil
	}
	return (*Periph)(unsafe.Pointer(mmap.UART0_BASE))
}

type Task uint8

const (
	STARTRX Task = 0 // Start UART receiver.
	STOPRX  Task = 1 // Stop UART receiver.
	STARTTX Task = 2 // Start UART transmitter.
	STOPTX  Task = 3 // Stop UART transmitter.
	SUSPEND Task = 7 // Suspend UART. nRF52.
)

type Event uint8

const (
	CTS    Event = 0  // CTS is activated (set low). nRF51v3+.
	NCTS   Event = 1  // CTS is deactivated (set high). nRF51v3+.
	RXDRDY Event = 2  // Data received in RXD.
	TXDRDY Event = 7  // Data sent from TXD.
	ERROR  Event = 9  // Error detected.
	RXTO   Event = 17 // Receiver timeout.
)

func (p *Periph) Task(t Task) *te.Task    { return p.Regs.Task(int(t)) }
func (p *Periph) Event(e Event) *te.Event { return p.Regs.Event(int(e)) }

type Shorts uint32

const (
	CTS_STARTRX Shorts = 1 << 3
	NCTS_STOPRX Shorts = 1 << 4
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
	s := "uart:"
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
	return p.enable.Load() == 4
}

// StoreENABLE enables or disables UART peripheral.
func (p *Periph) StoreENABLE(en bool) {
	p.enable.Store(4 * internal.BoolToUint32(en))
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

// LoadRXD returns data received in previous transfers.
func (p *Periph) LoadRXD() byte {
	return byte(p.rxd.Load())
}

// StoreTX stores data to be transferred.
func (p *Periph) StoreTXD(b byte) {
	p.txd.Store(uint32(b))
}

type Baudrate uint32

const (
	Baud1200   Baudrate = 0x0004F000 // Actual rate: 1205 baud.
	Baud2400   Baudrate = 0x0009D000 // Actual rate: 2396 baud.
	Baud4800   Baudrate = 0x0013B000 // Actual rate: 4808 baud.
	Baud9600   Baudrate = 0x00275000 // Actual rate: 9598 baud.
	Baud14400  Baudrate = 0x003B0000 // Actual rate: 14414 baud.
	Baud19200  Baudrate = 0x004EA000 // Actual rate: 19208 baud.
	Baud28800  Baudrate = 0x0075F000 // Actual rate: 28829 baud.
	Baud31250  Baudrate = 0x00800000
	Baud38400  Baudrate = 0x009D5000 // Actual rate: 38462 baud.
	Baud56000  Baudrate = 0x00E50000 // Actual rate: 55944 baud.
	Baud57600  Baudrate = 0x00EBF000 // Actual rate: 57762 baud.
	Baud76800  Baudrate = 0x013A9000 // Actual rate: 76923 baud.
	Baud115200 Baudrate = 0x01D7E000 // Actual rate: 115942 baud.
	Baud230400 Baudrate = 0x03AFB000 // Actual rate: 231884 baud.
	Baud250000 Baudrate = 0x04000000
	Baud460800 Baudrate = 0x075F7000 // Actual rate: 470588 baud.
	Baud921600 Baudrate = 0x0EBED000 // Actual rate: 941176 baud.
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
