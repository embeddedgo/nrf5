// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import "sync/atomic"

func BitAlloc(unused *uint32) int {
	for {
		m := atomic.LoadUint32(unused)
		if m == 0 {
			return -1
		}
		i := uint(0)
		for (m>>i)&1 == 0 {
			i++
		}
		if atomic.CompareAndSwapUint32(unused, m, m&^(1<<i)) {
			return int(i)
		}
	}
}

func BitFree(unused *uint32, mask uint32) {
	for {
		um := atomic.LoadUint32(unused)
		if um&mask != 0 {
			panic("already unused")
		}
		if atomic.CompareAndSwapUint32(unused, um, um|mask) {
			return
		}
	}
}
