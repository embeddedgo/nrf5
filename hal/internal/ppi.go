// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import "embedded/mmio"

type PPI struct {
	_        [64]uint32
	CHEN     mmio.U32
	CHENSET  mmio.U32
	CHENCLR  mmio.U32
	_        uint32
	CH       [20]struct{ EEP, TEP mmio.U32 }
	_        [148]uint32
	CHG      [6]mmio.U32
	_        [62]uint32
	FORK_TEP [32]mmio.U32
}
