// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"strings"
	"time"

	"github.com/embeddedgo/display/images"
	"github.com/embeddedgo/display/math2d"
	"github.com/embeddedgo/display/pix"
	"github.com/embeddedgo/display/pix/displays"
	"github.com/embeddedgo/display/testdata"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/init"

	"github.com/embeddedgo/nrf5/dci/tftdci"
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/spim"
	"github.com/embeddedgo/nrf5/hal/spim/spim3"
)

const frac = 6

var (
	black = color.Gray{}
	white = color.Gray{255}
	red   = color.RGBA{255, 0, 0, 255}
)

type Hand struct {
	Width int
	Base  [4]image.Point
	cur   [4]image.Point
	next  [4]image.Point
}

func (h *Hand) bounds() image.Rectangle {
	r := image.Rectangle{h.cur[0], h.cur[0]}
	for i := 1; i < 4; i++ {
		p := h.cur[i]
		if p.X < r.Min.X {
			r.Min.X = p.X
		} else if p.X > r.Max.X {
			r.Max.X = p.X
		}
		if p.Y < r.Min.Y {
			r.Min.Y = p.Y
		} else if p.Y > r.Max.Y {
			r.Max.Y = p.Y
		}
	}
	r.Max.X++
	r.Max.Y++
	return r
}

func (h *Hand) Next(alpha int32) bool {
	for i := range h.Base {
		v := math2d.F(h.Base[i], frac)
		v = math2d.Rotate(v, alpha)
		h.next[i] = math2d.I(v, frac)
	}
	return h.next != h.cur
}

func (h *Hand) Draw(a *pix.Area) {
	if h.cur[0] == h.cur[1] && h.cur[2] == h.cur[3] {
		a.SetColor(red)
		a.Line(h.cur[0], h.cur[2])
	} else {
		a.SetColor(black)
		a.Quad(h.cur[0], h.cur[1], h.cur[2], h.cur[3], true)
		a.SetColor(white)
		a.Quad(h.cur[0], h.cur[1], h.cur[2], h.cur[3], false)
	}
}

func (h *Hand) Clear(a *pix.Area, img image.Image, offset image.Point) {
	r := h.bounds()
	sp := r.Min.Add(offset)
	a.Draw(r, img, sp, nil, image.Point{}, draw.Src)
	h.cur = h.next
}

func main() {
	// External peripherals with their signals (original names)

	var ips struct{ scl, sda, res, dc, cs, blk gpio.Pin }

	// Assigning GPIO pins to the signals of external peripherals

	p0 := gpio.P(0)
	ips.scl = p0.Pin(13)
	ips.sda = p0.Pin(15)
	ips.res = p0.Pin(17)
	ips.dc = p0.Pin(20)
	ips.cs = p0.Pin(22)
	ips.blk = p0.Pin(24)

	// Configure internal peripherals

	ips.res.Set()
	ips.res.Setup(gpio.ModeOut)
	ips.blk.Clear()
	ips.blk.Setup(gpio.ModeOut)

	spi := spim3.Driver()
	spi.UsePin(ips.scl, spim.SCK)
	spi.UsePin(ips.sda, spim.MOSI)
	dci := tftdci.NewSPIM(spi, ips.dc, spim.CPOL0|spim.CPHA0, spim.F16MHz, spim.F16MHz)
	dci.UseCSN(ips.cs, false)

	// Reset

	ips.res.Clear()
	time.Sleep(time.Millisecond)
	ips.res.Set()
	ips.blk.Set()

	// Initialize display and create the drawing area with origin in the center

	disp := displays.ERTFTM_1i54_240x240_IPS_ST7789(dci)
	a := disp.NewArea(disp.Bounds())
	a.SetOrigin(a.Bounds().Max.Div(-2))

	// Draw a clock dial with the hour and minute marks and the gopher image
	// in the center.

	r := a.Bounds()
	a.SetColor(black)
	a.Fill(r)
	bgimg, err := png.Decode(strings.NewReader(testdata.GopherPNG))
	if err == nil {
		bgimg = images.Magnify(bgimg, 4, 4, images.Nearest)
	} else {
		bgimg = &image.Uniform{color.Gray{128}}
	}
	offset := r.Size().Sub(bgimg.Bounds().Size()).Div(2)
	sp := bgimg.Bounds().Min
	r = bgimg.Bounds().Sub(sp).Add(r.Min).Add(offset)
	offset = sp.Sub(r.Min)
	a.Draw(r, bgimg, sp, nil, image.Point{}, draw.Src)
	a.SetColor(white)
	for i := 0; i < 60; i++ {
		v0 := image.Pt(0, a.Bounds().Min.Y+1)
		v1 := v0
		if i%5 == 0 {
			v1.Y = v1.Y * 15 / 16 // hour mark
		} else {
			v1.Y = v1.Y * 31 / 32 // minute mark
		}
		alpha := int32(math2d.FullAngle * int64(i) / 60)
		v0 = math2d.Rotate(math2d.F(v0, frac), alpha)
		v1 = math2d.Rotate(math2d.F(v1, frac), alpha)
		a.Line(math2d.I(v0, frac), math2d.I(v1, frac))
	}

	// Setup hands

	sec := Hand{
		Base: [4]image.Point{
			{0, 0}, {0, 0},
			{0, r.Min.Y * 15 / 16}, {0, r.Min.Y * 15 / 16},
		},
	}
	min := Hand{
		Base: [4]image.Point{
			{-3, 0}, {3, 0},
			{2, r.Min.Y * 3 / 4}, {-2, r.Min.Y * 3 / 4},
		},
	}
	hou := Hand{
		Base: [4]image.Point{
			{-4, 0}, {4, 0},
			{3, r.Min.Y / 2}, {-3, r.Min.Y / 2},
		},
	}

	// Run the clock

	t := time.Now()
	for {
		h, m, s := t.Clock()
		if hou.Next(int32(h*4+m/15) << 23 / 6 << 5) {
			hou.Clear(a, bgimg, offset)
			hou.Draw(a)
			min.Draw(a)
			sec.Draw(a)
		}
		if min.Next(int32(m*4+s/15) << 23 / 15 << 5) {
			min.Clear(a, bgimg, offset)
			hou.Draw(a)
			min.Draw(a)
			sec.Draw(a)
		}
		if sec.Next(int32(s) << 25 / 15 << 5) {
			sec.Clear(a, bgimg, offset)
			hou.Draw(a)
			min.Draw(a)
			sec.Draw(a)
		}
		t = t.Add(time.Second)
		time.Sleep(t.Sub(time.Now()))
	}
}
