// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build nrf52840

package gpio

func ain(pin Pin) AIN {
	psel := pin.PSEL()
	if psel>>5 != 0 {
		return -1
	}
	n := psel & 31
	switch {
	case n < 2:
		return -1
	case n < 6:
		return AIN(n - 2)
	case n < 28:
		return -1
	case n < 32:
		return AIN(n - 24)
	default:
		return -1
	}
}

func apin(a AIN) Pin {
	if a < 0 {
		return Pin{}
	}
	h := portaddr(0) + uintptr(a)
	if a < 4 {
		h += 2
	} else {
		h += 24
	}
	return Pin{h}
}
