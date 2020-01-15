// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppi

import (
	"github.com/embeddedgo/nrf5/hal/te"
)

// Pre-programmed channels. nrf52-.
const (
	TIMER0_COMPARE0__RADIO_TXEN    te.Chan = 20
	TIMER0_COMPARE0__RADIO_RXEN    te.Chan = 21
	TIMER0_COMPARE1__RADIO_DISABLE te.Chan = 22
	RADIO_BCMATCH__AAR_START       te.Chan = 23
	RADIO_READY__CCM_KSGEN         te.Chan = 24
	RADIO_ADDRESS__CCM_CRYPT       te.Chan = 25
	RADIO_ADDRESS__TIMER0_CAPTURE1 te.Chan = 26
	RADIO_END__TIMER0_CAPTURE2     te.Chan = 27
	RTC0_COMPARE0__RADIO_TXEN      te.Chan = 28
	RTC0_COMPARE0__RADIO_RXEN      te.Chan = 29
	RTC0_COMPARE0__TIMER0_CLEAR    te.Chan = 30
	RTC0_COMPARE0__TIMER0_START    te.Chan = 31
)

// LoadEEP returns the channel Event End Point. nRF52-.
func LoadEEP(c te.Chan) *te.Event {
	return loadEEP(c)
}

// StoreEEP sets the channel Event End Point. nRF52-.
func StoreEEP(c te.Chan, e *te.Event) {
	storeEEP(c, e)
}

// LoadTEP returns the channel Task End Point. nRF52-.
func LoadTEP(c te.Chan) *te.Task {
	return loadTEP(c)
}

// StoreTEP sets the value channel Task End Point. nRF52-.
func StoreTEP(c te.Chan, t *te.Task) {
	storeTEP(c, t)
}

// TEP1 returns the value of channel Fork Task End Point. nRF52.
func TEP1(c te.Chan) *te.Task {
	return loadTEP1(c)
}

// SetTEP1 sets the value of channel Fork Task End Point. nRF52.
func SetTEP1(c te.Chan, t *te.Task) {
	storeTEP1(c, t)
}
