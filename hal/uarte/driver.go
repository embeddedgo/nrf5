// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uarte

import (
	"embedded/rtos"
	"time"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/gpio"
)

type Driver struct {
	p *Periph

	// tx state
	txdone rtos.Note

	timeoutRx time.Duration
	timeoutTx time.Duration
}

// NewDriver returns a new driver for p.
func NewDriver(p *Periph) *Driver {
	p.EnableIRQ(1 << ENDTX)
	return &Driver{p: p, timeoutRx: -1, timeoutTx: -1}
}

func (d *Driver) Periph() *Periph {
	return d.p
}

// Enable enables UART peripheral.
func (d *Driver) Enable() {
	d.p.StoreENABLE(true)
}

// Disable disables UART peripheral.
func (d *Driver) Disable() {
	d.p.StoreENABLE(false)
}

// UsePin configurs the specified pin to be used as signal s.
func (d *Driver) UsePin(pin gpio.Pin, s Signal) {
	switch s {
	case TXD, RTSn:
		pin.Set()
		pin.Setup(gpio.ModeOut)
	default:
		pin.Setup(gpio.ModeIn)
	}
	d.p.StorePSEL(s, pin.PSEL(), true)
}

// SetBaudrate sets Tx and Rx baudrate.
func (d *Driver) SetBaudrate(br Baudrate) {
	d.p.StoreBAUDRATE(br)
}

// SetConfig allows to configure UART peripheral to use hardware flow controll,
// add and check parity bit, and use two stop bits instead default one.
func (d *Driver) SetConfig(cfg Config) {
	d.p.StoreCONFIG(cfg)
}

// IRQ returns the interrupt assigned to UART peripheral used by driver.
func (d *Driver) IRQ() rtos.IRQ {
	return d.p.IRQ()
}

func (d *Driver) ISR() {
	p := d.p
	if endtx := p.Event(ENDTX); endtx.IsSet() {
		endtx.Clear()
		d.txdone.Wakeup()
	}
}

func (d *Driver) WriteByte(b byte) (err error) {
	d.txdone.Clear()
	p := d.p
	p.StoreTXDPTR(unsafe.Pointer(&b))
	p.StoreTXDMAXCNT(1)
	p.Task(STARTTX).Trigger()
	if !d.txdone.Sleep(d.timeoutTx) {
		err = ErrTimeout
	}
	p.Task(STOPTX).Trigger()
	return
}

func (d *Driver) Write(s []byte) (n int, err error) {
	if len(s) == 0 {
		return 0, nil
	}
	d.txdone.Clear()
	p := d.p
	p.StoreTXDPTR(unsafe.Pointer(&s[0]))
	p.StoreTXDMAXCNT(len(s))
	p.Task(STARTTX).Trigger()
	if !d.txdone.Sleep(d.timeoutTx) {
		err = ErrTimeout
	}
	p.Task(STOPTX).Trigger()
	return p.LoadTXDAMOUNT(), err
}

type DriverError uint8

const (
	// ErrBufOverflow is returned if one or more received bytes has been dropped
	// because of the lack of free space in the driver's receive buffer.
	ErrBufOverflow DriverError = iota + 1

	// ErrTimeout is returned if timeout occured. It means that the read/write
	// operation has been interrupted. In case of write you can not determine
	// the exact number of bytes sent to the remote party.
	ErrTimeout
)

// Error implements error interface.
func (e DriverError) Error() string {
	switch e {
	case ErrBufOverflow:
		return "uarte: buffer overflow"
	case ErrTimeout:
		return "uarte: timeout"
	}
	return ""
}
