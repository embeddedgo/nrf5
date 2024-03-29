// DO NOT EDIT THIS FILE. GENERATED BY xgen.

//go:build nrf52840

package uarte

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/nrf5/p/mmap"
)

type Periph struct {
	TASK_STARTRX    RTASK_STARTRX
	TASK_STOPRX     RTASK_STOPRX
	TASK_STARTTX    RTASK_STARTTX
	TASK_STOPTX     RTASK_STOPTX
	_               [7]uint32
	TASK_FLUSHRX    RTASK_FLUSHRX
	_               [52]uint32
	EVENT_CTS       REVENT_CTS
	EVENT_NCTS      REVENT_NCTS
	EVENT_RXDRDY    REVENT_RXDRDY
	_               uint32
	EVENT_ENDRX     REVENT_ENDRX
	_               [2]uint32
	EVENT_TXDRDY    REVENT_TXDRDY
	EVENT_ENDTX     REVENT_ENDTX
	EVENT_ERROR     REVENT_ERROR
	_               [7]uint32
	EVENT_RXTO      REVENT_RXTO
	_               uint32
	EVENT_RXSTARTED REVENT_RXSTARTED
	EVENT_TXSTARTED REVENT_TXSTARTED
	_               uint32
	EVENT_TXSTOPPED REVENT_TXSTOPPED
	_               [41]uint32
	SHORTS          RSHORTS
	_               [63]uint32
	INTEN           RINTEN
	INTENSET        RINTENSET
	INTENCLR        RINTENCLR
	_               [93]uint32
	ERRORSRC        RERRORSRC
	_               [31]uint32
	ENABLE          RENABLE
	_               uint32
	PSEL_RTS        RPSEL_RTS
	PSEL_TXD        RPSEL_TXD
	PSEL_CTS        RPSEL_CTS
	PSEL_RXD        RPSEL_RXD
	_               [3]uint32
	BAUDRATE        RBAUDRATE
	_               [3]uint32
	RXD_PTR         RRXD_PTR
	RXD_MAXCNT      RRXD_MAXCNT
	RXD_AMOUNT      RRXD_AMOUNT
	_               uint32
	TXD_PTR         RTXD_PTR
	TXD_MAXCNT      RTXD_MAXCNT
	TXD_AMOUNT      RTXD_AMOUNT
	_               [7]uint32
	CONFIG          RCONFIG
}

func UARTE0() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.UARTE0_BASE))) }
func UARTE1() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.UARTE1_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type TASK_STARTRX uint32

type RTASK_STARTRX struct{ mmio.U32 }

func (r *RTASK_STARTRX) LoadBits(mask TASK_STARTRX) TASK_STARTRX {
	return TASK_STARTRX(r.U32.LoadBits(uint32(mask)))
}
func (r *RTASK_STARTRX) StoreBits(mask, b TASK_STARTRX) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTASK_STARTRX) SetBits(mask TASK_STARTRX)      { r.U32.SetBits(uint32(mask)) }
func (r *RTASK_STARTRX) ClearBits(mask TASK_STARTRX)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTASK_STARTRX) Load() TASK_STARTRX             { return TASK_STARTRX(r.U32.Load()) }
func (r *RTASK_STARTRX) Store(b TASK_STARTRX)           { r.U32.Store(uint32(b)) }

type RMTASK_STARTRX struct{ mmio.UM32 }

func (rm RMTASK_STARTRX) Load() TASK_STARTRX   { return TASK_STARTRX(rm.UM32.Load()) }
func (rm RMTASK_STARTRX) Store(b TASK_STARTRX) { rm.UM32.Store(uint32(b)) }

type TASK_STOPRX uint32

type RTASK_STOPRX struct{ mmio.U32 }

func (r *RTASK_STOPRX) LoadBits(mask TASK_STOPRX) TASK_STOPRX {
	return TASK_STOPRX(r.U32.LoadBits(uint32(mask)))
}
func (r *RTASK_STOPRX) StoreBits(mask, b TASK_STOPRX) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTASK_STOPRX) SetBits(mask TASK_STOPRX)      { r.U32.SetBits(uint32(mask)) }
func (r *RTASK_STOPRX) ClearBits(mask TASK_STOPRX)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTASK_STOPRX) Load() TASK_STOPRX             { return TASK_STOPRX(r.U32.Load()) }
func (r *RTASK_STOPRX) Store(b TASK_STOPRX)           { r.U32.Store(uint32(b)) }

type RMTASK_STOPRX struct{ mmio.UM32 }

func (rm RMTASK_STOPRX) Load() TASK_STOPRX   { return TASK_STOPRX(rm.UM32.Load()) }
func (rm RMTASK_STOPRX) Store(b TASK_STOPRX) { rm.UM32.Store(uint32(b)) }

type TASK_STARTTX uint32

type RTASK_STARTTX struct{ mmio.U32 }

