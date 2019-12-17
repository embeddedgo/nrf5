// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmap

const (
	APB_BASE uintptr = 0x40000000 // accessed by APB,
	AHB_BASE uintptr = 0x50000000 // accessed by AHB.

	FICR_BASE uintptr = 0x10000000
)
