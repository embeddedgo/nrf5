// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package te

import "github.com/embeddedgo/nrf5/hal/internal"

// NumChanGroup is the number of implemented channel groups (4 in case of nRF51,
// 6 in case of nRF52+).
const NumChanGroup = numChanGroup

// ChanGroup repersents PPI channel group.
type ChanGroup int8

var unusedGroups uint32 = 1<<numChanGroup - 1

// AllocChanGroup returns first unused PPI channel group. It returns negative
// number if there is no free channel group.
func AllocChanGroup() ChanGroup {
	return ChanGroup(internal.BitAlloc(&unusedGroups))
}

// Free returns group to the pool of unused channel groups.
func (g ChanGroup) Free() {
	internal.BitFree(&unusedGroups, 1<<g)
}

// Channels returns channels that belongs to the group g.
func (g ChanGroup) Channels() Channels {
	return Channels(ppi().CHG[g].Load())
}

// SetChannels sets channels that belongs to the group g.
func (g ChanGroup) SetChannels(c Channels) {
	ppi().CHG[g].Store(uint32(c))
}

type ChanGroupTask uint8

// EN returns task that can be used to enable channel group g.
func (g ChanGroup) EN() ChanGroupTask {
	return ChanGroupTask(g * 2)
}

// DIS returns task that can be used to disable channel group g.
func (g ChanGroup) DIS() ChanGroupTask {
	return ChanGroupTask(g*2 + 1)
}

func (t ChanGroupTask) Task() *Task { return ppi().Regs.Task(int(t)) }
