// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package sdm

import "unsafe"

// LFCLKA represents accuracy of external low-frequency clock source.
type LFCLKA uint8

const (
	LFCLK250ppm LFCLKA = 0  // 250 ppm (default)
	LFCLK500ppm LFCLKA = 1  // 500 ppm
	LFCLK150ppm LFCLKA = 2  // 150 ppm
	LFCLK100ppm LFCLKA = 3  // 100 ppm
	LFCLK75ppm  LFCLKA = 4  // 75 ppm
	LFCLK50ppm  LFCLKA = 5  // 50 ppm
	LFCLK30ppm  LFCLKA = 6  // 30 ppm
	LFCLK20ppm  LFCLKA = 7  // 20 ppm
	LFCLK10ppm  LFCLKA = 8  // 10 ppm
	LFCLK5ppm   LFCLKA = 9  // 5 ppm
	LFCLK2ppm   LFCLKA = 10 // 2 ppm
	LFCLK1ppm   LFCLKA = 11 // 1 ppm
)

// LFCLKS represents possible low-frequency clock sources.
type LFCLKS uint8

const (
	LFCLKRC    LFCLKS = 0 // RC oscillator.
	LFCLKXTAL  LFCLKS = 1 // crystal oscillator.
	LFCLKSYNTH LFCLKS = 2 // synthesized from HFCLK.
)

// LFCLKC represents configuration of low-frequency clock.
type LFCLKC struct {
	Source     LFCLKS // clock source
	CalInt     uint8  // RC oscilator calibration interval in 1/4 second units
	CalTempInt uint8  // how often calibrate RC oscilator if temp. has changed
	Accuracy   LFCLKA // external clock accuracy
}

var sdfh func(id, pc, info uint32)

func EnableSoftdevice(lfclkc *LFCLKC, faultHandler func(id, pc, info uint32)) error {
	sdfh = faultHandler
	return mkerr(enableSoftdevice(lfclkc))
}

func DisableSoftdevice() error {
	return mkerr(disableSoftdevice())
}

func SoftdeviceEnabled() (en bool, err error) {
	e := softdeviceEnabled(&en)
	return en, mkerr(e)
}

func SetSoftdeviceVectorTableBase(addr unsafe.Pointer) error {
	return mkerr(setSoftdeviceVectorTableBase(addr))
}
