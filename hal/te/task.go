// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package te

import "embedded/mmio"

// Task represents a peripheral register that is used to trigger events.
type Task struct {
	u32 mmio.U32
}

// Trigger starts action corresponding to task.
func (r *Task) Trigger() {
	r.u32.Store(1)
}

// Chan returns the PPI/DPPI channel the task is connected to and reports
// whether subscription of events on this channel is enabled. In case of nRF52-
// which PPI peripheral allows to connect task to multiple PPI channels Chan
// returns the first channel found.
func (r *Task) Chan() (c Chan, en bool) {
	return getTaskChan(r)
}

// SetChan connects the task to the PPI/DPPI channel. En allows to enable or
// disable subscription of events on connected channel.
//
// SetChan allows to write portable code that can work with PPI (nRF52-) and
// DPPI (nRF53) but has some limitations in case of PPI:
//
// - you can connect only up to two tasks to the same channel,
//
// - if en is false the channel is set to -1 which as a result disconnects the
//   task from all channels.
//
// The one advantage of PPI over DPPI is ablility to conect one task to multiple
// channels. If you need this feature use SetTEP from ppi package.
func (r *Task) SetChan(c Chan, en bool) {
	setTaskChan(r, c, en)
}
