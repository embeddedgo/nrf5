// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppi

import (
	"embedded/mmio"
	"unsafe"
	"rtos"

	"github.com/embeddedgo/nrf5/hal/te"
	"github.com/embeddedgo/nrf5/hal/internal/mmap"
)

type channel struct {
	eep mmio.U32
	tep mmio.U32
}

type regs struct {
	te.Regs

	_       [64]mmio.U32
	chen    mmio.U32
	chenset mmio.U32
	chenclr mmio.U32
	_       mmio.U32
	ch      [20]channel
	_       [148]mmio.U32
	chg     [6]mmio.U32
	_       [62]mmio.U32
	forktep [32]mmio.U32
}

func r() *regs {
	return (*regs)(unsafe.Pointer(mmap.APB_BASE + 0x1F000))
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
