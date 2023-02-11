// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package saadc

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/embeddedgo/nrf5/hal/gpio"
)

// TODO: we need []sync.Mutex here in case of multiple SAADC peripherals.
var analogMutex sync.Mutex

// AnalogInit provides a convenient way to configure the SAADC peripheral to use
// it by AnalogRead function. It configures and calibrates SAADC(0),
// disconnnects all irs channels from analog inputs and changes the
// configuration of channel 0.
//
// The SAADC resolution is set to resBits. The 4x hardware oversampling in burst
// mode is used to reduce the noise. The channel 0 configuration is set to
// G1_4|RVDD1_4|T10us|Burst. You can modify these settings using SAADC(0)
// StoreOVERSAMPLE and StoreCONFIG methods.
func AnalogInit(resBits int) {
	p := SAADC(0)
	p.IRQ().Disable(0)
	p.StoreENABLE(false)
	p.StoreRESOLUTION(resBits)
	p.StoreOVERSAMPLE(Over4x)
	p.StoreSAMPLERATE(0, false)
	p.StoreMAXCNT(1)
	for ch := 0; ch < 8; ch++ {
		p.StorePSELP(ch, gpio.ANC)
		p.StorePSELN(ch, gpio.ANC)
	}
	p.StoreCONFIG(0, G1_4|RVDD1_4|T10us|Burst)
	p.StoreENABLE(true)

	p.Event(CALIBRATEDONE).Clear()
	p.Task(CALIBRATE).Trig()
	for !p.Event(CALIBRATEDONE).IsSet() {
		runtime.Gosched()
	}
}

// AnalogRead works like Arduino analogRead function. It connects SAADC channel
// 0 to ain and performs a conversion triggering START and SAMPLE tasks. If the
// 4x oversampling is used (AnalogInit default) the conversion takes about 100
// microsecond just like the original analogRead on Arduino UNO, Nano, Mini,
// Mega. Use SAADC directly in continuous mode to achieve the maximum 200 kHz
// sampling rate (5 microseconds per sample).
//
// The SAADC peripheral and its channel 0 must be properly configured before
// using this function (see AnalogInit).
func AnalogRead(ain gpio.AIN) int {
	analogMutex.Lock()
	var a int16
	p := SAADC(0)
	p.StorePSELP(0, ain)
	p.StorePTR(unsafe.Pointer(&a))
	p.Event(STARTED).Clear()
	p.Task(START).Trig()
	p.Event(END).Clear() // clear END here to give more time for STARTED event
	for !p.Event(STARTED).IsSet() {
		// In practice the CPU never enters here
	}
	p.Task(SAMPLE).Trig()
	for !p.Event(END).IsSet() {
		// With Gosched this loop is executed 1-2 times instead of 181 times
		// (10bit, Over4x).
		runtime.Gosched()
	}
	analogMutex.Unlock()
	return int(a)
}
