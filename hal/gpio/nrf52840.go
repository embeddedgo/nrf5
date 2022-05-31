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
	switch {
	case psel <= 1:
		return -1
	case psel <= 5:
		return AIN(psel - 1)
	case psel <= 27:
		return -1
	default:
		return AIN(psel - 23)
	}
}

func apin(a AIN) Pin {
	if a <= 0 {
		return Pin{}
	}
	h := portaddr(0) + uintptr(a)
	if a <= AIN3 {
		h += 1
	} else {
		h += 23
	}
	return Pin{h}
}