func (r *RTASK_STARTTX) LoadBits(mask TASK_STARTTX) TASK_STARTTX {
	return TASK_STARTTX(r.U32.LoadBits(uint32(mask)))
}
func (r *RTASK_STARTTX) StoreBits(mask, b TASK_STARTTX) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTASK_STARTTX) SetBits(mask TASK_STARTTX)      { r.U32.SetBits(uint32(mask)) }
func (r *RTASK_STARTTX) ClearBits(mask TASK_STARTTX)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTASK_STARTTX) Load() TASK_STARTTX             { return TASK_STARTTX(r.U32.Load()) }
func (r *RTASK_STARTTX) Store(b TASK_STARTTX)           { r.U32.Store(uint32(b)) }

type RMTASK_STARTTX struct{ mmio.UM32 }

func (rm RMTASK_STARTTX) Load() TASK_STARTTX   { return TASK_STARTTX(rm.UM32.Load()) }
func (rm RMTASK_STARTTX) Store(b TASK_STARTTX) { rm.UM32.Store(uint32(b)) }

type TASK_STOPTX uint32

type RTASK_STOPTX struct{ mmio.U32 }

func (r *RTASK_STOPTX) LoadBits(mask TASK_STOPTX) TASK_STOPTX {
	return TASK_STOPTX(r.U32.LoadBits(uint32(mask)))
}
func (r *RTASK_STOPTX) StoreBits(mask, b TASK_STOPTX) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTASK_STOPTX) SetBits(mask TASK_STOPTX)      { r.U32.SetBits(uint32(mask)) }
func (r *RTASK_STOPTX) ClearBits(mask TASK_STOPTX)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTASK_STOPTX) Load() TASK_STOPTX             { return TASK_STOPTX(r.U32.Load()) }
func (r *RTASK_STOPTX) Store(b TASK_STOPTX)           { r.U32.Store(uint32(b)) }

type RMTASK_STOPTX struct{ mmio.UM32 }

func (rm RMTASK_STOPTX) Load() TASK_STOPTX   { return TASK_STOPTX(rm.UM32.Load()) }
func (rm RMTASK_STOPTX) Store(b TASK_STOPTX) { rm.UM32.Store(uint32(b)) }

type TASK_FLUSHRX uint32

type RTASK_FLUSHRX struct{ mmio.U32 }

func (r *RTASK_FLUSHRX) LoadBits(mask TASK_FLUSHRX) TASK_FLUSHRX {
	return TASK_FLUSHRX(r.U32.LoadBits(uint32(mask)))
}
func (r *RTASK_FLUSHRX) StoreBits(mask, b TASK_FLUSHRX) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTASK_FLUSHRX) SetBits(mask TASK_FLUSHRX)      { r.U32.SetBits(uint32(mask)) }
func (r *RTASK_FLUSHRX) ClearBits(mask TASK_FLUSHRX)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTASK_FLUSHRX) Load() TASK_FLUSHRX             { return TASK_FLUSHRX(r.U32.Load()) }
func (r *RTASK_FLUSHRX) Store(b TASK_FLUSHRX)           { r.U32.Store(uint32(b)) }

type RMTASK_FLUSHRX struct{ mmio.UM32 }

func (rm RMTASK_FLUSHRX) Load() TASK_FLUSHRX   { return TASK_FLUSHRX(rm.UM32.Load()) }
func (rm RMTASK_FLUSHRX) Store(b TASK_FLUSHRX) { rm.UM32.Store(uint32(b)) }

type EVENT_CTS uint32

type REVENT_CTS struct{ mmio.U32 }

func (r *REVENT_CTS) LoadBits(mask EVENT_CTS) EVENT_CTS {
	return EVENT_CTS(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_CTS) StoreBits(mask, b EVENT_CTS) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_CTS) SetBits(mask EVENT_CTS)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_CTS) ClearBits(mask EVENT_CTS)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_CTS) Load() EVENT_CTS             { return EVENT_CTS(r.U32.Load()) }
func (r *REVENT_CTS) Store(b EVENT_CTS)           { r.U32.Store(uint32(b)) }

type RMEVENT_CTS struct{ mmio.UM32 }

func (rm RMEVENT_CTS) Load() EVENT_CTS   { return EVENT_CTS(rm.UM32.Load()) }
func (rm RMEVENT_CTS) Store(b EVENT_CTS) { rm.UM32.Store(uint32(b)) }

type EVENT_NCTS uint32

type REVENT_NCTS struct{ mmio.U32 }

func (r *REVENT_NCTS) LoadBits(mask EVENT_NCTS) EVENT_NCTS {
	return EVENT_NCTS(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_NCTS) StoreBits(mask, b EVENT_NCTS) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_NCTS) SetBits(mask EVENT_NCTS)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_NCTS) ClearBits(mask EVENT_NCTS)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_NCTS) Load() EVENT_NCTS             { return EVENT_NCTS(r.U32.Load()) }
func (r *REVENT_NCTS) Store(b EVENT_NCTS)           { r.U32.Store(uint32(b)) }

type RMEVENT_NCTS struct{ mmio.UM32 }

