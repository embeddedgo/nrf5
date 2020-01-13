// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package psel

import (
	"embedded/mmio"

	"github.com/embeddedgo/nrf5/hal/gpio"
)

type Reg struct {
	u32 mmio.U32
}

func (r *Reg) Load() (psel gpio.PSEL, en bool) {
	pseldis := r.u32.Load()
	return gpio.PSEL(pseldis), pseldis>>31 == 0
}

func (r *Reg) Store(psel gpio.PSEL, en bool) {
	pseldis := uint32(int32(psel))
	if !en {
		pseldis |= 1 << 31
	}
	r.u32.Store(pseldis)
}
