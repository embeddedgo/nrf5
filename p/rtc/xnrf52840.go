// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build nrf52840

package rtc

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/nrf5/p/mmap"
)

type Periph struct {
	TASK_START      RTASK_START
	TASK_STOP       RTASK_STOP
	TASK_CLEAR      RTASK_CLEAR
	TASK_TRIGOVRFLW RTASK_TRIGOVRFLW
	_               [60]uint32
	EVENT_TICK      REVENT_TICK
	EVENT_OVRFLW    REVENT_OVRFLW
	_               [14]uint32
	EVENT_COMPARE   [4]REVENT_COMPARE
	_               [109]uint32
	INTENSET        RINTENSET
	INTENCLR        RINTENCLR
	_               [13]uint32
	EVTEN           REVTEN
	EVTENSET        REVTENSET
	EVTENCLR        REVTENCLR
	_               [110]uint32
	COUNTER         RCOUNTER
	PRESCALER       RPRESCALER
	_               [13]uint32
	CC              [4]RCC
}

func RTC0() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.RTC0_BASE))) }
func RTC1() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.RTC1_BASE))) }
func RTC2() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.RTC2_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type TASK_START uint32

type RTASK_START struct{ mmio.U32 }

func (r *RTASK_START) LoadBits(mask TASK_START) TASK_START {
	return TASK_START(r.U32.LoadBits(uint32(mask)))
}
func (r *RTASK_START) StoreBits(mask, b TASK_START) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTASK_START) SetBits(mask TASK_START)      { r.U32.SetBits(uint32(mask)) }
func (r *RTASK_START) ClearBits(mask TASK_START)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTASK_START) Load() TASK_START             { return TASK_START(r.U32.Load()) }
func (r *RTASK_START) Store(b TASK_START)           { r.U32.Store(uint32(b)) }

type RMTASK_START struct{ mmio.UM32 }

func (rm RMTASK_START) Load() TASK_START   { return TASK_START(rm.UM32.Load()) }
func (rm RMTASK_START) Store(b TASK_START) { rm.UM32.Store(uint32(b)) }

type TASK_STOP uint32

type RTASK_STOP struct{ mmio.U32 }

func (r *RTASK_STOP) LoadBits(mask TASK_STOP) TASK_STOP {
	return TASK_STOP(r.U32.LoadBits(uint32(mask)))
}
func (r *RTASK_STOP) StoreBits(mask, b TASK_STOP) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTASK_STOP) SetBits(mask TASK_STOP)      { r.U32.SetBits(uint32(mask)) }
func (r *RTASK_STOP) ClearBits(mask TASK_STOP)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTASK_STOP) Load() TASK_STOP             { return TASK_STOP(r.U32.Load()) }
func (r *RTASK_STOP) Store(b TASK_STOP)           { r.U32.Store(uint32(b)) }

type RMTASK_STOP struct{ mmio.UM32 }

func (rm RMTASK_STOP) Load() TASK_STOP   { return TASK_STOP(rm.UM32.Load()) }
func (rm RMTASK_STOP) Store(b TASK_STOP) { rm.UM32.Store(uint32(b)) }

type TASK_CLEAR uint32

type RTASK_CLEAR struct{ mmio.U32 }

func (r *RTASK_CLEAR) LoadBits(mask TASK_CLEAR) TASK_CLEAR {
	return TASK_CLEAR(r.U32.LoadBits(uint32(mask)))
}
func (r *RTASK_CLEAR) StoreBits(mask, b TASK_CLEAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTASK_CLEAR) SetBits(mask TASK_CLEAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTASK_CLEAR) ClearBits(mask TASK_CLEAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTASK_CLEAR) Load() TASK_CLEAR             { return TASK_CLEAR(r.U32.Load()) }
func (r *RTASK_CLEAR) Store(b TASK_CLEAR)           { r.U32.Store(uint32(b)) }

type RMTASK_CLEAR struct{ mmio.UM32 }

func (rm RMTASK_CLEAR) Load() TASK_CLEAR   { return TASK_CLEAR(rm.UM32.Load()) }
func (rm RMTASK_CLEAR) Store(b TASK_CLEAR) { rm.UM32.Store(uint32(b)) }

