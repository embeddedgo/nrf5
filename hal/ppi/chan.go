// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppi

import (
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/internal"
	"github.com/embeddedgo/nrf5/hal/te"
)

// Pre-programmed channels.
const (
	TIMER0_COMPARE0__RADIO_TXEN    Chan = 20
	TIMER0_COMPARE0__RADIO_RXEN    Chan = 21
	TIMER0_COMPARE1__RADIO_DISABLE Chan = 22
	RADIO_BCMATCH__AAR_START       Chan = 23
	RADIO_READY__CCM_KSGEN         Chan = 24
	RADIO_ADDRESS__CCM_CRYPT       Chan = 25
	RADIO_ADDRESS__TIMER0_CAPTURE1 Chan = 26
	RADIO_END__TIMER0_CAPTURE2     Chan = 27
	RTC0_COMPARE0__RADIO_TXEN      Chan = 28
	RTC0_COMPARE0__RADIO_RXEN      Chan = 29
	RTC0_COMPARE0__TIMER0_CLEAR    Chan = 30
	RTC0_COMPARE0__TIMER0_START    Chan = 31
)

// Chan represents PPI channel. There are 32 channels numbered from 0 to 31.
// Channels from 20 to 31 are pre-programmed.
type Chan int8

var unusedChannels uint32 = 1<<20 - 1

// ChanAlloc returns first unused PPI channel. It returns negative number if
// there is no free channels.
func ChanAlloc() Chan {
	return Chan(internal.BitAlloc(&unusedChannels))
}

// Free disables channel and returns it to the pool of unused channels.
func (c Chan) Free() {
	m := c.Mask()
	m.Disable()
	internal.BitFree(&unusedChannels, uint32(m))
}

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

// EEP returns the value of Event End Point register for channel c.
func (c Chan) EEP() *te.Event {
	return (*te.Event)(unsafe.Pointer(uintptr(r().ch[c].eep.Load())))
}

// SetEEP sets the value of Event End Point register for channel c.
func (c Chan) SetEEP(e *te.Event) {
	r().ch[c].eep.Store(uint32(uintptr(unsafe.Pointer(e))))
}

// TEP returns the value of Task End Point register for channel c.
func (c Chan) TEP() *te.Task {
	return (*te.Task)(unsafe.Pointer(uintptr(r().ch[c].tep.Load())))
}

// SetTEP sets the value of Task End Point register for channel c.
func (c Chan) SetTEP(t *te.Task) {
	r().ch[c].tep.Store(uint32(uintptr(unsafe.Pointer(t))))
}

// TEP1 returns the value of Fork Task End Point register for channel c. nRF52.
func (c Chan) TEP1() *te.Task {
	return (*te.Task)(unsafe.Pointer(uintptr(r().forktep[c].Load())))
}

// SetTEP1 sets the value of Fork Task End Point register for channel c. nRF52.
func (c Chan) SetTEP1(t *te.Task) {
	r().forktep[c].Store(uint32(uintptr(unsafe.Pointer(t))))
}
