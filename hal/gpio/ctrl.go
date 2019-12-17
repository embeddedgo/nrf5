// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpio

// Pins returns input value of pins.
func (p *Port) Pins(pins Pins) Pins {
	return Pins(p.in.LoadBits(uint32(pins)))
}

// PinsOut returns output value of pins.
func (p *Port) PinsOut(pins Pins) Pins {
	return Pins(p.out.LoadBits(uint32(pins)))
}

// SetPins sets output value of pins to 1 in one atomic operation.
func (p *Port) SetPins(pins Pins) {
	p.outset.Store(uint32(pins))
}

// ClearPins sets output value of pins to 0 in one atomic operation.
func (p *Port) ClearPins(pins Pins) {
	p.outclr.Store(uint32(pins))
}

// Load returns input value of all pins.
func (p *Port) Load() Pins {
	return Pins(p.in.Load())
}

// LoadOut returns output value of all pins.
func (p *Port) LoadOut() Pins {
	return Pins(p.out.Load())
}

// Store sets output value of all pins to value specified by val.
func (p *Port) Store(val Pins) {
	p.out.Store(uint32(val))
}