func (rm RMEVENT_NCTS) Load() EVENT_NCTS   { return EVENT_NCTS(rm.UM32.Load()) }
func (rm RMEVENT_NCTS) Store(b EVENT_NCTS) { rm.UM32.Store(uint32(b)) }

type EVENT_RXDRDY uint32

type REVENT_RXDRDY struct{ mmio.U32 }

func (r *REVENT_RXDRDY) LoadBits(mask EVENT_RXDRDY) EVENT_RXDRDY {
	return EVENT_RXDRDY(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_RXDRDY) StoreBits(mask, b EVENT_RXDRDY) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_RXDRDY) SetBits(mask EVENT_RXDRDY)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_RXDRDY) ClearBits(mask EVENT_RXDRDY)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_RXDRDY) Load() EVENT_RXDRDY             { return EVENT_RXDRDY(r.U32.Load()) }
func (r *REVENT_RXDRDY) Store(b EVENT_RXDRDY)           { r.U32.Store(uint32(b)) }

type RMEVENT_RXDRDY struct{ mmio.UM32 }

func (rm RMEVENT_RXDRDY) Load() EVENT_RXDRDY   { return EVENT_RXDRDY(rm.UM32.Load()) }
func (rm RMEVENT_RXDRDY) Store(b EVENT_RXDRDY) { rm.UM32.Store(uint32(b)) }

type EVENT_ENDRX uint32

type REVENT_ENDRX struct{ mmio.U32 }

func (r *REVENT_ENDRX) LoadBits(mask EVENT_ENDRX) EVENT_ENDRX {
	return EVENT_ENDRX(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_ENDRX) StoreBits(mask, b EVENT_ENDRX) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_ENDRX) SetBits(mask EVENT_ENDRX)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_ENDRX) ClearBits(mask EVENT_ENDRX)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_ENDRX) Load() EVENT_ENDRX             { return EVENT_ENDRX(r.U32.Load()) }
func (r *REVENT_ENDRX) Store(b EVENT_ENDRX)           { r.U32.Store(uint32(b)) }

type RMEVENT_ENDRX struct{ mmio.UM32 }

func (rm RMEVENT_ENDRX) Load() EVENT_ENDRX   { return EVENT_ENDRX(rm.UM32.Load()) }
func (rm RMEVENT_ENDRX) Store(b EVENT_ENDRX) { rm.UM32.Store(uint32(b)) }

type EVENT_TXDRDY uint32

type REVENT_TXDRDY struct{ mmio.U32 }

func (r *REVENT_TXDRDY) LoadBits(mask EVENT_TXDRDY) EVENT_TXDRDY {
	return EVENT_TXDRDY(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_TXDRDY) StoreBits(mask, b EVENT_TXDRDY) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_TXDRDY) SetBits(mask EVENT_TXDRDY)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_TXDRDY) ClearBits(mask EVENT_TXDRDY)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_TXDRDY) Load() EVENT_TXDRDY             { return EVENT_TXDRDY(r.U32.Load()) }
func (r *REVENT_TXDRDY) Store(b EVENT_TXDRDY)           { r.U32.Store(uint32(b)) }

type RMEVENT_TXDRDY struct{ mmio.UM32 }

func (rm RMEVENT_TXDRDY) Load() EVENT_TXDRDY   { return EVENT_TXDRDY(rm.UM32.Load()) }
func (rm RMEVENT_TXDRDY) Store(b EVENT_TXDRDY) { rm.UM32.Store(uint32(b)) }

type EVENT_ENDTX uint32

type REVENT_ENDTX struct{ mmio.U32 }

func (r *REVENT_ENDTX) LoadBits(mask EVENT_ENDTX) EVENT_ENDTX {
	return EVENT_ENDTX(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_ENDTX) StoreBits(mask, b EVENT_ENDTX) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_ENDTX) SetBits(mask EVENT_ENDTX)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_ENDTX) ClearBits(mask EVENT_ENDTX)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_ENDTX) Load() EVENT_ENDTX             { return EVENT_ENDTX(r.U32.Load()) }
func (r *REVENT_ENDTX) Store(b EVENT_ENDTX)           { r.U32.Store(uint32(b)) }

type RMEVENT_ENDTX struct{ mmio.UM32 }

func (rm RMEVENT_ENDTX) Load() EVENT_ENDTX   { return EVENT_ENDTX(rm.UM32.Load()) }
func (rm RMEVENT_ENDTX) Store(b EVENT_ENDTX) { rm.UM32.Store(uint32(b)) }

type EVENT_ERROR uint32

type REVENT_ERROR struct{ mmio.U32 }

func (r *REVENT_ERROR) LoadBits(mask EVENT_ERROR) EVENT_ERROR {
	return EVENT_ERROR(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_ERROR) StoreBits(mask, b EVENT_ERROR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_ERROR) SetBits(mask EVENT_ERROR)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_ERROR) ClearBits(mask EVENT_ERROR)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_ERROR) Load() EVENT_ERROR             { return EVENT_ERROR(r.U32.Load()) }
