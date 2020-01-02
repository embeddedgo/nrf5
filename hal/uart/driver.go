// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uart

import (
	"embedded/rtos"
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/gpio"
)

// Driver is an interrupt based driver for UART peripheral. It is optimized for
// error free links (it is fast and efficient but has some limitations when it
// comes to reporting Rx errors).
//
// Reading from UART can cause hardware errors (there is no hardware error
// detection on writing). This driver treats all Rx hardware errors as
// asynchronous events (at least the ErrOverrun is in fact synchronous) so they
// are simply imformative about the connection quallity and reading performance
// but you can not determine which data has been affected by error (use other
// techniques to ensure data integrity).
//
// Set the read timeout to ensure wake-up in case of missing data. The hardware
// may not detect some Rx errors or the error can be signaled before you try to
// read affected data because of the internal hardware and software buffering.
// This means that the reader can not rely on waking up in case of Rx error.
// Consider also that the remote party can reset unexpectedly and depending on
// the protocol it can wait for a data request or an initialization sequence
// before sending anything.
//
// The write operation can block only if the hardware flow control is enabled.
// In this case you can use write timeout to detect problems with the remote
// party or RTS/CTS signaling.
//
// The driver does not support concurent reading or writing by multiple
// gorutines. There can be one reading goroutine and one writing goroutine at
// the same time. Both can work concurrently with the driver interrupt handler
// (parallel execution on multi-core systems is supported).
type Driver struct {
	P *Periph

	// rx state
	rxbuf   []byte
	lastrw  uint32
	rxready rtos.Note

	// tx state
	txdata string
	txn    int
	txdone rtos.Note

	timeoutRx int64
	timeoutTx int64
}

// NewDriver returns a new driver for p. The rxbuf can be nil in case of Tx-only
// use case. Otherwise at least 2-byte buffer is required because one byte
// remains unused for efficient checking of an empty state. Reading from 2-byte
// buffer always returns ErrBufOverflow because for performance reasons full
// rxbuf is treated as overflowed. At least 8-byte buffer is recommended to
// ensure the full 6-byte hardware buffer can be copied to the empty rxbuf and
// there is still one byte left to avoid ErrBufOverflow error. Please note that
// 1 Mbaud means 100 bytes in millisecond. In case of 2 ms scheduling period and
// another busy thread at least 202-byte buffer is need to have a chance to
// avoid overflows. The driver can use up to 65536 bytes for its Rx buffer.
func NewDriver(p *Periph, rxbuf []byte) *Driver {
	if len(rxbuf) > 65536 {
		rxbuf = rxbuf[:65536]
	}
	return &Driver{timeoutRx: -1, timeoutTx: -1, P: p, rxbuf: rxbuf}
}

func (d *Driver) Enable() {
	d.P.StoreENABLE(true)
}

func (d *Driver) Disable() {
	d.P.StoreENABLE(false)
}

// EnableRx enables UART receiver.
func (d *Driver) EnableRx() {
	if len(d.rxbuf) < 2 {
		panic("rxbuf too short")
	}
	p := d.P
	p.Event(RXDRDY).EnableIRQ()
	p.Task(STARTRX).Trigger()
}

// DisableRx disables UART receiver.
func (d *Driver) DisableRx() {
	d.P.Task(STOPRX).Trigger()
}

// EnableTx enables UART transmitter.
func (d *Driver) EnableTx() {
	p := d.P
	p.Event(TXDRDY).EnableIRQ()
	p.Task(STARTTX).Trigger()
}

// DisableTx disables UART transmitter.
func (d *Driver) DisableTx() {
	d.P.Task(STOPTX).Trigger()
}

func (d *Driver) SetReadTimeout(ns int64) {
	d.timeoutRx = ns
}

func (d *Driver) SetWriteTimeout(ns int64) {
	d.timeoutTx = ns
}

func (d *Driver) UsePin(s Signal, pin gpio.Pin) {
	d.P.StorePSEL(s, pin.PSEL())
}

func (d *Driver) SetBaudrate(br Baudrate) {
	d.P.StoreBAUDRATE(br)
}

func (d *Driver) SetConfig(cfg Config) {
	d.P.StoreCONFIG(cfg)
}

func (d *Driver) IRQ() rtos.IRQ {
	return d.P.IRQ()
}

