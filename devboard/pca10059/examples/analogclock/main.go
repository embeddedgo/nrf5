// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/embeddedgo/display/images"
	"github.com/embeddedgo/display/math2d"
	"github.com/embeddedgo/display/pix"
	"github.com/embeddedgo/display/pix/displays"
	"github.com/embeddedgo/display/pix/driver/imgdrv"
	"github.com/embeddedgo/display/testdata"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/system"

	"github.com/embeddedgo/nrf5/dci/tftdci"
	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/spim"
	"github.com/embeddedgo/nrf5/hal/spim/spim3"
)

const frac = 8

var (
	black = color.Gray{}
	white = color.Gray{255}
	red   = color.RGBA{255, 0, 0, 255}
)

var bgimg = images.Magnify(
	images.NewImmRGB16(image.Rect(0, 0, 48, 48), testdata.Gopher48x48RGB16),
	4, 4, images.Nearest,
)

func drawDial(a *pix.Area, r image.Rectangle) int {
	a.SetColor(black)
	a.Fill(r)

	// Draw the background image in the center of r
	o := r.Size().Sub(bgimg.Bounds().Size()).Div(2)
	a.Draw(r.Add(o), bgimg, bgimg.Bounds().Min, nil, image.Point{}, draw.Src)

	// Calculate the length of hour and minute marks on the assumption that
	// a is centerad around (0, 0) so r.Min is roughly equal to -r.Max.
	d := r.Max.X
	if d < r.Max.Y {
		d = r.Max.Y
	}
	d--
	// Hour and minute marks at twelve o'clock.
	hour := [2]image.Point{{0, -d}, {0, -d * 30 / 32}}
	min := [2]image.Point{{0, -d}, {0, -d * 31 / 32}}

	// Draw the marks around the dial.
	a.SetColor(white)
	for i := 0; i < 60; i++ {
		v := min
		if i%5 == 0 {
			v = hour
		}
		alpha := int32(math2d.FullAngle * int64(i) / 60)
		v[0] = math2d.Rotate(math2d.F(v[0], frac), alpha)
		v[1] = math2d.Rotate(math2d.F(v[1], frac), alpha)
		a.Line(math2d.I(v[0], frac), math2d.I(v[1], frac))
	}

	// Return the length of the second hand.
	return d
}

func drawDial1(ar *pix.Area, r image.Rectangle) {
	img := images.NewRGB16(ar.Bounds())
	disp := pix.NewDisplay(imgdrv.New(img))
	a := disp.NewArea(disp.Bounds())
	a.SetOrigin(ar.Bounds().Min)

	// Draw the background image in the center of a
	o := r.Size().Sub(bgimg.Bounds().Size()).Div(2)
	a.Draw(r.Add(o), bgimg, bgimg.Bounds().Min, nil, image.Point{}, draw.Src)

	// Calculate the length of hour and minute marks on the assumption that
	// a is centerad around (0, 0) so r.Min is roughly equal to -r.Max.
	d := r.Max.X
	if d < r.Max.Y {
		d = r.Max.Y
	}
	d--
	// Hour and minute marks at twelve o'clock.
	hour := [2]image.Point{{0, -d}, {0, -d * 30 / 32}}
	min := [2]image.Point{{0, -d}, {0, -d * 31 / 32}}

	// Draw the marks around the dial.
	a.SetColor(white)
	for i := 0; i < 60; i++ {
		v := min
		if i%5 == 0 {
			v = hour
		}
		alpha := int32(math2d.FullAngle * int64(i) / 60)
		v[0] = math2d.Rotate(math2d.F(v[0], frac), alpha)
		v[1] = math2d.Rotate(math2d.F(v[1], frac), alpha)
		a.Line(math2d.I(v[0], frac), math2d.I(v[1], frac))
	}

	ar.Draw(ar.Bounds(), img, img.Bounds().Min, nil, image.Point{}, draw.Src)
}

func drawDial2(ar *pix.Area, r image.Rectangle, h1, h2, h3 *Hand) {
	img := images.NewRGB16(ar.Bounds())
	disp := pix.NewDisplay(imgdrv.New(img))
	a := disp.NewArea(disp.Bounds())
	a.SetOrigin(ar.Bounds().Min)

	// Draw the background image in the center of a
	o := r.Size().Sub(bgimg.Bounds().Size()).Div(2)
	a.Draw(r.Add(o), bgimg, bgimg.Bounds().Min, nil, image.Point{}, draw.Src)

	// Calculate the length of hour and minute marks on the assumption that
	// a is centerad around (0, 0) so r.Min is roughly equal to -r.Max.
	d := r.Max.X
	if d < r.Max.Y {
		d = r.Max.Y
	}
	d--
	// Hour and minute marks at twelve o'clock.
	hour := [2]image.Point{{0, -d}, {0, -d * 30 / 32}}
	min := [2]image.Point{{0, -d}, {0, -d * 31 / 32}}

	// Draw the marks around the dial.
	a.SetColor(white)
	for i := 0; i < 60; i++ {
		v := min
		if i%5 == 0 {
			v = hour
		}
		alpha := int32(math2d.FullAngle * int64(i) / 60)
		v[0] = math2d.Rotate(math2d.F(v[0], frac), alpha)
		v[1] = math2d.Rotate(math2d.F(v[1], frac), alpha)
		a.Line(math2d.I(v[0], frac), math2d.I(v[1], frac))
	}
	h1.Draw(a)
	h2.Draw(a)
	h3.Draw(a)
	ar.Draw(ar.Bounds(), img, img.Bounds().Min, nil, image.Point{}, draw.Src)
}

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

