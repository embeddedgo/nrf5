// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppi

import (
	"github.com/embeddedgo/nrf5/hal/internal"
	"github.com/embeddedgo/nrf5/hal/te"
)

// GroupNum is the number of implemented channel groups (4 in case of nRF51, 6
// in case of nRF52).
const GroupNum = groupNum

// Group repersents PPI channel group.
type Group int8

var unusedGroups uint32 = 1<<groupNum - 1

// GroupAlloc returns first unused PPI channel group. It returns negative number
// if there is no free channel group.
func GroupAlloc() Group {
	return Group(internal.BitAlloc(&unusedGroups))
}

// Free returns group to the pool of unused channel groups.
func (g Group) Free() {
	internal.BitFree(&unusedGroups, 1<<g)
}

// Channels returns channels that belongs to the group g.
func (g Group) Channels() Channels {
	return Channels(r().chg[g].Load())
}

// SetChannels sets channels that belongs to the group g.
func (g Group) SetChannels(c Channels) {
	r().chg[g].Store(uint32(c))
}

type Task byte

// EN returns task that can be used to enable channel group g.
func (g Group) EN() Task {
	return Task(g * 2)
}

// DIS returns task that can be used to disable channel group g.
func (g Group) DIS() Task {
	return Task(g*2 + 1)
}

func (t Task) Task() *te.Task { return r().Regs.Task(int(t)) }