type TASK_TRIGOVRFLW uint32

type RTASK_TRIGOVRFLW struct{ mmio.U32 }

func (r *RTASK_TRIGOVRFLW) LoadBits(mask TASK_TRIGOVRFLW) TASK_TRIGOVRFLW {
	return TASK_TRIGOVRFLW(r.U32.LoadBits(uint32(mask)))
}
func (r *RTASK_TRIGOVRFLW) StoreBits(mask, b TASK_TRIGOVRFLW) {
	r.U32.StoreBits(uint32(mask), uint32(b))
}
func (r *RTASK_TRIGOVRFLW) SetBits(mask TASK_TRIGOVRFLW)   { r.U32.SetBits(uint32(mask)) }
func (r *RTASK_TRIGOVRFLW) ClearBits(mask TASK_TRIGOVRFLW) { r.U32.ClearBits(uint32(mask)) }
func (r *RTASK_TRIGOVRFLW) Load() TASK_TRIGOVRFLW          { return TASK_TRIGOVRFLW(r.U32.Load()) }
func (r *RTASK_TRIGOVRFLW) Store(b TASK_TRIGOVRFLW)        { r.U32.Store(uint32(b)) }

type RMTASK_TRIGOVRFLW struct{ mmio.UM32 }

func (rm RMTASK_TRIGOVRFLW) Load() TASK_TRIGOVRFLW   { return TASK_TRIGOVRFLW(rm.UM32.Load()) }
func (rm RMTASK_TRIGOVRFLW) Store(b TASK_TRIGOVRFLW) { rm.UM32.Store(uint32(b)) }

type EVENT_TICK uint32

type REVENT_TICK struct{ mmio.U32 }

func (r *REVENT_TICK) LoadBits(mask EVENT_TICK) EVENT_TICK {
	return EVENT_TICK(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_TICK) StoreBits(mask, b EVENT_TICK) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_TICK) SetBits(mask EVENT_TICK)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_TICK) ClearBits(mask EVENT_TICK)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_TICK) Load() EVENT_TICK             { return EVENT_TICK(r.U32.Load()) }
func (r *REVENT_TICK) Store(b EVENT_TICK)           { r.U32.Store(uint32(b)) }

type RMEVENT_TICK struct{ mmio.UM32 }

func (rm RMEVENT_TICK) Load() EVENT_TICK   { return EVENT_TICK(rm.UM32.Load()) }
func (rm RMEVENT_TICK) Store(b EVENT_TICK) { rm.UM32.Store(uint32(b)) }

type EVENT_OVRFLW uint32

type REVENT_OVRFLW struct{ mmio.U32 }

func (r *REVENT_OVRFLW) LoadBits(mask EVENT_OVRFLW) EVENT_OVRFLW {
	return EVENT_OVRFLW(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_OVRFLW) StoreBits(mask, b EVENT_OVRFLW) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_OVRFLW) SetBits(mask EVENT_OVRFLW)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_OVRFLW) ClearBits(mask EVENT_OVRFLW)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_OVRFLW) Load() EVENT_OVRFLW             { return EVENT_OVRFLW(r.U32.Load()) }
func (r *REVENT_OVRFLW) Store(b EVENT_OVRFLW)           { r.U32.Store(uint32(b)) }

type RMEVENT_OVRFLW struct{ mmio.UM32 }

func (rm RMEVENT_OVRFLW) Load() EVENT_OVRFLW   { return EVENT_OVRFLW(rm.UM32.Load()) }
func (rm RMEVENT_OVRFLW) Store(b EVENT_OVRFLW) { rm.UM32.Store(uint32(b)) }

type EVENT_COMPARE uint32

type REVENT_COMPARE struct{ mmio.U32 }

func (r *REVENT_COMPARE) LoadBits(mask EVENT_COMPARE) EVENT_COMPARE {
	return EVENT_COMPARE(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_COMPARE) StoreBits(mask, b EVENT_COMPARE) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_COMPARE) SetBits(mask EVENT_COMPARE)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_COMPARE) ClearBits(mask EVENT_COMPARE)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_COMPARE) Load() EVENT_COMPARE             { return EVENT_COMPARE(r.U32.Load()) }
