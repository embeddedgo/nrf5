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
// of event.
type Event struct {
	u32 mmio.U32
}

// IsSet reports whether the event has been generated since the last clearing of
// r. The unhandled event remains in active state as long as its software
// handler (usually an interrupt handler) clears it.
func (r *Event) IsSet() bool {
	return r.u32.Load() != 0
}

// Clear clears active event flag in r.
func (r *Event) Clear() {
	r.u32.Store(0)
}

func regsMask(r *Event) (*Regs, uint32) {
	addr := r.u32.Addr()
	n := addr&0xFFF/4 - 64
	return (*Regs)(unsafe.Pointer(addr &^ 0xFFF)), 1 << n
}

// IRQEnabled reports whether the event in active state generates interrupt.
func (r *Event) IRQEnabled() bool {
	rr, mask := regsMask(r)
	return rr.intEnSet.Load()&mask != 0
}

// EnableIRQ enables generating of interrupt by the event in active state.
func (r *Event) EnableIRQ() {
	rr, mask := regsMask(r)
	rr.intEnSet.Store(mask)
}

// DisableIRQ disables generating of interrupt by event.
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

// EnablePPI enables routing event to PPI. Only RTC supports this method.
func (r *Event) EnablePPI() {
	rr, mask := regsMask(r)
	rr.evtEnSet.Store(mask)
}

// DisablePPI disable routing event to PPI. Only RTC supports this method.
func (r *Event) DisablePPI() {
	rr, mask := regsMask(r)
	rr.evtEnClr.Store(mask)
}

// IRQ returns NVIC IRQ number associated to peripheral.
func (r *Event) IRQ() rtos.IRQ {
	rr, _ := regsMask(r)
	return rr.IRQ()
}

