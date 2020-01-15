// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build nrf52832 nrf52833 nrf52840

package ppi

import (
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/internal"
	"github.com/embeddedgo/nrf5/hal/te"
	"github.com/embeddedgo/nrf5/p/mmap"
)

const numGroup = 6

type regs struct {
	te.Regs
	internal.PPI
}

func r() *regs {
	return (*regs)(unsafe.Pointer(mmap.PPI_BASE))
}

func loadEEP(c te.Chan) *te.Event {
	return (*te.Event)(unsafe.Pointer(uintptr(r().CH[c].EEP.Load())))
}

func storeEEP(c te.Chan, e *te.Event) {
	r().CH[c].EEP.Store(uint32(uintptr(unsafe.Pointer(e))))
}

func loadTEP(c te.Chan) *te.Task {
	return (*te.Task)(unsafe.Pointer(uintptr(r().CH[c].TEP.Load())))
}

func storeTEP(c te.Chan, t *te.Task) {
	r().CH[c].TEP.Store(uint32(uintptr(unsafe.Pointer(t))))
}

func loadTEP1(c te.Chan) *te.Task {
	return (*te.Task)(unsafe.Pointer(uintptr(r().FORK_TEP[c].Load())))
}

func storeTEP1(c te.Chan, t *te.Task) {
	r().FORK_TEP[c].Store(uint32(uintptr(unsafe.Pointer(t))))
}
