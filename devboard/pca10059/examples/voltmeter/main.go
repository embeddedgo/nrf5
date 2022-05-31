// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This code shows how to use SAADC to implement a simple voltmeter. It can
// mesure voltages in single ended and differential mode with 8 ranges.
//
// Do not exceed the permissible voltage range (VSS <= Vin <= VDD) on any input
// because it may damage the microcontroller.
//
// The user interface consists of small monochrome OLED display and the onboard
// button.
//
// An additional function is a clock and simple oscilloscope.
package main

import (
	"embedded/rtos"
	"time"

	"github.com/embeddedgo/display/pix/displays"

	"github.com/embeddedgo/nrf5/devboard/pca10059/board/buttons"

	"github.com/embeddedgo/nrf5/dci/tftdci"
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/gpiote"
	"github.com/embeddedgo/nrf5/hal/saadc"
	"github.com/embeddedgo/nrf5/hal/spim"
	"github.com/embeddedgo/nrf5/hal/spim/spim3"
)

var (
	button *Button
	ticker = time.NewTicker(time.Second / 8)
)

func main() {
	// External peripherals with their signals (original names)

	var oled struct{ scl, sda, res, dc gpio.Pin }

	// Assigning GPIO pins to the signals of external peripherals

	p0 := gpio.P(0)
	oled.scl = p0.Pin(13)
	oled.sda = p0.Pin(15)
	oled.res = p0.Pin(17)
	oled.dc = p0.Pin(20)

	// Configure internal peripherals

	oled.res.Set()
	oled.res.Setup(gpio.ModeOut)

	spi := spim3.Driver()
	spi.UsePin(oled.scl, spim.SCK)
	spi.UsePin(oled.sda, spim.MOSI)
	dci := tftdci.NewSPIM(spi, oled.dc, spim.CPOL0|spim.CPHA0, spim.F16MHz, spim.F16MHz)

	// Initialize button

	button = NewButton(buttons.User.Pin(), buttons.User.Pin().Load())
	gpiote.IRQ().Enable(rtos.IntPrioLow, 0)

	// Initialize OLED

	oled.res.Clear()
	time.Sleep(time.Millisecond)
	oled.res.Set()
	disp.Init(displays.Adafruit_0i96_128x64_OLED_SSD1306(dci))

	// Run

	//time.Set(time.Now(), time.Date(2022, 6, 2, 0, 55, 20, 0, time.Local))

	for {
		clock()
		voltmeter()
	}

}

func clock() {
start:
	disp.H.Clear()
	disp.H.WriteString("Clock")
	disp.P.Align = Acenter
	for {
		disp.P.Clear()
		t := time.Now()
		disp.P.WriteString(t.Format("Monday\n15:04:05\n2 Jan 2006"))
		disp.Flush()

		select {
		case long := <-button.C:
			if !long {
				return
			}
			setClock()
			goto start
		case <-ticker.C:
		}
	}
}

func itos(v int, s []byte) {
	for i := len(s) - 1; i >= 0; i-- {
		s[i], v = byte(v%10+'0'), v/10
	}
}

func stoi(s []byte) int {
	v := 0
	for _, b := range s {
		v = v*10 + int(b) - '0'
	}
	return v
}

var ddmax = [14]byte{
	'9', '9', '9', '9', '1', '9', '3', '9',
	'2', '9', '5', '9', '5', '9',
}

func setClock() {
	var dd [14]byte
	t := time.Now()
	ye, mo, da := t.Date()
	itos(ye*1e4+int(mo)*1e2+da, dd[0:8])
	h, m, s := t.Clock()
	itos(h*1e4+m*1e2+s, dd[8:14])
	p := 0
	d := dd[0]

	disp.H.Clear()
	disp.H.WriteString("Set clock")
	disp.P.Align = Acenter
	for {
		disp.P.Clear()
		disp.P.Printf(
			"\n%s-%s-%s\n%s:%s:%s\n",
			dd[0:4], dd[4:6], dd[6:8], dd[8:10], dd[10:12], dd[12:14],
		)
		disp.Flush()

		select {
		case long := <-button.C:
			if long {
				dd[p] = d
				if p++; p == len(dd) {
					t := time.Now().Add(-1050 * time.Millisecond)
					d := time.Date(
						stoi(dd[0:4]), time.Month(stoi(dd[4:6])), stoi(dd[6:8]),
						stoi(dd[8:10]), stoi(dd[10:12]), stoi(dd[12:14]), 0,
						time.Local,
					)
					time.Set(t, d)
					return
				}
				d = dd[p]
			} else {
				if d++; d > ddmax[p] {
					d = '0'
				}
				dd[p] = d
			}
		case <-ticker.C:
			if dd[p] == ' ' {
				dd[p] = d
			} else {
				dd[p] = ' '
			}
		}
	}
}

func voltmeter() {
	const res = 10
	saadc.AnalogInit(res)

	disp.H.Clear()
	disp.H.WriteString("Voltmeter")
	disp.P.Align = Acenter
	for {
		disp.P.Clear()
		v := float32(saadc.AnalogRead(gpio.AIN7)) * 3.0 / (1 << res)
		disp.P.Printf("\n% 4.2f V", v)
		disp.Flush()

		select {
		case long := <-button.C:
			if !long {
				return
			}
		case <-ticker.C:
		}
	}
}

//go:interrupthandler
func GPIOTE_Handler() {
	if button.Event.IsSet() {
		button.ISR()
	}
}
