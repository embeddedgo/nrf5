// Copyright 2019 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example forces the garbage collector to work hard and periodically
// prints statistics of the memory allocator.
package main

import (
	"math/rand"
	"runtime"

	"github.com/embeddedgo/nrf5/hal/clock"

	_ "github.com/embeddedgo/nrf5/devboard/pca10059/board/system"
)

func main() {
	clock.StoreTRACECONFIG(clock.T4MHz, clock.Serial) // enable SWO on P1.00

	runtime.GOMAXPROCS(2)

	c := make(chan []uint32, 3)

	go func() {
		for {
			c <- make([]uint32, rand.Intn(512)+1)
		}
	}()

	for i := 0; ; i++ {
		<-c
		if i&1023 == 0 {
			if i&32767 == 0 {
				printMemstatDescr()
			}
			printMemstat(i)
		}
	}
}

func printMemstatDescr() {
	print("#Recv  Sys HeapSys StackSys GCSys OtherSys  " +
		"HeapAlloc/HeapInuse StackInuse HeapIdle  " +
		"MSpanInuse/MSpanSys MCacheInuse/MCacheSys\n")
}

var ms runtime.MemStats

func printMemstat(i int) {
	runtime.ReadMemStats(&ms)
	print(
		i>>10, "K\t",
		ms.Sys, "\t", ms.HeapSys, "\t", ms.StackSys, "\t", ms.GCSys, "\t",
		ms.OtherSys, "\t\t",
		ms.HeapAlloc, " / ", ms.HeapInuse, "\t", ms.StackInuse, "\t",
		ms.HeapIdle, "\t\t",
		ms.MSpanInuse, " / ", ms.MSpanSys, "\t", ms.MCacheInuse, " / ",
		ms.MCacheSys,
		"\n",
	)
}
