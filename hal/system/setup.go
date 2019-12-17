// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import "github.com/embeddedgo/nrf5/hal/clock"

// Setup setups nRF51 to operate using specified HFCLK and LFCLK clock sources..
func Setup(hfsrc, lfsrc clock.Source, lfena bool) {
	clock.StoreLFCLKSRC(lfsrc)
	if hfsrc == clock.XTAL {
		clock.HFCLKSTART.Task().Trigger()
	}
	if lfena {
		clock.LFCLKSTART.Task().Trigger()
	}
	for {
		src, run := clock.LoadHFCLKSTAT()
		if src == hfsrc && run {
			break
		}
	}
	for lfena {
		src, run := clock.LoadLFCLKSTAT()
		if src == lfsrc && run {
			break
		}
	}
}