func (r *REVENT_ERROR) Store(b EVENT_ERROR)           { r.U32.Store(uint32(b)) }

type RMEVENT_ERROR struct{ mmio.UM32 }

func (rm RMEVENT_ERROR) Load() EVENT_ERROR   { return EVENT_ERROR(rm.UM32.Load()) }
func (rm RMEVENT_ERROR) Store(b EVENT_ERROR) { rm.UM32.Store(uint32(b)) }

type EVENT_RXTO uint32

type REVENT_RXTO struct{ mmio.U32 }

func (r *REVENT_RXTO) LoadBits(mask EVENT_RXTO) EVENT_RXTO {
	return EVENT_RXTO(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_RXTO) StoreBits(mask, b EVENT_RXTO) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REVENT_RXTO) SetBits(mask EVENT_RXTO)      { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_RXTO) ClearBits(mask EVENT_RXTO)    { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_RXTO) Load() EVENT_RXTO             { return EVENT_RXTO(r.U32.Load()) }
func (r *REVENT_RXTO) Store(b EVENT_RXTO)           { r.U32.Store(uint32(b)) }

type RMEVENT_RXTO struct{ mmio.UM32 }

func (rm RMEVENT_RXTO) Load() EVENT_RXTO   { return EVENT_RXTO(rm.UM32.Load()) }
func (rm RMEVENT_RXTO) Store(b EVENT_RXTO) { rm.UM32.Store(uint32(b)) }

type EVENT_RXSTARTED uint32

type REVENT_RXSTARTED struct{ mmio.U32 }

func (r *REVENT_RXSTARTED) LoadBits(mask EVENT_RXSTARTED) EVENT_RXSTARTED {
	return EVENT_RXSTARTED(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_RXSTARTED) StoreBits(mask, b EVENT_RXSTARTED) {
	r.U32.StoreBits(uint32(mask), uint32(b))
}
func (r *REVENT_RXSTARTED) SetBits(mask EVENT_RXSTARTED)   { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_RXSTARTED) ClearBits(mask EVENT_RXSTARTED) { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_RXSTARTED) Load() EVENT_RXSTARTED          { return EVENT_RXSTARTED(r.U32.Load()) }
func (r *REVENT_RXSTARTED) Store(b EVENT_RXSTARTED)        { r.U32.Store(uint32(b)) }

type RMEVENT_RXSTARTED struct{ mmio.UM32 }

func (rm RMEVENT_RXSTARTED) Load() EVENT_RXSTARTED   { return EVENT_RXSTARTED(rm.UM32.Load()) }
func (rm RMEVENT_RXSTARTED) Store(b EVENT_RXSTARTED) { rm.UM32.Store(uint32(b)) }

type EVENT_TXSTARTED uint32

type REVENT_TXSTARTED struct{ mmio.U32 }

func (r *REVENT_TXSTARTED) LoadBits(mask EVENT_TXSTARTED) EVENT_TXSTARTED {
	return EVENT_TXSTARTED(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_TXSTARTED) StoreBits(mask, b EVENT_TXSTARTED) {
	r.U32.StoreBits(uint32(mask), uint32(b))
}
func (r *REVENT_TXSTARTED) SetBits(mask EVENT_TXSTARTED)   { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_TXSTARTED) ClearBits(mask EVENT_TXSTARTED) { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_TXSTARTED) Load() EVENT_TXSTARTED          { return EVENT_TXSTARTED(r.U32.Load()) }
func (r *REVENT_TXSTARTED) Store(b EVENT_TXSTARTED)        { r.U32.Store(uint32(b)) }

type RMEVENT_TXSTARTED struct{ mmio.UM32 }

func (rm RMEVENT_TXSTARTED) Load() EVENT_TXSTARTED   { return EVENT_TXSTARTED(rm.UM32.Load()) }
func (rm RMEVENT_TXSTARTED) Store(b EVENT_TXSTARTED) { rm.UM32.Store(uint32(b)) }

type EVENT_TXSTOPPED uint32

type REVENT_TXSTOPPED struct{ mmio.U32 }

func (r *REVENT_TXSTOPPED) LoadBits(mask EVENT_TXSTOPPED) EVENT_TXSTOPPED {
	return EVENT_TXSTOPPED(r.U32.LoadBits(uint32(mask)))
}
func (r *REVENT_TXSTOPPED) StoreBits(mask, b EVENT_TXSTOPPED) {
	r.U32.StoreBits(uint32(mask), uint32(b))
}
func (r *REVENT_TXSTOPPED) SetBits(mask EVENT_TXSTOPPED)   { r.U32.SetBits(uint32(mask)) }
func (r *REVENT_TXSTOPPED) ClearBits(mask EVENT_TXSTOPPED) { r.U32.ClearBits(uint32(mask)) }
func (r *REVENT_TXSTOPPED) Load() EVENT_TXSTOPPED          { return EVENT_TXSTOPPED(r.U32.Load()) }
func (r *REVENT_TXSTOPPED) Store(b EVENT_TXSTOPPED)        { r.U32.Store(uint32(b)) }

type RMEVENT_TXSTOPPED struct{ mmio.UM32 }

func (rm RMEVENT_TXSTOPPED) Load() EVENT_TXSTOPPED   { return EVENT_TXSTOPPED(rm.UM32.Load()) }
func (rm RMEVENT_TXSTOPPED) Store(b EVENT_TXSTOPPED) { rm.UM32.Store(uint32(b)) }

type SHORTS uint32

type RSHORTS struct{ mmio.U32 }

func (r *RSHORTS) LoadBits(mask SHORTS) SHORTS { return SHORTS(r.U32.LoadBits(uint32(mask))) }
func (r *RSHORTS) StoreBits(mask, b SHORTS)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSHORTS) SetBits(mask SHORTS)         { r.U32.SetBits(uint32(mask)) }
func (r *RSHORTS) ClearBits(mask SHORTS)       { r.U32.ClearBits(uint32(mask)) }
func (r *RSHORTS) Load() SHORTS                { return SHORTS(r.U32.Load()) }
func (r *RSHORTS) Store(b SHORTS)              { r.U32.Store(uint32(b)) }

type RMSHORTS struct{ mmio.UM32 }

func (rm RMSHORTS) Load() SHORTS   { return SHORTS(rm.UM32.Load()) }
func (rm RMSHORTS) Store(b SHORTS) { rm.UM32.Store(uint32(b)) }

func ENDRX_STARTRX_(p *Periph) RMSHORTS {
	return RMSHORTS{mmio.UM32{&p.SHORTS.U32, uint32(ENDRX_STARTRX)}}
}

func ENDRX_STOPRX_(p *Periph) RMSHORTS {
	return RMSHORTS{mmio.UM32{&p.SHORTS.U32, uint32(ENDRX_STOPRX)}}
}

type INTEN uint32

type RINTEN struct{ mmio.U32 }

func (r *RINTEN) LoadBits(mask INTEN) INTEN { return INTEN(r.U32.LoadBits(uint32(mask))) }
func (r *RINTEN) StoreBits(mask, b INTEN)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RINTEN) SetBits(mask INTEN)        { r.U32.SetBits(uint32(mask)) }
func (r *RINTEN) ClearBits(mask INTEN)      { r.U32.ClearBits(uint32(mask)) }
func (r *RINTEN) Load() INTEN               { return INTEN(r.U32.Load()) }
func (r *RINTEN) Store(b INTEN)             { r.U32.Store(uint32(b)) }