func (r *REVENT_COMPARE) Store(b EVENT_COMPARE)           { r.U32.Store(uint32(b)) }

type RMEVENT_COMPARE struct{ mmio.UM32 }

func (rm RMEVENT_COMPARE) Load() EVENT_COMPARE   { return EVENT_COMPARE(rm.UM32.Load()) }
func (rm RMEVENT_COMPARE) Store(b EVENT_COMPARE) { rm.UM32.Store(uint32(b)) }

type INTENSET uint32

type RINTENSET struct{ mmio.U32 }

func (r *RINTENSET) LoadBits(mask INTENSET) INTENSET { return INTENSET(r.U32.LoadBits(uint32(mask))) }
func (r *RINTENSET) StoreBits(mask, b INTENSET)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RINTENSET) SetBits(mask INTENSET)           { r.U32.SetBits(uint32(mask)) }
func (r *RINTENSET) ClearBits(mask INTENSET)         { r.U32.ClearBits(uint32(mask)) }
func (r *RINTENSET) Load() INTENSET                  { return INTENSET(r.U32.Load()) }
func (r *RINTENSET) Store(b INTENSET)                { r.U32.Store(uint32(b)) }

type RMINTENSET struct{ mmio.UM32 }

func (rm RMINTENSET) Load() INTENSET   { return INTENSET(rm.UM32.Load()) }
func (rm RMINTENSET) Store(b INTENSET) { rm.UM32.Store(uint32(b)) }

type INTENCLR uint32

type RINTENCLR struct{ mmio.U32 }

func (r *RINTENCLR) LoadBits(mask INTENCLR) INTENCLR { return INTENCLR(r.U32.LoadBits(uint32(mask))) }
func (r *RINTENCLR) StoreBits(mask, b INTENCLR)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RINTENCLR) SetBits(mask INTENCLR)           { r.U32.SetBits(uint32(mask)) }
func (r *RINTENCLR) ClearBits(mask INTENCLR)         { r.U32.ClearBits(uint32(mask)) }
func (r *RINTENCLR) Load() INTENCLR                  { return INTENCLR(r.U32.Load()) }
func (r *RINTENCLR) Store(b INTENCLR)                { r.U32.Store(uint32(b)) }

type RMINTENCLR struct{ mmio.UM32 }

func (rm RMINTENCLR) Load() INTENCLR   { return INTENCLR(rm.UM32.Load()) }
func (rm RMINTENCLR) Store(b INTENCLR) { rm.UM32.Store(uint32(b)) }

type EVTEN uint32

type REVTEN struct{ mmio.U32 }

func (r *REVTEN) LoadBits(mask EVTEN) EVTEN { return EVTEN(r.U32.LoadBits(uint32(mask))) }
func (r *REVTEN) StoreBits(mask, b EVTEN)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVTEN) SetBits(mask EVTEN)        { r.U32.SetBits(uint32(mask)) }
func (r *REVTEN) ClearBits(mask EVTEN)      { r.U32.ClearBits(uint32(mask)) }
func (r *REVTEN) Load() EVTEN               { return EVTEN(r.U32.Load()) }
func (r *REVTEN) Store(b EVTEN)             { r.U32.Store(uint32(b)) }

type RMEVTEN struct{ mmio.UM32 }

func (rm RMEVTEN) Load() EVTEN   { return EVTEN(rm.UM32.Load()) }
func (rm RMEVTEN) Store(b EVTEN) { rm.UM32.Store(uint32(b)) }

type EVTENSET uint32

type REVTENSET struct{ mmio.U32 }

func (r *REVTENSET) LoadBits(mask EVTENSET) EVTENSET { return EVTENSET(r.U32.LoadBits(uint32(mask))) }
func (r *REVTENSET) StoreBits(mask, b EVTENSET)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVTENSET) SetBits(mask EVTENSET)           { r.U32.SetBits(uint32(mask)) }
func (r *REVTENSET) ClearBits(mask EVTENSET)         { r.U32.ClearBits(uint32(mask)) }
func (r *REVTENSET) Load() EVTENSET                  { return EVTENSET(r.U32.Load()) }
func (r *REVTENSET) Store(b EVTENSET)                { r.U32.Store(uint32(b)) }