func (h *Hand) Clear(a *pix.Area) {
	rect := a.Rect()
	bounds := a.Bounds()
	newr := h.bounds().Add(rect.Min.Sub(bounds.Min))
	delta := newr.Min.Sub(rect.Min)
	a.SetRect(newr)
	a.SetOrigin(bounds.Min.Add(delta))
	drawDial1(a, bounds)
	a.SetRect(rect)
	a.SetOrigin(bounds.Min)
}

func (h *Hand) Update1(a *pix.Area, h1, h2, h3 *Hand) {
	rect := a.Rect()
	bounds := a.Bounds()
	hb := h.bounds()
	h.cur = h.next
	hb = hb.Union(h.bounds())
	newr := hb.Add(rect.Min.Sub(bounds.Min))
	delta := newr.Min.Sub(rect.Min)
	a.SetRect(newr)
	a.SetOrigin(bounds.Min.Add(delta))
	drawDial2(a, bounds, h1, h2, h3)
	a.SetRect(rect)
	a.SetOrigin(bounds.Min)
}

func (h *Hand) Next(alpha int32) bool {
	for i := range h.Base {
		v := math2d.F(h.Base[i], frac)
		v = math2d.Rotate(v, alpha)
		h.next[i] = math2d.I(v, frac)
	}
	return h.next != h.cur
}

func (h *Hand) Update() {
	h.cur = h.next
}

func main() {
	// Declare signal lines of external peripherals

	// IPS display.
	var ips struct{ scl, sda, res, dc, cs, blk gpio.Pin }

	// Assign GPIO pins to the signals of external peripherals

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
	dci := tftdci.NewSPIM(spi, ips.dc, spim.CPOL0|spim.CPHA0, spim.F8MHz, spim.F16MHz)
	dci.UseCSN(ips.cs, false)

	// Reset

	ips.res.Clear()
	time.Sleep(time.Millisecond)
	ips.res.Set()
	ips.blk.Set()

	// Initialize display and create the drawing area with origin in the center

	disp := displays.ERTFTM_1i54_240x240_IPS_ST7789().New(dci)
	a := disp.NewArea(disp.Bounds())
	a.SetOrigin(a.Bounds().Max.Div(-2))

	// Draw a clock dial with the hour and minute marks and the gopher image
	// in the center.

	d := drawDial(a, a.Bounds())

	// Setup hands

	sec := Hand{
		Base: [4]image.Point{
			{0, d / 8}, {0, d / 8},
			{0, -d}, {0, -d},
		},
	}
	min := Hand{
		Base: [4]image.Point{
			{-3, d / 8}, {3, d / 8},
			{3, -d * 3 / 4}, {-3, -d * 3 / 4},
		},
	}
	hou := Hand{
		Base: [4]image.Point{
			{-4, d / 8}, {4, d / 8},
			{4, -d / 2}, {-4, -d / 2},
		},
	}

	// Run the clock
	t := time.Now()
	/*for {
	  	h, m, s := t.Clock()
	  	h26 := int32(h*8+m*10/75) << 23 // positive 6.26 fixed point hour
	  	hou.Next(h26 / 3 << 3)
	  	hou.Update()
	  	m25 := int32(m*4+s/15) << 23 // positive 7.25 fixed point minute
	  	min.Next(m25 / 15 << 5)
	  	min.Update()
	  	s25 := int32(s) << 25 // positive 7.25 fixed point second
	  	sec.Next(s25 / 15 << 5)
	  	sec.Update()
	  	drawDial(a, a.Bounds())
	  	hou.Draw(a)
	  	min.Draw(a)
	  	sec.Draw(a)
	  	t = t.Add(time.Second)
	  	time.Sleep(t.Sub(time.Now()))
	  }
	*/
	/*for {
		h, m, s := t.Clock()
		h26 := int32(h*8+m*10/75) << 23 // positive 6.26 fixed point hour
		if hou.Next(h26 / 3 << 3) {
			hou.Clear(a)
			hou.Update()
			hou.Draw(a)
			min.Draw(a)
			sec.Draw(a)
		}
		m25 := int32(m*4+s/15) << 23 // positive 7.25 fixed point minute
		if min.Next(m25 / 15 << 5) {
			min.Clear(a)
			min.Update()
			hou.Draw(a)
			min.Draw(a)
			sec.Draw(a)
		}
		s25 := int32(s) << 25 // positive 7.25 fixed point second
		if sec.Next(s25 / 15 << 5) {
			sec.Clear(a)
			sec.Update()
			hou.Draw(a)
			min.Draw(a)
			sec.Draw(a)
		}
		t = t.Add(time.Second)
		time.Sleep(t.Sub(time.Now()))
	}*/
	for {
		h, m, s := t.Clock()
		h26 := int32(h*8+m*10/75) << 23 // positive 6.26 fixed point hour
		if hou.Next(h26 / 3 << 3) {
			hou.Update1(a, &hou, &min, &sec)
		}
		m25 := int32(m*4+s/15) << 23 // positive 7.25 fixed point minute
		if min.Next(m25 / 15 << 5) {
			min.Update1(a, &hou, &min, &sec)
		}
		s25 := int32(s) << 25 // positive 7.25 fixed point second
		if sec.Next(s25 / 15 << 5) {
			sec.Update1(a, &hou, &min, &sec)
		}
		t = t.Add(time.Second)
		time.Sleep(t.Sub(time.Now()))
	}
}
