// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uart

import (
	"embedded/rtos"
	"runtime"
	"sync/atomic"
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
// about the connection quallity or the reading performance. You can not
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
	lastrw   uint32
	rxready  rtos.Note
	overflow bool

	// tx state
	txdata string
	txn    int
	txdone rtos.Note

	timeoutRx int64
	timeoutTx int64
}

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

// EnableRx enables the UART receiver using the provided slice to buffer
// received data. At least 2-byte buffer is required, which is effectively one
// byte buffer because the other byte always remains unused for efficient
// checking of an empty state. You can not rely on 6-byte hardware buffer as
// extension of software buffer because for performance reasons the ISR do not
// return until it reads all bytes from hardware. If the software buffer is full
// the ISR simply drops read bytes until there is no more data to read. The
// driver uses at most 65536 bytes of rxbuf. EnableRx panics if the receiving is
// already enabled or rxbuf is too short.
func (d *Driver) EnableRx(rxbuf []byte) {
	if d.rxbuf != nil {
		panic("enabled before")
	}
	if len(rxbuf) < 2 {
		panic("rxbuf too short")
	}
	if len(rxbuf) > 65536 {
		rxbuf = rxbuf[:65536]
	}
	d.rxbuf = rxbuf
	d.lastrw = 0
	p := d.p
	p.Event(RXDRDY).EnableIRQ()
	p.Task(STARTRX).Trigger()
}

// DisableRx disables the UART receiver. The receive buffer is returned and no
// longer used by driver allowing GC to collect its memory. You can use the
// STOPRX and STARTRX tasks directly if you want to temporary disable the
// receiver leaving the driver intact.
func (d *Driver) DisableRx() (rxbuf []byte) {
	p := d.p
	rxto := p.Event(RXTO)
	rxto.Clear()
	p.Task(STOPRX).Trigger()
	p.Event(RXDRDY).DisableIRQ()
	for !rxto.IsSet() {
		runtime.Gosched()
	}
	rxbuf = d.rxbuf
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
	p := d.p

tryAgain:
	rxwakeup := false
	txwakeup := false
	pollmore := false

	// send next byte if the previous one was sent (little work so do it first)
	if p.Event(TXDRDY).IsSet() {
		p.Event(TXDRDY).Clear()
		if n := d.txn + 1; n < len(d.txdata) {
			d.txn = n
			p.StoreTXD(d.txdata[n])
			pollmore = true
		} else {
			// txdone.Wakeup takes a lot of time so defer it after receive
			txwakeup = true
		}
	}

	// Empty the hardware buffer even if there is no space in rxbuf. It
	// simplifies and speeds up the receiving code and makes it possible to
	// distinguish between the hardware buffer overrun (interrupt latency too
	// high) and the software buffer overlow (the reading goroutine too slow).
	if rxrdy := p.Event(RXDRDY).IsSet(); rxrdy {
		lastrw := atomic.LoadUint32(&d.lastrw)
		lastr, lastw := lastrw>>16, lastrw&0xFFFF
		rxwakeup = (lastr == lastw) // reader can sleep on empty rxbuf
		for rxrdy {
			p.Event(RXDRDY).Clear()
			b := p.LoadRXD()
			nextw := lastw + 1
			if int(nextw) == len(d.rxbuf) {
				nextw = 0
			}
			if nextw != lastr {
				d.rxbuf[nextw] = b
				lastw = nextw
			} else {
				d.overflow = true
			}
			rxrdy = p.Event(RXDRDY).IsSet()
		}
		for !atomic.CompareAndSwapUint32(&d.lastrw, lastrw, lastr<<16|lastw) {
			// in the meantime the rxbuf was read (multicore system)
			lastrw = atomic.LoadUint32(&d.lastrw)
			lastr = lastrw >> 16
			rxwakeup = (lastr == lastrw&0xFFFF)
		}
		pollmore = true
	}

	if txwakeup {
		d.txdone.Wakeup()
	}
	if rxwakeup {
		d.rxready.Wakeup()
	}
	if pollmore {
		// Avoid expensive Go ISR exit/enter code in case of high baudrates.
		// In case of nRF52+mbr+softdevice(off) useful for baudrates >= 921600.
		goto tryAgain
	}
}

