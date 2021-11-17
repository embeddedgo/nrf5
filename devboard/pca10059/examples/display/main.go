// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"
	"time"

	"github.com/embeddedgo/display/pixd"
	"github.com/embeddedgo/display/pixd/driver/tftdrv/st7789"
	"github.com/embeddedgo/nrf5/dci/tftdci"
	"github.com/embeddedgo/nrf5/hal/clock"
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/spim"
	"github.com/embeddedgo/nrf5/hal/spim/spim3"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/init"
)

func main() {
	clock.StoreTRACECONFIG(clock.T4MHz, clock.Serial) // enable SWO on P1.00

	// Assign GPIO pins
	/*
		p0 := gpio.P(0)
		dc := p0.Pin(2)
		reset := p0.Pin(29)
		csn := p0.Pin(31)

		p1 := gpio.P(1)
		miso := p1.Pin(10)
		sck := p1.Pin(13)
		mosi := p1.Pin(15)
	*/
	p0 := gpio.P(0)
	reset := p0.Pin(2)
	sda := p0.Pin(29)
	scl := p0.Pin(31)

	p1 := gpio.P(1)
	blk := p1.Pin(10)
	csn := p1.Pin(13)
	dc := p1.Pin(15)

	// Configure peripherals

	blk.Clear()
	blk.Setup(gpio.ModeOut)

	reset.Clear()
	reset.Setup(gpio.ModeOut)
	time.Sleep(time.Millisecond)
	reset.Set() // deassert RESET
	time.Sleep(5 * time.Millisecond)

	spidrv := spim3.Driver()
	spidrv.UsePin(scl, spim.SCK)
	//spidrv.UsePin(miso, spim.MISO)
	spidrv.UsePin(sda, spim.MOSI)
	dci := tftdci.NewSPIM(spidrv, dc, spim.CPOL0|spim.CPHA0, spim.F16MHz, spim.F16MHz)
	dci.UseCSN(csn, false)

	// Run

	drv := st7789.New(dci)
	drv.Init(st7789.GFX)
	//drv := ili9486.NewOver(dci)
	//drv.Init(ili9486.MSP4022)
	disp := pixd.NewDisplay(drv)
	disp.SetRect(image.Rect(0, 0+40, 240, 240+40))
	r := disp.Bounds()

	a := disp.NewArea(r)
	a.SetColorRGBA(77, 78, 79, 255)
	a.Fill(a.Bounds())
	blk.Set()

	for i := 0; ; i++ {
		time.Sleep(2 * time.Second)
		println(i & 3)
		disp.SetDir(i)
		a.SetRect(a.Rect())
		r := a.Bounds()

		a.SetColorRGBA(0, 0, 0, 255)
		a.Fill(a.Bounds())
		a.SetColorRGBA(255, 0, 0, 255)
		a.Fill(image.Rect(10, 10, 50, 200))
		a.SetColorRGBA(0, 255, 0, 255)
		a.Fill(image.Rect(60, 10, 100, 200))
		a.SetColorRGBA(0, 0, 255, 255)
		a.Fill(image.Rect(110, 10, 150, 200))
		a.SetColorRGBA(128, 0, 128, 128)
		a.RoundRect(r.Min, r.Max.Sub(image.Pt(1, 1)), 0, 0, false)

	}

}