type RMINTEN struct{ mmio.UM32 }

func (rm RMINTEN) Load() INTEN   { return INTEN(rm.UM32.Load()) }
func (rm RMINTEN) Store(b INTEN) { rm.UM32.Store(uint32(b)) }

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

type ERRORSRC uint32

type RERRORSRC struct{ mmio.U32 }

func (r *RERRORSRC) LoadBits(mask ERRORSRC) ERRORSRC { return ERRORSRC(r.U32.LoadBits(uint32(mask))) }
func (r *RERRORSRC) StoreBits(mask, b ERRORSRC)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RERRORSRC) SetBits(mask ERRORSRC)           { r.U32.SetBits(uint32(mask)) }
func (r *RERRORSRC) ClearBits(mask ERRORSRC)         { r.U32.ClearBits(uint32(mask)) }
func (r *RERRORSRC) Load() ERRORSRC                  { return ERRORSRC(r.U32.Load()) }
func (r *RERRORSRC) Store(b ERRORSRC)                { r.U32.Store(uint32(b)) }

type RMERRORSRC struct{ mmio.UM32 }

func (rm RMERRORSRC) Load() ERRORSRC   { return ERRORSRC(rm.UM32.Load()) }
func (rm RMERRORSRC) Store(b ERRORSRC) { rm.UM32.Store(uint32(b)) }

func EOVERRUN_(p *Periph) RMERRORSRC {
	return RMERRORSRC{mmio.UM32{&p.ERRORSRC.U32, uint32(EOVERRUN)}}
}

func EPARITY_(p *Periph) RMERRORSRC {
	return RMERRORSRC{mmio.UM32{&p.ERRORSRC.U32, uint32(EPARITY)}}
}

func EFRAMING_(p *Periph) RMERRORSRC {
	return RMERRORSRC{mmio.UM32{&p.ERRORSRC.U32, uint32(EFRAMING)}}
}

func EBREAK_(p *Periph) RMERRORSRC {
	return RMERRORSRC{mmio.UM32{&p.ERRORSRC.U32, uint32(EBREAK)}}
}

type ENABLE uint32

type RENABLE struct{ mmio.U32 }

