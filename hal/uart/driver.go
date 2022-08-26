// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uart

import (
	"embedded/rtos"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/gpio"
)

// Driver is an interrupt based driver for UART peripheral. It is optimized for
// error free links (it is fast and efficient but has some limitations when it
// comes to reporting Rx errors).
//
// Reading from UART you may encounter errors detected by the hardware (there is
// no hardware error detection on writing). This driver treats all Rx hardware
// errors as asynchronous events but at least the ErrOverrun is in fact
// synchronous. So the Rx errors other than ErrTimeout are simply imformative
// about the connection quallity or the reading performance. You cannot
// determine which data has been affected (use other techniques to ensure data
// integrity.
//
// Set the read timeout to ensure wake-up in case of missing data. The hardware
// may not detect some Rx errors or the error can be signaled before you try to
// read affected data because of the hardware and software buffering. This means
// that the reader can not rely on waking up in case of Rx error. Consider also
// that the remote party can reset unexpectedly and depending on the protocol
// used it can be quiet waiting for data request or initialization sequence.
//
// The write operation can block only if the hardware flow control is enabled.
// In this case you can use write timeout to detect problems with the remote
// party or RTS/CTS signaling.
//
// The driver supports one reading goroutine and one writing goroutine that both
// can work concurrently with the driver.
type Driver struct {
	p *Periph

	// rx state
	rxbuf    []byte
	nextr    uint32
	nextw    uint32
	rxcmd    uint32
	overflow uint32
	rxready  rtos.Note

	// tx state
	txdata string
	txn    int
	txdone rtos.Note

	timeoutRx time.Duration
	timeoutTx time.Duration
}

const (
	cmdNone = iota
	cmdWakeup
	cmdStop
)

