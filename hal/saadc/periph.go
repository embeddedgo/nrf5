// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package saadc

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/gpio"
	"github.com/embeddedgo/nrf5/hal/internal"
	"github.com/embeddedgo/nrf5/hal/te"
	"github.com/embeddedgo/nrf5/p/mmap"
)

type Periph struct {
	te.Regs

	status mmio.U32
	_      [63]uint32
	enable mmio.U32
	_      [3]uint32
	ch     [8]struct {
		pselp  mmio.U32
		pseln  mmio.U32
		config mmio.U32
		limit  mmio.U32
	}
	_          [24]uint32
	resolution mmio.U32
	oversample mmio.U32
	samplerate mmio.U32
	_          [12]uint32
	ptr        mmio.U32
	maxcnt     mmio.U32
	amount     mmio.U32
}

func SAADC(n int) *Periph {
	if n != 0 {
		return nil
	}
	return (*Periph)(unsafe.Pointer(mmap.SAADC_BASE))
}

type Task uint8

const (
	START     Task = 0 // Start the SAADC and prepare the result buffer in RAM.
	SAMPLE    Task = 1 // Take one SAADC sample.
	STOP      Task = 2 // Stop the SAADC and terminate all on-going conversions.
	CALIBRATE Task = 3 // Start offset auto-calibration.
)

type Event uint8

const (
	STARTED       Event = 0  // The SAADC has started.
	END           Event = 1  // The SAADC has filled up the result buffer.
	DONE          Event = 2  // A conversion task has been completed.
	RESULTDONE    Event = 3  // Result ready for transfer to RAM.
	CALIBRATEDONE Event = 4  // Calibration is complete.
	STOPPED       Event = 5  // The SAADC has stopped.
	CH0LIMITH     Event = 6  // Last result is equal or above CH[0].LIMIT.HIGH
	CH0LIMITL     Event = 7  // Last result is equal or below CH[0].LIMIT.LOW
	CH1LIMITH     Event = 8  // Last result is equal or above CH[1].LIMIT.HIGH
	CH1LIMITL     Event = 9  // Last result is equal or below CH[1].LIMIT.LOW
	CH2LIMITH     Event = 10 // Last result is equal or above CH[2].LIMIT.HIGH
	CH2LIMITL     Event = 11 // Last result is equal or below CH[2].LIMIT.LOW
	CH3LIMITH     Event = 12 // Last result is equal or above CH[3].LIMIT.HIGH
	CH3LIMITL     Event = 13 // Last result is equal or below CH[3].LIMIT.LOW
	CH4LIMITH     Event = 14 // Last result is equal or above CH[4].LIMIT.HIGH
	CH4LIMITL     Event = 15 // Last result is equal or below CH[4].LIMIT.LOW
	CH5LIMITH     Event = 16 // Last result is equal or above CH[5].LIMIT.HIGH
	CH5LIMITL     Event = 17 // Last result is equal or below CH[5].LIMIT.LOW
	CH6LIMITH     Event = 18 // Last result is equal or above CH[6].LIMIT.HIGH
	CH6LIMITL     Event = 19 // Last result is equal or below CH[6].LIMIT.LOW
	CH7LIMITH     Event = 20 // Last result is equal or above CH[7].LIMIT.HIGH
	CH7LIMITL     Event = 21 // Last result is equal or below CH[7].LIMIT.LOW
)

func (p *Periph) Task(t Task) *te.Task    { return p.Regs.Task(int(t)) }
func (p *Periph) Event(e Event) *te.Event { return p.Regs.Event(int(e)) }

// LoadSTALSTAT reports busy status of the SAADC. It returns true if the
// conversion is in progress and false if there is no on-going conversions.
func (p *Periph) LoadSTATUS() bool {
	return p.status.Load()&1 != 0
}

// LoadENABLE reports whether the SAADC peripheral is enabled.
func (p *Periph) LoadENABLE() bool {
	return p.enable.Load()&1 != 0
}

// StoreENABLE enables or disables the SAADC peripheral.
func (p *Periph) StoreENABLE(en bool) {
	p.enable.Store(internal.BoolToUint32(en))
}

// LoadPSELP returns the analog input used as of positive input of channel ch.
func (p *Periph) LoadPSELP(ch int) gpio.AIN {
	return gpio.AIN(p.ch[ch].pselp.Load())
}

// StorePSELP configures the positive input of channel ch to use analog input a.
func (p *Periph) StorePSELP(ch int, a gpio.AIN) {
	p.ch[ch].pselp.Store(uint32(a))
}

// LoadPSELN returns the analog input used as of negative input of channel ch.
func (p *Periph) LoadPSELN(ch int) gpio.AIN {
	return gpio.AIN(p.ch[ch].pseln.Load())
}

// StorePSELN configures the negative input of channel ch to use analog input a.
func (p *Periph) StorePSELN(ch int, a gpio.AIN) {
	p.ch[ch].pseln.Store(uint32(a))
}

type Config uint32

