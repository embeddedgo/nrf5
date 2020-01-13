// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package power provides interface to the power managemnt peripheral.
package power

import (
	"embedded/mmio"
	"embedded/rtos"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/te"
	"github.com/embeddedgo/nrf5/p/mmap"
)

type regs struct {
	te.Regs

	resetreas mmio.U32     // 0x400
	_         [9]mmio.U32  //
	ramstatus mmio.U32     // 0x428
	_         [53]mmio.U32 //
	systemoff mmio.U32     // 0x500
	_         [3]mmio.U32  //
	pofcon    mmio.U32     // 0x510
	_         [2]mmio.U32  //
	gpregret  [2]mmio.U32  // 0x51C
	ramon     mmio.U32     // 0x524
	_         [7]mmio.U32  //
	reset     mmio.U32     // 0x544
	_         [3]mmio.U32  //
	ramonb    mmio.U32     // 0x554
	_         [8]mmio.U32  //
	dcdcen    mmio.U32     // 0x578
	_         [225]mmio.U32
	ram       [8]struct{ power, powerset, powerclr mmio.U32 }
}

func r() *regs { return (*regs)(unsafe.Pointer(mmap.POWER_BASE)) }

// IRQEnabled returns EventMask that lists events that have enabled
// generating interrupts..
func IRQEnabled() te.EventMask {
	return r().IRQEnabled()
}

// EnableIRQ enables generating interrupts by events specified by mask.
func EnableIRQ(mask te.EventMask) {
	r().EnableIRQ(mask)
}

// DisableIRQ disables generating interrupts by events specified by mask
func DisableIRQ(mask te.EventMask) {
	r().DisableIRQ(mask)
}

// IRQ returns the IRQ number in NVIC associated to the peripheral.
func IRQ() rtos.IRQ {
	return r().IRQ()
}
