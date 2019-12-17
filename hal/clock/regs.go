// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package clock provides interface to manage clocks source/generation.
package clock

import (
	"embedded/mmio"
	"embedded/rtos"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/internal/mmap"
	"github.com/embeddedgo/nrf5/hal/te"
)

type regs struct {
	te.Regs

	_            [2]mmio.U32
	hfclkrun     mmio.U32
	hfclkstat    mmio.U32
	_            mmio.U32
	lfclkrun     mmio.U32
	lfclkstat    mmio.U32
	lfclksrccopy mmio.U32
	_            [62]mmio.U32
	lfclksrc     mmio.U32
	_            [7]mmio.U32
	ctiv         mmio.U32
	_            [5]mmio.U32
	xtalfreq     mmio.U32
	_            [2]mmio.U32
	traceconfig  mmio.U32
}

func r() *regs {
	return (*regs)(unsafe.Pointer(mmap.APB_BASE))
}

func IRQ() rtos.IRQ {
	return r().IRQ()
}

func IRQEnabled() te.EventMask {
	return r().IRQEnabled()
}

func EnableIRQ(mask te.EventMask) {
	r().EnableIRQ(mask)
}

func DisableIRQ(mask te.EventMask) {
	r().DisableIRQ(mask)
}
