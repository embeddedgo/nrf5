// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"time"

	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/spim"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/init"
)

const (
	ILI9341_SLPOUT = 0x11
	ILI9341_DISPON = 0x29
	ILI9341_RAMWR  = 0x2C
	ILI9341_MADCTL = 0x36
	ILI9341_PIXFMT = 0x3A
)

type ILI9341 struct {
	SPI *spim.Driver

	reset, cs, dc gpio.Pin
}

func (ili *ILI9341) Reset() {
	time.Sleep(time.Millisecond)
	ili.reset.Clear()
	time.Sleep(time.Millisecond)
	ili.reset.Set()
}

func (ili *ILI9341) Cmd(cmd byte) {
	ili.dc.Clear()
	ili.SPI.WriteRead([]byte{cmd}, nil)
	ili.dc.Set()
}

func (ili *ILI9341) WriteByte(b byte) {
	ili.SPI.WriteRead([]byte{b}, nil)
}

func (ili *ILI9341) Write(data []byte) {
	ili.SPI.WriteRead(data, nil)
}

var ili ILI9341

func main() {
	//clock.StoreTRACECONFIG(clock.T4MHz, clock.Serial) // enable SWO on P1.00

	p0 := gpio.P(0)
	ili.dc = p0.Pin(2)
	ili.reset = p0.Pin(29)
	ili.cs = p0.Pin(31)

	p1 := gpio.P(1)
	miso := p1.Pin(10)
	sck := p1.Pin(13)
	mosi := p1.Pin(15)

	ili.dc.Clear()
	ili.dc.Setup(gpio.ModeOut)
	ili.reset.Set()
	ili.reset.Setup(gpio.ModeOut)
	ili.cs.Set()
	ili.cs.Setup(gpio.ModeOut)

	ili.SPI = spim.NewDriver(spim.SPIM(0))
	ili.SPI.UsePin(sck, spim.SCK)
	ili.SPI.UsePin(miso, spim.MISO)
	ili.SPI.UsePin(mosi, spim.MOSI)
	ili.SPI.SetFreq(spim.F8MHz)
	ili.SPI.Enable()
	ili.SPI.IRQ().Enable(rtos.IntPrioLow, 0)

	ili.Reset()
	time.Sleep(10 * time.Millisecond)
	ili.cs.Clear()
	ili.Cmd(ILI9341_SLPOUT)
	time.Sleep(120 * time.Millisecond)
	ili.Cmd(ILI9341_DISPON)

	ili.Cmd(ILI9341_PIXFMT)
	ili.WriteByte(0x55) // 16 bit 565 format.

	ili.Cmd(ILI9341_MADCTL)
	ili.WriteByte(0x48) // Screen orientation.

	ili.Cmd(ILI9341_RAMWR)

	line := make([]byte, 240*2)
	for i := 0; i < 320; i++ {
		for k := range line {
			x := i
			if k == x || k == x+1 {
				line[k] = 0xff
			} else {
				line[k] = 0
			}
		}
		ili.Write(line)
	}

	/*
		// Reset and select
		reset.Clear()
		time.Sleep(1 * time.Millisecond)
		reset.Set()
		time.Sleep(120 * time.Millisecond)
		cs.Clear()

		cmd := func(cmds ...byte) {
			dc.Clear()
			spi.WriteRead(cmds, nil)
			dc.Set()
		}

		cmd(ili9341.SLPOUT)
		time.Sleep(120 * time.Millisecond)
		cmd(ili9341.DISPON)
		cmd(ili9341.PIXSET)
		spi.WriteRead([]byte{byte(ili9341.PF16)}, nil)
		cmd(ili9341.MADCTL)
		spi.WriteRead([]byte{byte(ili9341.MY | ili9341.MX | ili9341.MV | ili9341.BGR)}, nil)

		cmd(ili9341.CASET)
		spi.WriteRead([]byte{0, 0, 320 >> 8, 320 & 0xFF}, nil)
		cmd(ili9341.PASET)
		spi.WriteRead([]byte{0, 0, 320 >> 8, 240 & 0xFF}, nil)
		cmd(ili9341.RAMWR)

		for i := 0; i < 100; i++ {
			spi.WriteRead([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0, 0, 0, 0}, nil)
		}
	*/
}

//go:interrupthandler
func SPIM0_SPIS0_TWIM0_TWIS0_SPI0_TWI0_Handler() {
	ili.SPI.ISR()
}