func (r *RENABLE) LoadBits(mask ENABLE) ENABLE { return ENABLE(r.U32.LoadBits(uint32(mask))) }
func (r *RENABLE) StoreBits(mask, b ENABLE)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RENABLE) SetBits(mask ENABLE)         { r.U32.SetBits(uint32(mask)) }
func (r *RENABLE) ClearBits(mask ENABLE)       { r.U32.ClearBits(uint32(mask)) }
func (r *RENABLE) Load() ENABLE                { return ENABLE(r.U32.Load()) }
func (r *RENABLE) Store(b ENABLE)              { r.U32.Store(uint32(b)) }

type RMENABLE struct{ mmio.UM32 }

func (rm RMENABLE) Load() ENABLE   { return ENABLE(rm.UM32.Load()) }
func (rm RMENABLE) Store(b ENABLE) { rm.UM32.Store(uint32(b)) }

func EN_(p *Periph) RMENABLE {
	return RMENABLE{mmio.UM32{&p.ENABLE.U32, uint32(EN)}}
}

type PSEL_RTS uint32

type RPSEL_RTS struct{ mmio.U32 }

func (r *RPSEL_RTS) LoadBits(mask PSEL_RTS) PSEL_RTS { return PSEL_RTS(r.U32.LoadBits(uint32(mask))) }
func (r *RPSEL_RTS) StoreBits(mask, b PSEL_RTS)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPSEL_RTS) SetBits(mask PSEL_RTS)           { r.U32.SetBits(uint32(mask)) }
func (r *RPSEL_RTS) ClearBits(mask PSEL_RTS)         { r.U32.ClearBits(uint32(mask)) }
func (r *RPSEL_RTS) Load() PSEL_RTS                  { return PSEL_RTS(r.U32.Load()) }
func (r *RPSEL_RTS) Store(b PSEL_RTS)                { r.U32.Store(uint32(b)) }

type RMPSEL_RTS struct{ mmio.UM32 }

func (rm RMPSEL_RTS) Load() PSEL_RTS   { return PSEL_RTS(rm.UM32.Load()) }
func (rm RMPSEL_RTS) Store(b PSEL_RTS) { rm.UM32.Store(uint32(b)) }

type PSEL_TXD uint32

type RPSEL_TXD struct{ mmio.U32 }

func (r *RPSEL_TXD) LoadBits(mask PSEL_TXD) PSEL_TXD { return PSEL_TXD(r.U32.LoadBits(uint32(mask))) }
func (r *RPSEL_TXD) StoreBits(mask, b PSEL_TXD)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPSEL_TXD) SetBits(mask PSEL_TXD)           { r.U32.SetBits(uint32(mask)) }
func (r *RPSEL_TXD) ClearBits(mask PSEL_TXD)         { r.U32.ClearBits(uint32(mask)) }
func (r *RPSEL_TXD) Load() PSEL_TXD                  { return PSEL_TXD(r.U32.Load()) }
func (r *RPSEL_TXD) Store(b PSEL_TXD)                { r.U32.Store(uint32(b)) }

type RMPSEL_TXD struct{ mmio.UM32 }

func (rm RMPSEL_TXD) Load() PSEL_TXD   { return PSEL_TXD(rm.UM32.Load()) }
func (rm RMPSEL_TXD) Store(b PSEL_TXD) { rm.UM32.Store(uint32(b)) }

type PSEL_CTS uint32

type RPSEL_CTS struct{ mmio.U32 }

func (r *RPSEL_CTS) LoadBits(mask PSEL_CTS) PSEL_CTS { return PSEL_CTS(r.U32.LoadBits(uint32(mask))) }
func (r *RPSEL_CTS) StoreBits(mask, b PSEL_CTS)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPSEL_CTS) SetBits(mask PSEL_CTS)           { r.U32.SetBits(uint32(mask)) }
func (r *RPSEL_CTS) ClearBits(mask PSEL_CTS)         { r.U32.ClearBits(uint32(mask)) }
func (r *RPSEL_CTS) Load() PSEL_CTS                  { return PSEL_CTS(r.U32.Load()) }
func (r *RPSEL_CTS) Store(b PSEL_CTS)                { r.U32.Store(uint32(b)) }

type RMPSEL_CTS struct{ mmio.UM32 }

func (rm RMPSEL_CTS) Load() PSEL_CTS   { return PSEL_CTS(rm.UM32.Load()) }
func (rm RMPSEL_CTS) Store(b PSEL_CTS) { rm.UM32.Store(uint32(b)) }

type PSEL_RXD uint32

type RPSEL_RXD struct{ mmio.U32 }

func (r *RPSEL_RXD) LoadBits(mask PSEL_RXD) PSEL_RXD { return PSEL_RXD(r.U32.LoadBits(uint32(mask))) }
func (r *RPSEL_RXD) StoreBits(mask, b PSEL_RXD)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPSEL_RXD) SetBits(mask PSEL_RXD)           { r.U32.SetBits(uint32(mask)) }
func (r *RPSEL_RXD) ClearBits(mask PSEL_RXD)         { r.U32.ClearBits(uint32(mask)) }
func (r *RPSEL_RXD) Load() PSEL_RXD                  { return PSEL_RXD(r.U32.Load()) }
func (r *RPSEL_RXD) Store(b PSEL_RXD)                { r.U32.Store(uint32(b)) }

