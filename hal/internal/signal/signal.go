// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package signal

import (
	"embedded/mmio"

	"github.com/embeddedgo/nrf5/hal/gpio"
)

// Digital represents a peripheral register that corresponds to one digital I/O
// signal. Its Set* methods can be used only if the peripheral is disabled.
type Digital struct {
	u32 mmio.U32
}

// SetPin sets GPIO pin that will be used for digital signal d. In fact the
// signal is connected to the selected GPIO pin if both the connect is true and
// the peripheral is enabled.
func (d *Digital) SetPin(pin gpio.Pin, connect bool) {
	psel := uint32(int32(pin.PSEL()))
	if !connect {
		psel |= 1 << 31
	}
	d.u32.Store(psel)
}

// Pin returns the current configuration of digital signal d.
func (d *Digital) Pin() (pin gpio.Pin, connected bool) {
	psel := d.u32.Load()
	return gpio.PSEL(psel).Pin(), psel>>31 == 0
}

//// SetConnect allows to configure the signal d to be connected to or
//// disconnected from the selected GPIO pin.
//func (d *Digital) SetConnect(connect bool) {
//	if connect {
//		d.u32.ClearBits(1 << 31)
//	} else {
//		d.u32.SetBits(1 << 31)
//	}
//}

// Analog represents a peripheral register that corresponds to one analog I/O
// signal. Its Set* methods can be used only if the peripheral is disabled.
type Analog struct {
	u32 mmio.U32
}

// SetAIN sets analog input number used for analog signal a.
func (a *Analog) SetAIN(ain gpio.AIN) {
	a.u32.Store(uint32(ain))
}

// AIN returns analog input number used for analog signal a.
func (a *Analog) AIN() gpio.AIN {
	return gpio.AIN(a.u32.Load())
}
