// Copyright 2021 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"
	"time"

	"github.com/embeddedgo/nrf5/hal/clock"
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/spim"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/init"
)

const (
	ILI9341_SWRESET = 0x01
	ILI9341_SLPOUT  = 0x11
	ILI9341_DISPON  = 0x29
	ILI9341_RAMWR   = 0x2C
	ILI9341_MADCTL  = 0x36
	ILI9341_PIXSET  = 0x3A
)

type ILI9341 struct {
	SPI *spim.Driver

	reset, cs, dc gpio.Pin
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
	clock.StoreTRACECONFIG(clock.T4MHz, clock.Serial) // enable SWO on P1.00

	p0 := gpio.P(0)
	ili.dc = p0.Pin(2)
	ili.reset = p0.Pin(29)
	ili.cs = p0.Pin(31)

	p1 := gpio.P(1)
	//miso := p1.Pin(10)
	sck := p1.Pin(13)
	mosi := p1.Pin(15)

	ili.dc.Clear()
	ili.dc.Setup(gpio.ModeOut)
	ili.reset.Clear() // assert RESET
	ili.reset.Setup(gpio.ModeOut)
	ili.cs.Clear() // assert CS
	ili.cs.Setup(gpio.ModeOut)

	ili.SPI = spim.NewDriver(spim.SPIM(0))
	ili.SPI.UsePin(sck, spim.SCK)
	ili.SPI.UsePin(mosi, spim.MOSI)
	ili.SPI.SetFreq(spim.F8MHz)
	ili.SPI.Enable()
	ili.SPI.IRQ().Enable(rtos.IntPrioLow, 0)

	line := make([]byte, 240*2)
	for {
		ili.reset.Clear()
		time.Sleep(time.Millisecond)
		ili.reset.Set()
		resetTime := time.Now()

		time.Sleep(5 * time.Millisecond)

		ili.Cmd(ILI9341_PIXSET)
		ili.WriteByte(0x55) // 16 bit 565 format.

		ili.Cmd(ILI9341_MADCTL)
		ili.WriteByte(0x48) // Screen orientation.

		time.Sleep(resetTime.Add(120 * time.Millisecond).Sub(time.Now()))

		ili.Cmd(ILI9341_SLPOUT)

		time.Sleep(5 * time.Millisecond)

		ili.Cmd(ILI9341_DISPON)
		ili.Cmd(ILI9341_RAMWR)

		for k := 0; k < len(line); k += 2 {
			line[k] = 0
			line[k+1] = 0
		}
		for i := 0; i < 320; i++ {
			ili.Write(line)
		}

		time.Sleep(time.Second)

		var x int
		for i := 0; i < 320; i++ {
			for k := 0; k < len(line); k += 2 {
				x++
				line[k] = byte(x >> 8)
				line[k+1] = byte(x)
			}
			ili.Write(line)
			time.Sleep(5 * time.Millisecond)
		}

		time.Sleep(time.Second)
	}
}

//go:interrupthandler
func SPIM0_SPIS0_TWIM0_TWIS0_SPI0_TWI0_Handler() {
	ili.SPI.ISR()
}
