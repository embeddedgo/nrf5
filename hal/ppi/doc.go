// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ppi provides access to some features of PPI (Programmable Peripheral
// Interconnect) and DPPI (Distributed PPI). Most of the features of PPI/DPPI
// like channels, channel groups, channel group tasks are handled in unified way
// by te package. The ppi package allows to manage PPI/DPPI interrupts for
// multiple events simultaneously and exposes some PPI specific features that
// are not supported by te package: list of preprogrammed channels, connecting
// one event/task to multiple channels.
package ppi