// ISR handles UART interrupts. It supports the reading thread to run in
// parallel on another CPU.
func (d *Driver) ISR() {
	p := d.P

tryAgain:
	rxwakeup := false
	txwakeup := false
	active := false

	// send next byte if the previous one was sent (little work so do it first)
	if p.Event(TXDRDY).IsSet() {
		p.Event(TXDRDY).Clear()
		if n := d.txn + 1; n < len(d.txdata) {
			d.txn = n
			p.StoreTXD(d.txdata[n])
			active = true
		} else {
			// txdone.Wakeup takes a lot of time so defer it after receive
			txwakeup = true
		}
	}

	// Empty the hardware buffer even if there is no space in rxbuf. It
	// simplifies and speeds up the receiving code and makes it possible to
	// distinguish between the hardware buffer overrun error (interrupt latency
	// is too high) and the software buffer overlow error (the reading goroutine
	// is too slow).
	if rxrdy := p.Event(RXDRDY).IsSet(); rxrdy {
		active = true
	checkRxbuf:
		lastrw := atomic.LoadUint32(&d.lastrw)
		lastr, lastw := lastrw>>16, lastrw&0xFFFF
		rxwakeup = (lastr == lastw) // gorutine can sleep on empty rxbuf
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
			}
			rxrdy = p.Event(RXDRDY).IsSet()
		}
		if !atomic.CompareAndSwapUint32(&d.lastrw, lastrw, lastr<<16|lastw) {
			goto checkRxbuf // in the meantime the rxbuf was read
		}
	}

	if txwakeup {
		d.txdone.Wakeup()
	}
	if rxwakeup {
		d.rxready.Wakeup()
	}
	if active {
		// avoid expensive Go ISR exit/enter code in case of high baudrates
		goto tryAgain
	}
}

// Len returns number of bytes that are ready to read from internal Rx buffer.
func (d *Driver) Len() int {
	lastrw := atomic.LoadUint32(&d.lastrw)
	n := int(lastrw&0xFFFF) - int(lastrw>>16)
	if n < 0 {
		n += len(d.rxbuf)
	}
	return n
}

// In case of timeout the transmission is stopped but you can not be sure the
// byte was sent or not. Use EnableTx before subsequent write.
func (d *Driver) WriteByte(b byte) error {
	d.txdone.Clear()
	d.P.StoreTXD(b)
	if d.txdone.Sleep(d.timeoutTx) {
		return nil
	}
	d.P.Task(STOPTX).Trigger()
	d.P.Event(TXDRDY).Clear()
	return ErrTimeout
}

// In case of timeout the transmission is stopped and you can not rely on the
// returned number of bytes sent. Use EnableTx before subsequent write.
func (d *Driver) WriteString(s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}
	d.txdata = s
	d.txn = 0
	d.txdone.Clear()
	d.P.StoreTXD(s[0])
	if d.txdone.Sleep(d.timeoutTx) {
		return len(s), nil
	}
	d.P.Task(STOPTX).Trigger()
	d.P.Event(TXDRDY).Clear()
	return d.txn, ErrTimeout
}

func (d *Driver) Write(p []byte) (int, error) {
	return d.WriteString(*(*string)(unsafe.Pointer(&p)))
}

func (d *Driver) waitRxData() (lastr, lastw uint32) {
	lastrw := atomic.LoadUint32(&d.lastrw)
	lastr, lastw = lastrw>>16, lastrw&0xFFFF
	if lastr != lastw || !d.rxready.Sleep(d.timeoutRx) {
		return
	}
	lastrw = atomic.LoadUint32(&d.lastrw)
	lastr, lastw = lastrw>>16, lastrw&0xFFFF
	if lastr != lastw {
		return
	}
	panic("wakeup on empty buffer")
}

func (d *Driver) markDataRead(lastr uint32) error {
	var lastrw uint32
	for {
		lastrw = atomic.LoadUint32(&d.lastrw)
		lastw := lastrw & 0xFFFF
		if lastw == lastr {
			d.rxready.Clear()
		}
		if atomic.CompareAndSwapUint32(&d.lastrw, lastrw, lastr<<16|lastw) {
			break
		}
	}
	n := int(lastrw>>16) - int(lastrw&0xFFFF)
	if n < 0 {
		n += len(d.rxbuf)
	}
	if n == 1 {
		return ErrBufOverflow
	}
	if e := d.P.LoadERRORSRC(); e != 0 {
		d.P.ClearERRORSRC(e)
		return e
	}
	return nil
}

// ReadByte reads one byte from the internal buffer.
func (d *Driver) ReadByte() (b byte, err error) {
	lastr, lastw := d.waitRxData()
	if lastr == lastw {
		return 0, ErrTimeout
	}
	if lastr++; int(lastr) == len(d.rxbuf) {
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
	if int(nextr) == len(d.rxbuf) {
		nextr = 0
	}
	if nextr <= lastw {
		n = copy(p, d.rxbuf[nextr:lastw+1])
	} else {
		n = copy(p, d.rxbuf[nextr:])
		n += copy(p[n:], d.rxbuf[:lastw+1])
	}
	return n, d.markDataRead(lastw)
}

type DriverError uint8

const (
	ErrBufOverflow DriverError = iota + 1
	ErrTimeout
)

func (e DriverError) Error() string {
	switch e {
	case ErrBufOverflow:
		return "uart: buffer overflow"
	case ErrTimeout:
		return "uart: timeout"
	}
	return ""
}
