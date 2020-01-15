// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build nrf52832 nrf52833 nrf52840

package te

import (
	"sync/atomic"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/internal"
	"github.com/embeddedgo/nrf5/p/mmap"
)

const numChan = 20

type ppiRegs struct {
	Regs
	internal.PPI
}

func ppi() *ppiRegs {
	return (*ppiRegs)(unsafe.Pointer(mmap.PPI_BASE))
}

func getEventChan(r *Event) (c Chan, en bool) {
	chs := &ppi().CH
	for i := range chs {
		if uintptr(chs[i].EEP.Load()) == uintptr(unsafe.Pointer(r)) {
			return Chan(i), true
		}
	}
	return -1, false
}

func setEventChan(r *Event, c Chan, en bool) {
	chs := &ppi().CH
	for i := range chs {
		eep := &chs[i].EEP
		if uintptr(eep.Load()) == uintptr(unsafe.Pointer(r)) {
			eep.Store(0)
		}
	}
	if en {
		v := uint32(uintptr(unsafe.Pointer(r)))
		p := (*uint32)(unsafe.Pointer(&chs[c].EEP))
		if !atomic.CompareAndSwapUint32(p, 0, v) {
			panic("PPI EEP already in use")
		}
	}
}

func getTaskChan(r *Task) (c Chan, en bool) {
	chs := &ppi().CH
	for i := range chs {
		if uintptr(chs[i].TEP.Load()) == uintptr(unsafe.Pointer(r)) {
			return Chan(i), true
		}
	}
	fts := &ppi().FORK_TEP
	for i := range fts {
		if uintptr(fts[i].Load()) == uintptr(unsafe.Pointer(r)) {
			return Chan(i), true
		}
	}
	return -1, false
}

func setTaskChan(r *Task, c Chan, en bool) {
	chs := &ppi().CH
	for i := range chs {
		tep := &chs[i].TEP
		if uintptr(tep.Load()) == uintptr(unsafe.Pointer(r)) {
			tep.Store(0)
		}
	}
	fts := &ppi().FORK_TEP
	for i := range fts {
		tep := &fts[i]
		if uintptr(tep.Load()) == uintptr(unsafe.Pointer(r)) {
			tep.Store(0)
		}
	}
	if en {
		v := uint32(uintptr(unsafe.Pointer(r)))
		p := (*uint32)(unsafe.Pointer(&chs[c].TEP))
		if atomic.CompareAndSwapUint32(p, 0, v) {
			return
		}
		p = (*uint32)(unsafe.Pointer(&fts[c]))
		if atomic.CompareAndSwapUint32(p, 0, v) {
			return
		}
		panic("PPI TEPs already in use")
	}
}

