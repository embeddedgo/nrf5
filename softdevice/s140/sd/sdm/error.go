// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package sdm

import "github.com/embeddedgo/nrf5/softdevice/s140/sd"

const ErrorBase = 0x1000

type Error uint32

const (
	ErrLFCLKS Error = ErrorBase + 0 // unknown LFCLK source
	ErrIntCfg Error = ErrorBase + 1 // incorrect interrupt configuration
	ErrCLENR0 Error = ErrorBase + 2 // incorrect CLENR0
)

var errStr = [...]string{
	"unknown LFCLK source",
	"incorrect interrupt configuration",
	"incorrect CLENR0",
}

func (e Error) Error() string {
	s := "???"
	if e >= ErrorBase {
		if e -= ErrorBase; int(e) < len(errStr) {
			s = errStr[e]
		}
	}
	return "sdm: " + s
}

func mkerr(e uint32) error {
	switch {
	case e == 0:
		return nil
	case e >= ErrorBase:
		return Error(e)
	default:
		return sd.Error(e)
	}
}
