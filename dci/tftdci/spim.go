// Copyright 2021 The Embedded Go authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tftdci

import (
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/spim"
)

// SPIM is an implementation of the tftdrv.DCI that uses an SPIM peripheral
// to communicate with the display in what is known as 4-line serial mode.
type SPIM struct {
	spi     *spim.Driver
	dc      gpio.Pin
	csn     gpio.Pin
	mode    spim.Config
	rf, wf  spim.Freq
	started bool
	reconf  bool
}

// NewSPIM returns new SPI based implementation of tftdrv.DCI. It properly
// configures the provided SPI driver and DC pin to communicate with a display
// controller. Select the SPI mode (CPOL,CPHA), write and read clock speed
// according to the display controller specification. Note that the maximum
// speed may be limited by the concrete instance of nRF5 SPI peripheral, the
// bus topology and the specific display design.
func NewSPIM(drv *spim.Driver, dc gpio.Pin, mode spim.Config, rf, wf spim.Freq) *SPIM {
	dci := &SPIM{spi: drv, dc: dc, mode: mode, wf: wf, rf: rf}
	dc.Clear()
	dc.Setup(gpio.ModeOut)
	drv.Setup(dci.mode, wf)
	return dci
}

// UseCSN selects csn as slave select pin. If reconf is true the SPI peripheral
// is reconfigured by any Cmd call so it can be shared with other applications
// (exclusive acces is required until End call).
func (dci *SPIM) UseCSN(csn gpio.Pin, reconf bool) {
	dci.csn = csn
	dci.reconf = reconf
	csn.Set()
	csn.Setup(gpio.ModeOut)
}

func (dci *SPIM) Driver() *spim.Driver { return dci.spi }
func (dci *SPIM) Err(clear bool) error { return nil }
func (dci *SPIM) DC() gpio.Pin         { return dci.dc }

func start(dci *SPIM) {
	dci.started = true
	if dci.csn.IsValid() {
		dci.csn.Clear()
		if dci.reconf {
			dci.spi.Setup(dci.mode, dci.wf)
		}
	}
	dci.spi.Enable()
}

func (dci *SPIM) Cmd(p []byte, _ int) {
	if !dci.started {
		start(dci)
	}
	dci.dc.Clear()
	dci.spi.WriteRead(p, nil)
	dci.dc.Set()
}

func (dci *SPIM) End() {
	dci.started = false
	if dci.csn.IsValid() {
		dci.csn.Set()
	}
	dci.spi.Disable()
}

func (dci *SPIM) WriteBytes(p []uint8) {
	if !dci.started {
		start(dci)
	}
	dci.spi.WriteRead(p, nil)
}

func (dci *SPIM) ReadBytes(p []byte) {
	if !dci.started {
		start(dci)
	}
	dci.spi.Periph().StoreFREQUENCY(dci.rf)
	dci.spi.WriteRead(nil, p)
	dci.spi.Periph().StoreFREQUENCY(dci.wf)
}
