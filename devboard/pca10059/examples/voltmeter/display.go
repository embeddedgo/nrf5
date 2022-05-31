// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/embeddedgo/display/font/subfont/font9/terminus16"
	"github.com/embeddedgo/display/pix"
)

var (
	bgcolor = color.Gray{0}
	fgcolor = color.Gray{255}
)

var disp Display

const (
	Aleft = iota
	Aright
	Acenter
)

// TextArea wraps *pix.TextWriter to add convenient Clear and Printf methods.
type TextArea struct {
	*pix.TextWriter
	Align byte
}

// Clear clears the text area using bgcolor and resets the drawing positon to
// its top-left corner.
func (a *TextArea) Clear() {
	r := a.Area.Bounds()
	a.Area.SetColor(bgcolor)
	a.Area.Fill(r)
	a.Pos = r.Min
}

// WriteString allows to print an aligned multiline text string.
func (a *TextArea) WriteString(text string) {
	for len(text) != 0 {
		var s string
		n := strings.IndexByte(text, '\n')
		if n >= 0 {
			s = text[:n+1]
			text = text[n+1:]
		} else {
			n = len(text)
			s = text
			text = text[n:]
		}
		r := a.Area.Bounds()
		width := pix.StringWidth(s[:n], a.Face)
		switch a.Align {
		case Acenter:
			a.Pos.X = (r.Min.X + r.Max.X - width) / 2
		case Aright:
			a.Pos.X = r.Max.X - width
		default: // Aleft
			a.Pos.X = r.Min.X
		}
		a.TextWriter.WriteString(s)
	}
}

// Printf works like fmt.Printf
func (a *TextArea) Printf(f string, v ...interface{}) {
	a.WriteString(fmt.Sprintf(f, v...))
}

// Display repersents the whole display divided in two areas. The top part takes
// up 1/4 of the screen height and the bottom part takes the rest. In case of
// small OLED displays these two areas are often physically implemented using
// two different colors.
//
// The areas are mainly designed to print text but you can use all pix drawing
// primitives on them using their Area fields.
//
// We use single letters to name the areas in a way that allude to HTML tags.
type Display struct {
	H TextArea // the top part of the display (header)
	P TextArea // the bottom part of the display
}

// Init initializes the display.
func (d *Display) Init(disp *pix.Display) {
	f := terminus16.NewFace(
		terminus16.X0020_007e,
		//terminus16.X00a0_0175,
	)

	r := disp.Bounds()
	r.Max.Y /= 4
	d.H.TextWriter = disp.NewArea(r).NewTextWriter(f)
	d.H.SetColor(fgcolor)

	r = disp.Bounds()
	r.Min.Y = r.Max.Y / 4
	d.P.TextWriter = disp.NewArea(r).NewTextWriter(f)
	d.P.SetColor(fgcolor)
}

// Flush ensures that the latest changes are visible on the screen.
func (d *Display) Flush() {
	d.H.Area.Flush()
}
