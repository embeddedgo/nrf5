// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Display draws on the connected display.
package main

import (
	"time"

	"github.com/embeddedgo/display/pix/displays"
	"github.com/embeddedgo/display/pix/examples"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/init"

	"github.com/embeddedgo/nrf5/dci/tftdci"
	"github.com/embeddedgo/nrf5/hal/clock"
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/spim"
	"github.com/embeddedgo/nrf5/hal/spim/spim3"
)

func main() {
	clock.StoreTRACECONFIG(clock.T4MHz, clock.Serial) // enable SWO on P1.00

	// Assign GPIO pins

	p0 := gpio.P(0)
	scl := p0.Pin(13)
	mosi := p0.Pin(15)
	reset := p0.Pin(17)
	dc := p0.Pin(20)
	cs := p0.Pin(22)
	miso := p0.Pin(24)

	// Configure peripherals

	reset.Set()
	reset.Setup(gpio.ModeOut)

	spi := spim3.Driver()
	spi.UsePin(scl, spim.SCK)
	spi.UsePin(miso, spim.MISO)
	spi.UsePin(mosi, spim.MOSI)
	dci := tftdci.NewSPIM(spi, dc, spim.CPOL0|spim.CPHA0, spim.F16MHz, spim.F16MHz)
	dci.UseCSN(cs, false)

	// Reset

	reset.Clear()
	time.Sleep(time.Millisecond)
	reset.Set()

	// Run

	//disp := displays.Adafruit_0i96_128x64_OLED_SSD1306(dci)
	//disp := displays.Adafruit_1i5_128x128_OLED_SSD1351(dci)
	//disp := displays.Adafruit_1i54_240x240_IPS_ST7789(dci)
	//disp := displays.Adafruit_2i8_240x320_TFT_ILI9341(dci)
	disp := displays.ERTFTM_1i54_240x240_IPS_ST7789(dci)
	//disp := displays.MSP4022_4i0_320x480_TFT_ILI9486(dci)
	//disp := displays.Waveshare_1i5_128x128_OLED_SSD1351(dci)

	for {
		examples.RotateDisplay(disp)
		examples.DrawText(disp)
		examples.GraphicsTest(disp)
	}
}
