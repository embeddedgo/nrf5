// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppi

// Channels is a bitfield that represents multiple PPI channels.
type Channels uint32

// Enabled returns enabled channels.
func Enabled() Channels {
	return Channels(r().chen.Load())
}

// Enable atomically enables channels c.
func (c Channels) Enable() {
	r().chenset.Store(uint32(c))
}

// Disable atomically disables channels.
func (c Channels) Disable() {
	r().chenclr.Store(uint32(c))
}