const (
	Rp       Config = 3 << 0 // Positive channel pull up/down resistors control
	RpNone   Config = 0 << 0 // Disconnect resistor ladder.
	RpVSS    Config = 1 << 0 // Pull down to VSS (GND).
	RpVDD    Config = 2 << 0 // Pull up to VDD.
	RpVDD1_2 Config = 3 << 0 // Activate both resistors to set input at VDD/2.

	Rn       Config = 3 << 4 // Negative channel pull up/down resistors control
	RnNone   Config = 0 << 4 // Disconnect resistor ladder.
	RnVSS    Config = 1 << 4 // Pull down to VSS (GND).
	RnVDD    Config = 2 << 4 // Pull up to VDD.
	RnVDD1_2 Config = 3 << 4 // Activate both resistors to set input at VDD/2.

	G    Config = 7 << 8 // Gain controll.
	G1_6 Config = 0 << 8 // 1/6
	G1_5 Config = 1 << 8 // 1/5
	G1_4 Config = 2 << 8 // 1/4
	G1_3 Config = 3 << 8 // 1/3
	G1_2 Config = 4 << 8 // 1/2
	G1   Config = 5 << 8 // 1
	G2   Config = 6 << 8 // 2
	G4   Config = 7 << 8 // 4

	RVDD1_4 Config = 1 << 12 // VDD/4 as reference instead of internal 0.6V.

	T     Config = 7 << 16 // Acquistion time.
	T3us  Config = 0 << 16 //  3 us
	T5us  Config = 1 << 16 //  5 us
	T10us Config = 2 << 16 // 10 us
	T15us Config = 3 << 16 // 15 us
	T20us Config = 4 << 16 // 20 us
	T40us Config = 5 << 16 // 40 us

	Diff Config = 1 << 20 // Enable differentail mode.

	Burst Config = 1 << 24 // Enable burst mode.
)

// LoadCONFIG returns the current configuratio of the SAADC channel ch.
func (p *Periph) LoadCONFIG(ch int) Config {
	return Config(p.ch[ch].config.Load())
}

// StoreCONFIG configrues the channel ch.
func (p *Periph) StoreCONFIG(ch int, cfg Config) {
	p.ch[ch].config.Store(uint32(cfg))
}

// LoadLIMIT returns the current limits for channel ch.
func (p *Periph) LoadLIMIT(ch int) (low, high int16) {
	limit := p.ch[ch].limit.Load()
	return int16(limit & 0xFFFF), int16(limit >> 16)
}

// StoreLIMIT sets the limits for channel ch.
func (p *Periph) StoreLIMIT(ch int, low, high int16) {
	limit := uint32(low) | uint32(high)<<16
	p.ch[ch].limit.Store(limit)
}

// LoadRESOLUTION returns the current SAADC resolution: 8, 10, 12 or 14 (bits).
func (p *Periph) LoadRESOLUTION() int {
	return 8 + int(p.resolution.Load())<<1
}

// StoreRESOLUTION sets the SAADC resolution: 8, 10, 12 or 14 (bits).
func (p *Periph) StoreRESOLUTION(res int) {
	p.resolution.Store(uint32(res-8) >> 1)
}

const (
	Over1x   = 0 // Bypass oversampling.
	Over2x   = 1 // Oversample 2x
	Over4x   = 2 // Oversample 4x
	Over8x   = 3 // Oversample 8x
	Over16x  = 4 // Oversample 16x
	Over32x  = 5 // Oversample 32x
	Over64x  = 6 // Oversample 64x
	Over128x = 7 // Oversample 128x
	Over256x = 8 // Oversample 256x
)

// LoadOVERSAMPLE returns the current SAADC oversampling configuration.
func (p *Periph) LoadOVERSAMPLE() int {
	return int(p.oversample.Load())
}

// StoreOVERSAMPLE sets the SAADC oversampling configuration. The oversampling
// should only be used when single input channel is enabled.
func (p *Periph) StoreOVERSAMPLE(over int) {
	p.oversample.Store(uint32(over))
}

// LoadSAMPLERATE returns the "capture and compare" value and reports whether
// the SAADC local timer and cc is used to control SR. If localTimer is true
// the sample rate is equal to 16 MHz / cc, otherwise the SR is controlled by
// SAMPLE task.
func (p *Periph) LoadSAMPLERATE() (cc int, localTimer bool) {
	sr := p.samplerate.Load()
	return int(sr & 0x7ff), sr>>12 != 0
}

// StoreSAMPLERATE configures the source of sample rate. If localTimer is true
// the SAADC local timer is used as sample rate source (SR = 16 MHz / cc).
// Otherwise the SAMPLE task is used. CC must be in the range [80..2047].
// The local timer can only be used when single input channel is enabled.
func (p *Periph) StoreSAMPLERATE(cc int, localTimer bool) {
	p.samplerate.Store(uint32(cc) | internal.BoolToUint32(localTimer)<<12)
}

// LoadPTR returns the last set result data pointer. The PTR register is double
// buffered and you have no access to the internal one.
func (p *Periph) LoadPTR() uintptr {
	return uintptr(p.ptr.Load())
}

// StorePTR sets the result data pointer. The PTR register is double buffered
// and the internal one is unaffected until the START task is triggered.
func (p *Periph) StorePTR(ptr unsafe.Pointer) {
	p.ptr.Store(uint32(uintptr(ptr)))
}

// LoadMAXCNT returns the configured maximum number of 16-bit result samples.
func (p *Periph) LoadMAXCNT() int {
	return int(p.maxcnt.Load())
}

// StoreMAXCNT sets the maximum number of 16-bit samples to be written to the
// output buffer set by StorePTR.
func (p *Periph) StoreMAXCNT(n int) {
	p.maxcnt.Store(uint32(n))
}

// LoadAMOUNT returns the number of 16-bit samples written to the output buffer
// since the previous START task.
func (p *Periph) LoadAMOUNT() int {
	return int(p.amount.Load())
}