type RMEVTENSET struct{ mmio.UM32 }

func (rm RMEVTENSET) Load() EVTENSET   { return EVTENSET(rm.UM32.Load()) }
func (rm RMEVTENSET) Store(b EVTENSET) { rm.UM32.Store(uint32(b)) }

type EVTENCLR uint32

type REVTENCLR struct{ mmio.U32 }

func (r *REVTENCLR) LoadBits(mask EVTENCLR) EVTENCLR { return EVTENCLR(r.U32.LoadBits(uint32(mask))) }
func (r *REVTENCLR) StoreBits(mask, b EVTENCLR)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVTENCLR) SetBits(mask EVTENCLR)           { r.U32.SetBits(uint32(mask)) }
func (r *REVTENCLR) ClearBits(mask EVTENCLR)         { r.U32.ClearBits(uint32(mask)) }
func (r *REVTENCLR) Load() EVTENCLR                  { return EVTENCLR(r.U32.Load()) }
func (r *REVTENCLR) Store(b EVTENCLR)                { r.U32.Store(uint32(b)) }

type RMEVTENCLR struct{ mmio.UM32 }

func (rm RMEVTENCLR) Load() EVTENCLR   { return EVTENCLR(rm.UM32.Load()) }
func (rm RMEVTENCLR) Store(b EVTENCLR) { rm.UM32.Store(uint32(b)) }

type COUNTER uint32

type RCOUNTER struct{ mmio.U32 }

func (r *RCOUNTER) LoadBits(mask COUNTER) COUNTER { return COUNTER(r.U32.LoadBits(uint32(mask))) }
func (r *RCOUNTER) StoreBits(mask, b COUNTER)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCOUNTER) SetBits(mask COUNTER)          { r.U32.SetBits(uint32(mask)) }
func (r *RCOUNTER) ClearBits(mask COUNTER)        { r.U32.ClearBits(uint32(mask)) }
func (r *RCOUNTER) Load() COUNTER                 { return COUNTER(r.U32.Load()) }
func (r *RCOUNTER) Store(b COUNTER)               { r.U32.Store(uint32(b)) }

type RMCOUNTER struct{ mmio.UM32 }

func (rm RMCOUNTER) Load() COUNTER   { return COUNTER(rm.UM32.Load()) }
func (rm RMCOUNTER) Store(b COUNTER) { rm.UM32.Store(uint32(b)) }

type PRESCALER uint32

type RPRESCALER struct{ mmio.U32 }

func (r *RPRESCALER) LoadBits(mask PRESCALER) PRESCALER {
	return PRESCALER(r.U32.LoadBits(uint32(mask)))
}
func (r *RPRESCALER) StoreBits(mask, b PRESCALER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPRESCALER) SetBits(mask PRESCALER)      { r.U32.SetBits(uint32(mask)) }
func (r *RPRESCALER) ClearBits(mask PRESCALER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPRESCALER) Load() PRESCALER             { return PRESCALER(r.U32.Load()) }
func (r *RPRESCALER) Store(b PRESCALER)           { r.U32.Store(uint32(b)) }

type RMPRESCALER struct{ mmio.UM32 }

func (rm RMPRESCALER) Load() PRESCALER   { return PRESCALER(rm.UM32.Load()) }
func (rm RMPRESCALER) Store(b PRESCALER) { rm.UM32.Store(uint32(b)) }

type CC uint32

type RCC struct{ mmio.U32 }

func (r *RCC) LoadBits(mask CC) CC  { return CC(r.U32.LoadBits(uint32(mask))) }
func (r *RCC) StoreBits(mask, b CC) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCC) SetBits(mask CC)      { r.U32.SetBits(uint32(mask)) }
func (r *RCC) ClearBits(mask CC)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCC) Load() CC             { return CC(r.U32.Load()) }
func (r *RCC) Store(b CC)           { r.U32.Store(uint32(b)) }

type RMCC struct{ mmio.UM32 }

func (rm RMCC) Load() CC   { return CC(rm.UM32.Load()) }
func (rm RMCC) Store(b CC) { rm.UM32.Store(uint32(b)) }
