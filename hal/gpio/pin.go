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
	return p.h&^0x7F != 0
}

// Port returns the port where the pin is located.
func (p Pin) Port() *Port {
	return (*Port)(unsafe.Pointer(p.h &^ 0x7F))
}

// PSEL returns the PSEL representation of GPIO pin in connected state.
func (p Pin) PSEL() PSEL {
	if p.h&^0x7F == 0 {
		return 0xFFFFFFFF
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

// Detect reports whether the pin have met the criteria set by Sense*
// configuration options. nRF52.
func (p Pin) Detect() bool {
	return p.Port().latch.Load()>>uint(p.Num())&1 != 0
}

// ClearDetect clears the detect state for pin. nRF52.
func (p Pin) ClearDetect() {
	p.Port().latch.Store(1 << uint(p.Num()))
}

// PSEL is numerical representation of GPIO pin used as peripheral digital
// signal. It can have two states: connected or disconnected to the peripheral.
type PSEL uint32

// IsConnected reports the connection state of ps.
func (ps PSEL) IsConnected() bool {
	return ps>>31 != 0
}

// Connected returns ps with connection state changed.
func (ps PSEL) Connected(connected bool) PSEL {
	if connected {
		return ps &^ 1 << 31
	}
	return ps | 1<<31
}

// Pin returns the GPIO pin corresponding to ps.
func (ps PSEL) Pin() Pin {
	if ps < 0 {
		return Pin{}
	}
	return Pin{portaddr(int(ps)>>5) | uintptr(ps)}
}

const (
	ANC     AIN = 0   // Not connected.
	AIN0    AIN = 1   // Analog input 0.
	AIN1    AIN = 2   // Analog input 1.
	AIN2    AIN = 3   // Analog input 2.
	AIN3    AIN = 4   // Analog input 3.
	AIN4    AIN = 5   // Analog input 4.
	AIN5    AIN = 6   // Analog input 5.
	AIN6    AIN = 7   // Analog input 6.
	AIN7    AIN = 8   // Analog input 7.
	VDD     AIN = 9   // VDD
	VDDH1_5 AIN = 0xD // VDDH / 5
)

// AIN is an analog input number.
type AIN int8

// Pin returns the GPIO pin that can be used as analog input a.
func (a AIN) Pin() Pin {
	return apin(a)
}
