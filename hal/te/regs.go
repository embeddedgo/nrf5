// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package te

import (
	"embedded/mmio"
	"embedded/rtos"
	"unsafe"
)

// Regs should be the first field on any Periph struct.
// It takes 0x400 bytes of memory.
type Regs struct {
	tasks    [32]Task
	_        [32]uint32
	events   [32]Event
	_        [32]uint32
	shorts   mmio.U32
	_        [64]uint32
	intEnSet mmio.U32
	intEnClr mmio.U32
	_        [13]uint32
	evtEn    mmio.U32
	evtEnSet mmio.U32
	evtEnClr mmio.U32
	_        [45]uint32
}

// Task returns a pointer to n-th task register.
func (r *Regs) Task(n int) *Task { return &r.tasks[n] }

// Task returns a pointer to n-th event register.
func (r *Regs) Event(n int) *Event { return &r.events[n] }

// EvAll represents all 32 possible events of one peripheral.
const EvAll EventMask = 0xFFFFFFFF

// EventMask is a bitmask that can represent multiple events of one peripheral.
type EventMask uint32

// IRQEnabled returns EventMask that lists events that have enabled
// generating interrupts..
func (r *Regs) IRQEnabled() EventMask {
	return EventMask(r.intEnSet.Load())
}

// EnableIRQ enables generating interrupts by events specified by mask.
func (r *Regs) EnableIRQ(mask EventMask) {
	r.intEnSet.Store(uint32(mask))
}

// DisableIRQ disables generating interrupts by events specified by mask
func (r *Regs) DisableIRQ(mask EventMask) {
	r.intEnClr.Store(uint32(mask))
}

// PPIEnabled returns EventMask that lists events that have enabled routing to
// PPI. Only RTC supports this method.
func (r *Regs) PPIEnabled() EventMask {
	return EventMask(r.evtEnSet.Load())
}

// EnablePPI enables routing to PPI for events specified by mask. Only RTC
// supports this method.
func (r *Regs) EnablePPI(mask EventMask) {
	r.evtEnSet.Store(uint32(mask))
}

// DisablePPI disables routing to PPI for events specified by mask. Only RTC
// supports this method.
func (r *Regs) DisablePPI(mask EventMask) {
	r.evtEnClr.Store(uint32(mask))
}

// LoadSHORTS returns value of the SHORTS register.
func (r *Regs) LoadSHORTS() uint32 {
	return r.shorts.Load()
}

// StoreSHORTS stores s to the SHORTS register.
func (r *Regs) StoreSHORTS(s uint32) {
	r.shorts.Store(s)
}

// IRQ returns the IRQ number in NVIC associated to the peripheral.
func (r *Regs) IRQ() rtos.IRQ {
	addr := uintptr(unsafe.Pointer(r))
	return rtos.IRQ(addr >> 12 & 0x1FF)
}