// WriteByte sends one byte to the remote party and returns an error if detected// WriteByte can block if the hardware flow control is used. It does not provide
// any guarantee that the byte sent was received by the remote party.
func (d *Driver) WriteByte(b byte) (err error) {
	d.txdone.Clear()
	p := d.p
	p.Event(TXDRDY).EnableIRQ()
	p.Task(STARTTX).Trigger()
	p.StoreTXD(b)
	if !d.txdone.Sleep(d.timeoutTx) {
		err = ErrTimeout
	}
	p.Task(STOPTX).Trigger()
	p.Event(TXDRDY).DisableIRQ()
	p.Event(TXDRDY).Clear()
	return
}

// WriteString works like Write.
func (d *Driver) WriteString(s string) (n int, err error) {
	if len(s) == 0 {
		return 0, nil
	}
	d.txdata = s
	d.txn = 0
	d.txdone.Clear()
	p := d.p
	p.Event(TXDRDY).EnableIRQ()
	p.Task(STARTTX).Trigger()
	p.StoreTXD(s[0])
	if !d.txdone.Sleep(d.timeoutTx) {
		err = ErrTimeout
	}
	p.Task(STOPTX).Trigger()
	p.Event(TXDRDY).DisableIRQ()
	p.Event(TXDRDY).Clear()
	return d.txn, err
}

// Write sends bytes from p to the remote party. It return the number of bytes
// sent and error if detected. It does not provide any guarantee that the bytes
// sent were received by the remote party.
func (d *Driver) Write(p []byte) (int, error) {
	return d.WriteString(*(*string)(unsafe.Pointer(&p)))
}

// Len returns the number of bytes that are ready to read from Rx buffer.
func (d *Driver) Len() int {
	lastrw := atomic.LoadUint32(&d.lastrw)
	n := int(lastrw&0xFFFF) - int(lastrw>>16)
	if n < 0 {
		n += len(d.rxbuf)
	}
	return n
}

func (d *Driver) waitRxData() (lastr, lastw int) {
	lastrw := atomic.LoadUint32(&d.lastrw)
	lastr, lastw = int(lastrw>>16), int(lastrw&0xFFFF)
	if lastr != lastw || !d.rxready.Sleep(d.timeoutRx) {
		return
	}
	lastrw = atomic.LoadUint32(&d.lastrw)
	lastr, lastw = int(lastrw>>16), int(lastrw&0xFFFF)
	if lastr != lastw {
		return
	}
	panic("wakeup on empty buffer")
}

func (d *Driver) markDataRead(lastr int) error {
	ulastr := uint32(lastr)
	for {
		lastrw := atomic.LoadUint32(&d.lastrw)
		lastw := lastrw & 0xFFFF
		if lastw == ulastr {
			d.rxready.Clear()
		}
		if atomic.CompareAndSwapUint32(&d.lastrw, lastrw, ulastr<<16|lastw) {
			break
		}
	}
	if d.overflow {
		d.overflow = false
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
	lastr, lastw := d.waitRxData()
	if lastr == lastw {
		return 0, ErrTimeout
	}
	if lastr++; lastr == len(d.rxbuf) {
		lastr = 0
	}
	return d.rxbuf[lastr], d.markDataRead(lastr)
}

// Read reads up to len(p) bytes into p. It returns number of bytes read and an
// error if detected. Read blocks only if the internal buffer is empty (d.Len()
// > 0 ensure non-blocking read).
func (d *Driver) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	lastr, lastw := d.waitRxData()
	if lastr == lastw {
		return 0, ErrTimeout
	}
	nextr := lastr + 1
	if nextr == len(d.rxbuf) {
		nextr = 0
	}
	if nextr <= lastw {
		n = copy(p, d.rxbuf[nextr:lastw+1])
	} else {
		n = copy(p, d.rxbuf[nextr:])
		if n < len(p) {
			n += copy(p[n:], d.rxbuf[:lastw+1])
		}
	}
	lastr += n
	if lastr >= len(d.rxbuf) {
		lastr -= len(d.rxbuf)
	}
	return n, d.markDataRead(lastr)
}

// SetReadTimeout sets the read timeout used by Read* functions.
func (d *Driver) SetReadTimeout(ns int64) {
	d.timeoutRx = ns
}

// SetWriteTimeout sets the write timeout used by Write* functions.
func (d *Driver) SetWriteTimeout(ns int64) {
	d.timeoutTx = ns
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
		return "uart: buffer overflow"
	case ErrTimeout:
		return "uart: timeout"
	}
	return ""
}
