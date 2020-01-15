// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package te

import "github.com/embeddedgo/nrf5/hal/internal"

// NumChan is the number of channels available to the application. It does not
// include pre-programmed channels.
const NumChan = numChan

// Chan represents PPI/DPPI channel.
type Chan int8

var unusedChannels uint32 = 1<<numChan - 1

// AllocChan returns first unused PPI/DPPI channel. It returns negative number
// if all channels are in use.
func AllocChan() Chan {
	return Chan(internal.BitAlloc(&unusedChannels))
}

// Free disables channel and returns it to the pool of unused channels.
func (c Chan) Free() {
	m := c.Mask()
	m.Disable()
	internal.BitFree(&unusedChannels, uint32(m))
}

// Mask returns a Channels bitfield with c bit set.
func (c Chan) Mask() Channels {
	return Channels(1) << c
}

// Enabled reports whether channel c is enabled.
func (c Chan) Enabled() bool {
	return Enabled()&c.Mask() != 0
}

// Enable atomically enables channel c.
func (c Chan) Enable() {
	c.Mask().Enable()
}

// Enable atomically disables channel c.
func (c Chan) Disable() {
	c.Mask().Disable()
}

// Channels is a bitfield that represents multiple PPI/DPPI channels.
type Channels uint32

// Enabled returns enabled channels.
func Enabled() Channels {
	return Channels(ppi().CHEN.Load())
}

// Enable atomically enables channels.
func (c Channels) Enable() {
	ppi().CHENSET.Store(uint32(c))
}

// Disable atomically disables channels.
func (c Channels) Disable() {
	ppi().CHENCLR.Store(uint32(c))
}