type RMPSEL_RXD struct{ mmio.UM32 }

func (rm RMPSEL_RXD) Load() PSEL_RXD   { return PSEL_RXD(rm.UM32.Load()) }
func (rm RMPSEL_RXD) Store(b PSEL_RXD) { rm.UM32.Store(uint32(b)) }

type BAUDRATE uint32

type RBAUDRATE struct{ mmio.U32 }

func (r *RBAUDRATE) LoadBits(mask BAUDRATE) BAUDRATE { return BAUDRATE(r.U32.LoadBits(uint32(mask))) }
func (r *RBAUDRATE) StoreBits(mask, b BAUDRATE)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBAUDRATE) SetBits(mask BAUDRATE)           { r.U32.SetBits(uint32(mask)) }
func (r *RBAUDRATE) ClearBits(mask BAUDRATE)         { r.U32.ClearBits(uint32(mask)) }
func (r *RBAUDRATE) Load() BAUDRATE                  { return BAUDRATE(r.U32.Load()) }
func (r *RBAUDRATE) Store(b BAUDRATE)                { r.U32.Store(uint32(b)) }

type RMBAUDRATE struct{ mmio.UM32 }

func (rm RMBAUDRATE) Load() BAUDRATE   { return BAUDRATE(rm.UM32.Load()) }
func (rm RMBAUDRATE) Store(b BAUDRATE) { rm.UM32.Store(uint32(b)) }

func BR_(p *Periph) RMBAUDRATE {
	return RMBAUDRATE{mmio.UM32{&p.BAUDRATE.U32, uint32(BR)}}
}

type RXD_PTR uint32

type RRXD_PTR struct{ mmio.U32 }

func (r *RRXD_PTR) LoadBits(mask RXD_PTR) RXD_PTR { return RXD_PTR(r.U32.LoadBits(uint32(mask))) }
func (r *RRXD_PTR) StoreBits(mask, b RXD_PTR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXD_PTR) SetBits(mask RXD_PTR)          { r.U32.SetBits(uint32(mask)) }
func (r *RRXD_PTR) ClearBits(mask RXD_PTR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RRXD_PTR) Load() RXD_PTR                 { return RXD_PTR(r.U32.Load()) }
func (r *RRXD_PTR) Store(b RXD_PTR)               { r.U32.Store(uint32(b)) }

type RMRXD_PTR struct{ mmio.UM32 }

func (rm RMRXD_PTR) Load() RXD_PTR   { return RXD_PTR(rm.UM32.Load()) }
func (rm RMRXD_PTR) Store(b RXD_PTR) { rm.UM32.Store(uint32(b)) }

type RXD_MAXCNT uint32

type RRXD_MAXCNT struct{ mmio.U32 }

func (r *RRXD_MAXCNT) LoadBits(mask RXD_MAXCNT) RXD_MAXCNT {
	return RXD_MAXCNT(r.U32.LoadBits(uint32(mask)))
}
func (r *RRXD_MAXCNT) StoreBits(mask, b RXD_MAXCNT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXD_MAXCNT) SetBits(mask RXD_MAXCNT)      { r.U32.SetBits(uint32(mask)) }
func (r *RRXD_MAXCNT) ClearBits(mask RXD_MAXCNT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRXD_MAXCNT) Load() RXD_MAXCNT             { return RXD_MAXCNT(r.U32.Load()) }
func (r *RRXD_MAXCNT) Store(b RXD_MAXCNT)           { r.U32.Store(uint32(b)) }

type RMRXD_MAXCNT struct{ mmio.UM32 }

func (rm RMRXD_MAXCNT) Load() RXD_MAXCNT   { return RXD_MAXCNT(rm.UM32.Load()) }
func (rm RMRXD_MAXCNT) Store(b RXD_MAXCNT) { rm.UM32.Store(uint32(b)) }

type RXD_AMOUNT uint32

type RRXD_AMOUNT struct{ mmio.U32 }

func (r *RRXD_AMOUNT) LoadBits(mask RXD_AMOUNT) RXD_AMOUNT {
	return RXD_AMOUNT(r.U32.LoadBits(uint32(mask)))
}
func (r *RRXD_AMOUNT) StoreBits(mask, b RXD_AMOUNT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXD_AMOUNT) SetBits(mask RXD_AMOUNT)      { r.U32.SetBits(uint32(mask)) }
func (r *RRXD_AMOUNT) ClearBits(mask RXD_AMOUNT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRXD_AMOUNT) Load() RXD_AMOUNT             { return RXD_AMOUNT(r.U32.Load()) }
func (r *RRXD_AMOUNT) Store(b RXD_AMOUNT)           { r.U32.Store(uint32(b)) }

type RMRXD_AMOUNT struct{ mmio.UM32 }

func (rm RMRXD_AMOUNT) Load() RXD_AMOUNT   { return RXD_AMOUNT(rm.UM32.Load()) }
func (rm RMRXD_AMOUNT) Store(b RXD_AMOUNT) { rm.UM32.Store(uint32(b)) }

