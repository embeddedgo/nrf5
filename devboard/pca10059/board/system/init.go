// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	_ "unsafe"

	"github.com/embeddedgo/nrf5/hal/clock"
	"github.com/embeddedgo/nrf5/hal/rtc"
	"github.com/embeddedgo/nrf5/hal/system"
	"github.com/embeddedgo/nrf5/hal/system/timer/rtcst"
)

func init() {
	system.Setup(clock.XTAL, clock.XTAL, true)
	rtcst.Setup(rtc.RTC(1), 0)
}

//go:interrupthandler
func _RTC1_Handler() { rtcst.ISR() }

//go:linkname _RTC1_Handler IRQ17_Handler
