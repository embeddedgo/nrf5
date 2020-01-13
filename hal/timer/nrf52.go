// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build nrf52832 nrf52840

package timer

import (
	"unsafe"

	"github.com/embeddedgo/nrf5/p/mmap"
)

func timer(n int) *Periph {
	switch {
	case uint(n) > 4:
		return nil
	case n <= 2:
		return (*Periph)(unsafe.Pointer(mmap.TIMER0_BASE + uintptr(n)*0x1000))
	default:
		return (*Periph)(unsafe.Pointer(mmap.TIMER3_BASE + uintptr(n-3)*0x1000))
	}
}
