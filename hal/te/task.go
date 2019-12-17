// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package te

import "embedded/mmio"

// Task represents a peripheral register that is used to trigger events.
type Task struct {
	u32 mmio.U32
}

// Trigger starts action corresponding to task t.
func (r *Task) Trigger() {
	r.u32.Store(1)
}