// NewDriver returns a new driver for p.
func NewDriver(p *Periph) *Driver {
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

// EnableRx enables the UART receiver.
//
// It allocates an internal ring buffer of bufLen size. In most cases bufLen=64/
// is good choise. Minimum buffer size is 2 which means effectively one byte
// buffer because the other byte always remains unused for efficient checking
// of an empty state.
//
// You cannot rely on 6-byte hardware buffer as extension of the software buffer
// because for the performance reasons the ISR will not return until it has
// read all bytes from hardware. If the software buffer is full the ISR simply
// drops read bytes until there is no more data to read.
//
// EnableRx panics if the receiving is already enabled or bufLen is too small.
func (d *Driver) EnableRx(bufLen int) {
	if d.rxbuf != nil {
		panic("enabled before")
	}
	if bufLen < 2 {
		panic("rxbuf too short")
	}
	d.rxbuf = make([]byte, bufLen)
	d.nextr = 0
	d.nextw = 0
	d.p.Event(RXDRDY).EnableIRQ()
	d.p.Task(STARTRX).Trigger()
}

// DisableRx disables the UART receiver. The receive buffer is freed.
// You can use the STOPRX and STARTRX tasks directly if you want to temporary
// disable the receiver leaving the driver intact.
func (d *Driver) DisableRx() {
	atomic.StoreUint32(&d.rxcmd, cmdStop)
	d.p.Event(RXTO).Clear()
	d.p.Event(RXTO).EnableIRQ()
	d.p.Task(STOPRX).Trigger()
	d.rxready.Sleep(-1)
	d.p.Event(RXDRDY).DisableIRQ()
	d.p.Event(RXTO).DisableIRQ()
	d.rxbuf = nil
	return
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

// ISR handles UART interrupts. It supports the reading thread to run in
// parallel on another CPU.
func (d *Driver) ISR() {
	for d.p.Event(RXDRDY).IsSet() {
		d.p.Event(RXDRDY).Clear()
		nextw := d.nextw

		// Read from hardware buffer even if there is no space in d.rxbuf. It
		// simplifies the receiving code and makes it possible to distinguish
		// between the EOVERRUN (interrupt handler too slow or interrup latency
		// to high) and the ErrBufOverflow (reading goroutine too slow).
		d.rxbuf[nextw] = d.p.LoadRXD()

		if nextw++; int(nextw) == len(d.rxbuf) {
			nextw = 0
		}
		if nextw != atomic.LoadUint32(&d.nextr) {
			atomic.StoreUint32(&d.nextw, nextw)
			if atomic.CompareAndSwapUint32(&d.rxcmd, cmdWakeup, cmdNone) {
				d.rxready.Wakeup()
			}
		} else {
			atomic.StoreUint32(&d.overflow, 1)
		}
	}

	if d.p.Event(TXDRDY).IsSet() {
		d.p.Event(TXDRDY).Clear()
		if n := d.txn; n < len(d.txdata) {
			d.p.StoreTXD(d.txdata[n])
			d.txn = n + 1
		} else {
			d.txdone.Wakeup()
		}
	}

	if d.p.Event(RXTO).IsSet() {
		d.p.Event(RXTO).Clear()
		if atomic.CompareAndSwapUint32(&d.rxcmd, cmdStop, cmdNone) {
			d.rxready.Wakeup()
		}
	}
}

// WriteByte sends one byte to the remote party and returns an error if detected// WriteByte can block if the hardware flow control is used. It does not provide
// any guarantee that the byte sent was received by the remote party.
func (d *Driver) WriteByte(b byte) (err error) {
	d.txn = 1
	d.txdone.Clear()
	d.p.Task(STARTTX).Trigger()
	d.p.StoreTXD(b)
	d.p.Event(TXDRDY).EnableIRQ()
	if !d.txdone.Sleep(d.timeoutTx) {
		err = ErrTimeout
	}
	d.p.Task(STOPTX).Trigger()
	d.p.Event(TXDRDY).DisableIRQ()
	d.p.Event(TXDRDY).Clear()
	return // BUG: in case of timeout the ISR can still run in multicore system
}

// WriteString works like Write.
func (d *Driver) WriteString(s string) (n int, err error) {
	if len(s) == 0 {
		return 0, nil
	}
	d.txdata = s
	d.txn = 1
	d.txdone.Clear()
	d.p.Task(STARTTX).Trigger()
	d.p.StoreTXD(s[0])
	d.p.Event(TXDRDY).EnableIRQ()
	if !d.txdone.Sleep(d.timeoutTx) {
		err = ErrTimeout
	}
	d.p.Task(STOPTX).Trigger()
	d.p.Event(TXDRDY).DisableIRQ()
	d.p.Event(TXDRDY).Clear()
	d.txdata = ""
	return d.txn, err // BUG: in case of timeout the ISR can still run in multicore system

}

// Write sends bytes from p to the remote party. It return the number of bytes
// sent and error if detected. It does not provide any guarantee that the bytes
// sent were received by the remote party.
func (d *Driver) Write(p []byte) (int, error) {
	return d.WriteString(*(*string)(unsafe.Pointer(&p)))
}

// Len returns the number of bytes that are ready to read from Rx buffer.
func (d *Driver) Len() int {
	n := int(atomic.LoadUint32(&d.nextw)) - int(d.nextr)
	if n < 0 {
		n += len(d.rxbuf)
	}
	return n
}

func (d *Driver) waitRxData() int {
	nextw := atomic.LoadUint32(&d.nextw)
	if nextw != d.nextr {
		return int(nextw)
	}
	d.rxready.Clear()
	atomic.StoreUint32(&d.rxcmd, cmdWakeup)
	nextw = atomic.LoadUint32(&d.nextw)
	if nextw != d.nextr {
		if atomic.SwapUint32(&d.rxcmd, cmdNone) == cmdNone {
			d.rxready.Sleep(-1) // wait for the upcoming wake up
		}
		return int(nextw)
	}
	if !d.rxready.Sleep(d.timeoutRx) {
		if atomic.SwapUint32(&d.rxcmd, cmdNone) != cmdNone {
			return int(nextw)
		}
		d.rxready.Sleep(-1) // wait for the upcoming wake up
	}
	nextw = atomic.LoadUint32(&d.nextw)
	if nextw != d.nextr {
		return int(nextw)
	}
	panic("wakeup on empty buffer")
}

func (d *Driver) markDataRead(nextr int) error {
	if nextr >= len(d.rxbuf) {
		nextr -= len(d.rxbuf)
	}
	atomic.StoreUint32(&d.nextr, uint32(nextr))
	if atomic.CompareAndSwapUint32(&d.overflow, 1, 0) {
		return ErrBufOverflow
	}
	if e := d.p.LoadERRORSRC(); e != 0 {
		d.p.ClearERRORSRC(e)
		return e
	}
	return nil
}

// ReadByte reads one byte and returns error if detected. ReadByte blocks only
// if the internal buffer is empty (d.Len() > 0 ensure non-blocking read).
func (d *Driver) ReadByte() (b byte, err error) {
	nextw := d.waitRxData()
	nextr := int(d.nextr)
	if nextw == nextr {
		return 0, ErrTimeout
	}
	return d.rxbuf[nextr], d.markDataRead(nextr + 1)
}

// Read reads up to len(p) bytes into p. It returns number of bytes read and an
// error if detected. Read blocks only if the internal buffer is empty (d.Len()
// > 0 ensure non-blocking read).
func (d *Driver) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	nextw := d.waitRxData()
	nextr := int(d.nextr)
	if nextw == nextr {
		return 0, ErrTimeout
	}
	if nextr <= nextw {
		n = copy(p, d.rxbuf[nextr:nextw])
	} else {
		n = copy(p, d.rxbuf[nextr:])
		if n < len(p) {
			n += copy(p[n:], d.rxbuf[:nextw])
		}
	}
	return n, d.markDataRead(nextr + n)
}

// SetReadTimeout sets the read timeout used by Read* functions.
func (d *Driver) SetReadTimeout(timeout time.Duration) {
	d.timeoutRx = timeout
}

// SetWriteTimeout sets the write timeout used by Write* functions.
func (d *Driver) SetWriteTimeout(timeout time.Duration) {
	d.timeoutTx = timeout
}

type DriverError uint8

const (
	// ErrBufOverflow is returned if one or more received bytes has been dropped
	// because of the lack of free space in the driver's receive buffer.
	ErrBufOverflow DriverError = iota + 1

	// ErrTimeout is returned if a timeout occured. It means that the read/write
	// operation has been interrupted. In case of write you can not determine
	// the exact number of bytes sent to the remote party.
	ErrTimeout
)

// Error implements error interface.
func (e DriverError) Error() string {
	switch e {
	case ErrBufOverflow:
		return "uart: buffer overflow"
	case ErrTimeout:
		return "uart: timeout"
	}
	return ""
}
