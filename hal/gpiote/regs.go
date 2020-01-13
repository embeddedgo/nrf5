// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpiote

import (
	"embedded/mmio"
	"embedded/rtos"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/te"
	"github.com/embeddedgo/nrf5/p/mmap"
)

type regs struct {
	te.Regs

	_      [68]mmio.U32
	config [8]mmio.U32
}

func r() *regs {
	return (*regs)(unsafe.Pointer(mmap.GPIOTE_BASE))
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