type TXD_PTR uint32

type RTXD_PTR struct{ mmio.U32 }

func (r *RTXD_PTR) LoadBits(mask TXD_PTR) TXD_PTR { return TXD_PTR(r.U32.LoadBits(uint32(mask))) }
func (r *RTXD_PTR) StoreBits(mask, b TXD_PTR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXD_PTR) SetBits(mask TXD_PTR)          { r.U32.SetBits(uint32(mask)) }
func (r *RTXD_PTR) ClearBits(mask TXD_PTR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RTXD_PTR) Load() TXD_PTR                 { return TXD_PTR(r.U32.Load()) }
func (r *RTXD_PTR) Store(b TXD_PTR)               { r.U32.Store(uint32(b)) }

type RMTXD_PTR struct{ mmio.UM32 }

func (rm RMTXD_PTR) Load() TXD_PTR   { return TXD_PTR(rm.UM32.Load()) }
func (rm RMTXD_PTR) Store(b TXD_PTR) { rm.UM32.Store(uint32(b)) }

type TXD_MAXCNT uint32

type RTXD_MAXCNT struct{ mmio.U32 }

func (r *RTXD_MAXCNT) LoadBits(mask TXD_MAXCNT) TXD_MAXCNT {
	return TXD_MAXCNT(r.U32.LoadBits(uint32(mask)))
}
func (r *RTXD_MAXCNT) StoreBits(mask, b TXD_MAXCNT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXD_MAXCNT) SetBits(mask TXD_MAXCNT)      { r.U32.SetBits(uint32(mask)) }
func (r *RTXD_MAXCNT) ClearBits(mask TXD_MAXCNT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTXD_MAXCNT) Load() TXD_MAXCNT             { return TXD_MAXCNT(r.U32.Load()) }
func (r *RTXD_MAXCNT) Store(b TXD_MAXCNT)           { r.U32.Store(uint32(b)) }

type RMTXD_MAXCNT struct{ mmio.UM32 }

func (rm RMTXD_MAXCNT) Load() TXD_MAXCNT   { return TXD_MAXCNT(rm.UM32.Load()) }
func (rm RMTXD_MAXCNT) Store(b TXD_MAXCNT) { rm.UM32.Store(uint32(b)) }

type TXD_AMOUNT uint32

type RTXD_AMOUNT struct{ mmio.U32 }

func (r *RTXD_AMOUNT) LoadBits(mask TXD_AMOUNT) TXD_AMOUNT {
	return TXD_AMOUNT(r.U32.LoadBits(uint32(mask)))
}
func (r *RTXD_AMOUNT) StoreBits(mask, b TXD_AMOUNT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXD_AMOUNT) SetBits(mask TXD_AMOUNT)      { r.U32.SetBits(uint32(mask)) }
func (r *RTXD_AMOUNT) ClearBits(mask TXD_AMOUNT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTXD_AMOUNT) Load() TXD_AMOUNT             { return TXD_AMOUNT(r.U32.Load()) }
func (r *RTXD_AMOUNT) Store(b TXD_AMOUNT)           { r.U32.Store(uint32(b)) }

type RMTXD_AMOUNT struct{ mmio.UM32 }

func (rm RMTXD_AMOUNT) Load() TXD_AMOUNT   { return TXD_AMOUNT(rm.UM32.Load()) }
func (rm RMTXD_AMOUNT) Store(b TXD_AMOUNT) { rm.UM32.Store(uint32(b)) }

type CONFIG uint32

type RCONFIG struct{ mmio.U32 }

func (r *RCONFIG) LoadBits(mask CONFIG) CONFIG { return CONFIG(r.U32.LoadBits(uint32(mask))) }
func (r *RCONFIG) StoreBits(mask, b CONFIG)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCONFIG) SetBits(mask CONFIG)         { r.U32.SetBits(uint32(mask)) }
func (r *RCONFIG) ClearBits(mask CONFIG)       { r.U32.ClearBits(uint32(mask)) }
func (r *RCONFIG) Load() CONFIG                { return CONFIG(r.U32.Load()) }
func (r *RCONFIG) Store(b CONFIG)              { r.U32.Store(uint32(b)) }

type RMCONFIG struct{ mmio.UM32 }

func (rm RMCONFIG) Load() CONFIG   { return CONFIG(rm.UM32.Load()) }
func (rm RMCONFIG) Store(b CONFIG) { rm.UM32.Store(uint32(b)) }

func HWFC_(p *Periph) RMCONFIG {
	return RMCONFIG{mmio.UM32{&p.CONFIG.U32, uint32(HWFC)}}
}

func PARITY_(p *Periph) RMCONFIG {
	return RMCONFIG{mmio.UM32{&p.CONFIG.U32, uint32(PARITY)}}
}

func STOP_(p *Periph) RMCONFIG {
	return RMCONFIG{mmio.UM32{&p.CONFIG.U32, uint32(STOP)}}
}
