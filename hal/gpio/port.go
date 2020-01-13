// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpio

import (
	"embedded/mmio"
	"unsafe"
)

// Port represents a GPIO port.
type Port struct {
	_          mmio.U32
	out        mmio.U32
	outset     mmio.U32
	outclr     mmio.U32
	in         mmio.U32
	dir        mmio.U32
	dirset     mmio.U32
	dirclr     mmio.U32
	latch      mmio.U32 // TODO
	detectmode mmio.U32 // TODO
	_          [118]mmio.U32
	pincnf     [32]mmio.U32
}

// P returns n-th GPIO port
func P(n int) *Port {
	return (*Port)(unsafe.Pointer(portaddr(n)))
}

// Num returns the port number.
func (p *Port) Num() int {
	return portnum(p)
}

// Pins is a bitmask which represents the pins of GPIO port.
type Pins uint32

const (
	Pin0 Pins = 1 << iota
	Pin1
	Pin2
	Pin3
	Pin4
	Pin5
	Pin6
	Pin7
	Pin8
	Pin9
	Pin10
	Pin11
	Pin12
	Pin13
	Pin14
	Pin15
	Pin16
	Pin17
	Pin18
	Pin19
	Pin20
	Pin21
	Pin22
	Pin23
	Pin24
	Pin25
	Pin26
	Pin27
	Pin28
	Pin29
	Pin30
	Pin31
)

// Pin returns n-th pin in port p.
func (p *Port) Pin(n int) Pin {
	if uint(n) > 31 {
		panic("gpio: bad pin")
	}
	ptr := uintptr(unsafe.Pointer(p))
	return Pin{ptr | uintptr(p.Num())<<5 | uintptr(n)}
}
