// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpio

import (
	"unsafe"
)

// Pin represents one phisical pin (specific pin in specific port).
type Pin struct {
	h uintptr // [31:7] port address, [6:5] port number, [4:0] pin number
}

// IsValid reports whether p represents a valid pin.
func (p Pin) IsValid() bool {
	return p.h != 0
}

// Port returns the port where the pin is located.
func (p Pin) Port() *Port {
	return (*Port)(unsafe.Pointer(p.h &^ 0x7F))
}

// PSEL returns the numerical representation of digital Pin.
func (p Pin) PSEL() PSEL {
	if p.h == 0 {
		return -1
	}
	return PSEL(p.h & 0x7F)
}

// AIN retruns analog intput number that corresponds to p or -1 if
// pin can not be used as analog input.
func (p Pin) AIN() AIN {
	return ain(p)
}

// Num returns the pin number in the port.
func (p Pin) Num() int {
	return int(p.h & 0x1F)
}

// Setup configures pin.
func (p Pin) Setup(cfg Config) {
	p.Port().SetupPin(p.Num(), cfg)
}

// Config returns current configuration of pin.
func (p Pin) Config() Config {
	return p.Port().PinConfig(p.Num())
}

// Mask returns bitmask that represents the pin.
func (p Pin) Mask() Pins {
	return Pin0 << uint(p.Num())
}

// Load returns input value of the pin.
func (p Pin) Load() int {
	return int(p.Port().in.Load()) >> uint(p.Num()) & 1
}

// LoadOut returns output value of the pin.
func (p Pin) LoadOut() int {
	return int(p.Port().out.Load()) >> uint(p.Num()) & 1
}

// Set sets output value of the pin to 1 in one atomic operation.
func (p Pin) Set() {
	p.Port().outset.Store(uint32(Pin0) << uint(p.Num()))
}

// Clear sets output value of the pin to 0 in one atomic operation.
func (p Pin) Clear() {
	p.Port().outclr.Store(uint32(Pin0) << uint(p.Num()))
}

// Store sets output value of the pin to the least significant bit of val.
func (p Pin) Store(val int) {
	port := p.Port()
	n := uint(p.Num())
	if val&1 != 0 {
		port.outset.Store(uint32(Pin0) << n)
	} else {
		port.outclr.Store(uint32(Pin0) << n)
	}
}

// PSEL is numerical representation of GPIO pin.
type PSEL int8

// Pin returns the GPIO pin used for digital signal that corresponds to psel.
func (ps PSEL) Pin() Pin {
	if ps < 0 {
		return Pin{}
	}
	return Pin{portaddr(int(ps)>>5) | uintptr(ps)}
}

const (
	VDD      AIN = 9   // VDD
	VDDHDIV5 AIN = 0xD // VDDH / 5
)

// AIN is an analog input number.
type AIN int8

// Pin returns the GPIO pin that can be used as analog input a.
func (a AIN) Pin() Pin {
	return apin(a)
}
