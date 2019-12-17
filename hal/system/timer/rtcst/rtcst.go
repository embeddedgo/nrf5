// Package rtcst implements a tickless system timer using the real time counter.
package rtcst

import (
	"embedded/rtos"

	"github.com/embeddedgo/nrf5/hal/rtc"
	"github.com/embeddedgo/nrf5/hal/te"
)

var g struct {
	wakens  int64
	st      *rtc.Periph
	softcnt uint32
	scale   uint32
	ccn     byte
}

func cce() *te.Event {
	return g.st.Event(rtc.COMPARE(int(g.ccn)))
}

// Setup setups st as system timer using ccn compare channel number. Usually
// rtc.RTC1 and channel number 0 is used. Setup accepts st at its reset state or
// configured before for other purposes. For its needs it uses only rtc.OVRFLW
// and rtc.COMPARE(ccn) events and can work with any prescaler set before
// (avoid prescaler values that cause tick period > 1 ms). Setup starts st by
// triggering rtc.START task but accepts RTC started before. It setups st.IRQ()
// priority in NVIC and enables IRQ handling.
func Setup(st *rtc.Periph, ccn int) {
	if uint(ccn) > 3 {
		panic("rtcst: bad ccn")
	}
	g.st = st
	g.ccn = byte(ccn)
	cce := cce()
	cce.DisablePPI()
	cce.DisableIRQ()
	g.scale = (st.LoadPRESCALER() + 1) * (1e9 >> 9)
	ove := st.Event(rtc.OVRFLW)
	ove.DisablePPI()
	ove.EnableIRQ()
	st.IRQ().Enable(rtos.IntPrioSysTimer)
	st.Task(rtc.START).Trigger()
	rtos.SetSystemTimer(nanotime, setAlarm)
}

func counters() (ch, cl uint32) {
	ch = g.softcnt
	for {
		cl = g.st.LoadCOUNTER()
		ch1 := g.softcnt
		if ch1 == ch {
			return
		}
		ch = ch1
	}
}

func nanotime() int64 {
	ch, cl := counters()
	// exact and efficient calculation of: (ch<<24+cl)*1e9*(prescaler+1)/32768.
	scale := int64(g.scale)
	h := int64(ch) * scale << (24 - (15 - 9))
	l := int64(cl) * scale >> (15 - 9)
	return h + l
}

func ticks(ch, cl uint32) int64 {
	return int64(ch)<<24 | int64(cl)
}

func setAlarm(ns int64) bool {
	if g.wakens == ns {
		return true
	}
	g.wakens = ns
	scale := int64(g.scale)
	wkup := (ns<<(15-9) + scale - 1) / scale // round up, don't wake too early
	cce := cce()
	cce.Clear()
	ch, cl := counters()
	now := ticks(ch, cl)
	sleep := wkup - now
	switch {
	case sleep > 0xffffff:
		g.st.StoreCC(int(g.ccn), cl)
	case sleep > 0:
		g.st.StoreCC(int(g.ccn), uint32(wkup)&0xffffff)
		now = ticks(counters())
	}
	if now < wkup {
		cce.EnableIRQ()
		return true
	}
	// wkup in the past or there is a chance that CC was set to late.
	g.wakens = 0
	return false
}

func ISR() {
	ove := g.st.Event(rtc.OVRFLW)
	cce := cce()
	if ove.IsSet() {
		ove.Clear()
		g.softcnt++
	}
	if cce.IsSet() {
		cce.DisableIRQ()
		g.wakens = 0
		schedule()
	}
}

func schedule()
