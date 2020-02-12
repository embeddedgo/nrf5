// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package te

import (
	"embedded/mmio"
	"embedded/rtos"
	"unsafe"
)

// Event represents a peripheral register that is used to record an occurence
// of one kind of events.
type Event struct {
	u32 mmio.U32
}

// IsSet reports whether any event has been generated since the last clearing
// of r.
func (r *Event) IsSet() bool {
	return r.u32.Load() != 0
}

// Clear clears the event flag in r
func (r *Event) Clear() {
	r.u32.Store(0)
}

func regsMask(r *Event) (*Regs, uint32) {
	addr := r.u32.Addr()
	n := addr&0xFFF/4 - 64
	return (*Regs)(unsafe.Pointer(addr &^ 0xFFF)), 1 << n
}

// IRQEnabled reports whether the events recorded in r generete interrupts.
func (r *Event) IRQEnabled() bool {
	rr, mask := regsMask(r)
	return rr.intEnSet.Load()&mask != 0
}

// EnableIRQ enables generating of interrupt by events recorded in r.
func (r *Event) EnableIRQ() {
	rr, mask := regsMask(r)
	rr.intEnSet.Store(mask)
}

// DisableIRQ disables generating of interrupts by events recorded in r.
func (r *Event) DisableIRQ() {
	rr, mask := regsMask(r)
	rr.intEnClr.Store(mask)
}

// PPIEnabled reports whether the event has enabled routing to PPI. Only RTC
// supports this method.
func (r *Event) PPIEnabled() bool {
	rr, mask := regsMask(r)
	return rr.evtEnSet.Load()&mask != 0
}

// EnablePPI enables routing the event to PPI. Only RTC supports this method.
func (r *Event) EnablePPI() {
	rr, mask := regsMask(r)
	rr.evtEnSet.Store(mask)
}

// DisablePPI disable routing the event to PPI. Only RTC supports this method.
func (r *Event) DisablePPI() {
	rr, mask := regsMask(r)
	rr.evtEnClr.Store(mask)
}

// IRQ returns NVIC IRQ number associated to the peripheral.
func (r *Event) IRQ() rtos.IRQ {
	rr, _ := regsMask(r)
	return rr.IRQ()
}

// Chan returns the PPI/DPPI channel the event is connected to and reports
// whether publishing events on this channel is enabled. In case of nRF52- which
// PPI peripheral allows to connect one event to multiple PPI channels Chan
// returns the first channel found.
func (r *Event) Chan() (c Chan, en bool) {
	return getEventChan(r)
}

// SetChan connects the event to the PPI/DPPI channel. En allows to enable or
// disable publishing events on connected channel.
//
// SetChan allows to write portable code that can work with PPI (nRF52-) and
// DPPI (nRF53) but has some limitations in case of PPI:
//
// - you cannot connect one event to multiple channels,
//
// - if en is false the channel is set to -1 which as a result disconnects the
//   event from all channels.
//
// The one advantage of PPI over DPPI is its ablility to conect one event to
// multiple channels. If you need this feature use SetEEP from ppi package.
func (r *Event) SetChan(c Chan, en bool) {
	setEventChan(r, c, en)
}
