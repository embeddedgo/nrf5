// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpio

import (
	"unsafe"

	"github.com/embeddedgo/nrf5/p/mmap"
)

func portaddr(n int) uintptr {
	if uint(n) > 1 {
		return 0
	}
	return (mmap.P0_BASE + 0x500) + uintptr(n)*0x300
}

func portnum(p *Port) int {
	n := int(uintptr(unsafe.Pointer(p))-(mmap.P0_BASE+0x500)) / 0x300
	if uint(n) > 1 {
		return -1
	}
	return n
}
